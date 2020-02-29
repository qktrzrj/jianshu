package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

var (
	schema        graphql.Schema
	queryType     *graphql.Object
	mutationType  *graphql.Object
	subscriptType *graphql.Object
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func Register(e *gin.Engine) {
	queryType = graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: graphql.Fields{
		"test": {
			Name: "test",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "test", nil
			},
			Description: "test",
		},
	}})
	mutationType = graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: graphql.Fields{
		"test": {
			Name: "test",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "test", nil
			},
			Description: "test",
		},
	}})
	subscriptType = graphql.NewObject(graphql.ObjectConfig{Name: "Subscription", Fields: graphql.Fields{
		"test": {
			Name: "test",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "test", nil
			},
			Description: "test",
		},
	}})
	schemaConfig := graphql.SchemaConfig{
		Query:        queryType,
		Mutation:     mutationType,
		Subscription: subscriptType,
	}
	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	router := func(ctx *gin.Context) {
		h.ContextHandler(context.Background(), ctx.Writer, ctx.Request)
	}
	// graphql的web界面，只有admin才能进入
	e.GET("/graphql", router)
	e.POST("/graphql", router)
	e.OPTIONS("/graphql", router)

	e.GET("/query", query)
	e.OPTIONS("/query", query)
	e.POST("/query", query)

}

func query(ctx *gin.Context) {
	requestOption := &RequestOptions{}
	_ = ctx.Bind(requestOption)
	ctx.Set("operationName", requestOption.OperationName)
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  requestOption.Query,
		VariableValues: requestOption.Variables,
		OperationName:  requestOption.OperationName,
	})

	ctx.JSON(http.StatusOK, result)
}
