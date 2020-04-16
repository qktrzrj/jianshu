package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql/federation"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/graphql/system"
	"github.com/shyptr/graphql/system/ast"
	"github.com/shyptr/graphql/system/execution"
	"github.com/shyptr/graphql/system/introspection"
	"github.com/shyptr/hello-world-web/middleware"
	"github.com/shyptr/hello-world-web/zan_service/model"
	"github.com/shyptr/hello-world-web/zan_service/resolve"
	"sync"
)

var (
	zs   *zanService
	once sync.Once
)

type zanService struct {
	schema  *system.Schema
	execute *execution.Executor
}

func (s *zanService) Subscription(ctx context.Context, req *federation.FederationRequest, stream federation.FederationService_SubscriptionStream) error {
	return nil
}

func (s *zanService) Execute(ctx context.Context, r *federation.FederationRequest, resp *federation.FederationResponse) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	request := federation.ConvertRequest(r)
	root := s.schema.Query
	if request.Kind == string(ast.Mutation) {
		root = s.schema.Mutation
	}
	logger.Info().Interface("request", request).Send()
	execute, multiError := s.execute.Execute(ctx, root, nil, request.SelectionSet)
	response := federation.ConvertToResponse(execute, multiError)
	resp.Data, resp.Errors = response.GetData(), response.GetErrors()
	return nil
}

func (s *zanService) Introspection(ctx context.Context, null *federation.Null, response *federation.FederationResponse) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	schemaJSON, err := introspection.ComputeSchemaJSON(s.schema)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	response.Data = &any.Any{Value: schemaJSON}
	return nil
}

func registerZanType(schema *schemabuilder.Schema) {
	schema.Enum("ObjType", model.ObjType(0), model.ObjTypeString)
	schema.Object("Zan", model.Zan{})

	mutation := schema.Mutation()
	mutation.FieldFunc("Zan", resolve.ZanResolve.AddZan, middleware.LoginNeed())
	mutation.FieldFunc("CancelZan", resolve.ZanResolve.RemoveZan, middleware.LoginNeed())

}

func NewService() *zanService {
	once.Do(func() {
		schemabuilder.NewValidate()
		builder := schemabuilder.NewSchema()
		registerZanType(builder)
		schema := builder.MustBuild()
		introspection.AddIntrospectionToSchema(schema)
		zs = &zanService{schema: schema, execute: &execution.Executor{}}
	})
	return zs
}
