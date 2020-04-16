package router

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/errors"
	"github.com/shyptr/graphql/federation"
	"github.com/shyptr/graphql/system/execution"
	"github.com/shyptr/graphql/system/introspection"
	"github.com/shyptr/hello-world-web/setting"
	"github.com/shyptr/hello-world-web/util"
	metadata2 "google.golang.org/grpc/metadata"
	"net/http"
	"strings"
	"time"
)

var node map[string]federation.FederationServiceServer

var planner *federation.Planner

func init() {
	setting.InitConfig()
	util.InitLogger(nil)
	etcdReg := etcd.NewRegistry(func(options *registry.Options) {
		etcdCfg := setting.GetETCDConfig()
		options.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
		options.Timeout = time.Second * 15
	})
	node = make(map[string]federation.FederationServiceServer, 4)
	userService := micro.NewService(micro.Name("user_service.client"), micro.Registry(etcdReg))
	userService.Init()
	node["user_service"] = federation.NewFederationService("user_service", userService.Client())
	articleService := micro.NewService(micro.Name("article_service.client"), micro.Registry(etcdReg))
	articleService.Init()
	node["article_service"] = federation.NewFederationService("article_service", userService.Client())
	commentService := micro.NewService(micro.Name("comment_service.client"), micro.Registry(etcdReg))
	commentService.Init()
	node["comment_service"] = federation.NewFederationService("comment_service", userService.Client())
	zanService := micro.NewService(micro.Name("zan_service.client"), micro.Registry(etcdReg))
	zanService.Init()
	node["zan_service"] = federation.NewFederationService("zan_service", userService.Client())
}

func Register(engine *gin.Engine) {
	schemas := make(map[string]string)
	etcdCfg := setting.GetETCDConfig()
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())},
	})
	defer cli.Close()
	for name, c := range node {
		res, err := c.Introspection(context.Background(), &federation.Null{})
		if err != nil {
			panic(err)
		}
		schemas[name] = string(res.GetData().Value)
	}
	convertSchema, err := federation.ConvertSchema(schemas)
	if err != nil {
		panic(err)
	}
	introspection.AddIntrospectionToSchema(convertSchema.Schema)
	planner, err = federation.NewPlaner(convertSchema)
	if err != nil {
		panic(err)
	}
	engine.GET("/", func(c *gin.Context) {
		graphql.GraphiQLHandler().ServeHTTP(c.Writer, c.Request)
	})
	engine.OPTIONS("/query")
	engine.POST("/query", func(c *gin.Context) {
		logger := c.Value("logger").(zerolog.Logger)
		var params execution.Params
		err := c.BindJSON(&params)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			c.JSON(http.StatusBadRequest, graphql.Response{Errors: errors.MultiError{errors.New(err.Error())}})
			return
		}
		if params.OperationName == "IntrospectionQuery" || strings.Contains(params.Query, "IntrospectionQuery") {
			data, _ := execution.Do(convertSchema.Schema, params)
			c.JSON(http.StatusOK, graphql.Response{Data: data.(map[string]interface{})})
			return
		}
		planner := planner
		plan, err := federation.MustPlan(planner, params)
		if err != nil {
			logger.Error().Caller().Err(err).Send()
			c.JSON(http.StatusBadRequest, graphql.Response{Errors: errors.MultiError{errors.New(err.Error())}})
			return
		}
		response := graphql.Response{
			Data: make(map[string]interface{}),
		}
		for _, p := range plan.After {
			author, _ := c.Cookie("me")
			ctx := metadata.NewContext(context.Background(), metadata.Metadata{"me": author})
			federationResponse, err := node[p.Service].Execute(ctx, &federation.FederationRequest{
				Kind:         p.Kind,
				SelectionSet: federation.ConvertToSelectionSet(p.SelectionSet),
			})
			if err != nil {
				logger.Error().Caller().Err(err).Send()
				response.Errors = append(response.Errors, errors.New(err.Error()))
				continue
			}
			convertResponse := federation.ConvertResponse(federationResponse)
			for k, v := range convertResponse.Data.(map[string]interface{}) {
				response.Data.(map[string]interface{})[k] = v
			}
			response.Errors = append(response.Errors, convertResponse.Errors...)
			if p.Service == "user_service" {
				logger.Info().Msg("get cookie")
				mdata, _ := metadata2.FromOutgoingContext(ctx)
				logger.Info().Interface("metadata", mdata).Send()
				cookie := mdata.Get("cookie")
				logger.Info().Interface("cookie", cookie).Send()
				if len(cookie) > 0 {
					c.Set("Set-Cookie", cookie[0])
				}
			}
		}
		logger.Info().Interface("response", response).Send()
		c.JSON(http.StatusOK, response)
	})
}
