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
	"github.com/shyptr/hello-world-web/article_service/model"
	"github.com/shyptr/hello-world-web/article_service/resolve"
	"github.com/shyptr/hello-world-web/middleware"
	"sync"
)

var (
	as   *articleService
	once sync.Once
)

type articleService struct {
	schema  *system.Schema
	execute *execution.Executor
}

func (s *articleService) Execute(ctx context.Context, r *federation.FederationRequest) (*federation.FederationResponse, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	request := federation.ConvertRequest(r)
	root := s.schema.Query
	if request.Kind == string(ast.Mutation) {
		root = s.schema.Mutation
	}
	logger.Info().Interface("request", request).Send()
	execute, multiError := s.execute.Execute(ctx, root, nil, request.SelectionSet)
	return federation.ConvertToResponse(execute, multiError), nil
}

func (s *articleService) Introspection(ctx context.Context, r *federation.Null) (*federation.FederationResponse, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	schemaJSON, err := introspection.ComputeSchemaJSON(s.schema)
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, err
	}
	return &federation.FederationResponse{Data: &any.Any{Value: schemaJSON}}, nil
}

func (s *articleService) Subscription(r *federation.FederationRequest, server federation.FederationService_SubscriptionServer) error {
	return nil
}

func RegisterArticleType(schema *schemabuilder.Schema) {
	schema.Enum("ArticleState", model.ArticleState(0), model.ArticleStateString)

	article := schema.Object("Article", model.Article{})

	schema.Object("ArticleEx", model.ArticleEx{})
	article.FieldFunc("count", resolve.ArticleResolve.ArticleCount)

	query := schema.Query()
	query.FieldFunc("GetArticle", resolve.ArticleResolve.Article)
	schemabuilder.RelayKey(model.Article{}, "id")
	query.FieldFunc("Articles", resolve.ArticleResolve.Articles, schemabuilder.RelayConnection)

	mutation := schema.Mutation()
	mutation.FieldFunc("AddArticle", resolve.ArticleResolve.AddArticle, middleware.LoginNeed())
	mutation.FieldFunc("UpdateArticle", resolve.ArticleResolve.UpdateArticle, middleware.LoginNeed())
	mutation.FieldFunc("DeleteArticle", resolve.ArticleResolve.DeleteArticle, middleware.LoginNeed())
}

func RegisterTagType(schema *schemabuilder.Schema) {
	schema.Object("Tag", model.Tag{})

	schema.Query().FieldFunc("Tag", resolve.TagResolve.Tag)

	schema.Mutation().FieldFunc("AddTag", resolve.TagResolve.AddTag, middleware.LoginNeed())
}

func NewService() *articleService {
	once.Do(func() {
		schemabuilder.NewValidate()
		builder := schemabuilder.NewSchema()
		RegisterArticleType(builder)
		RegisterTagType(builder)
		schema := builder.MustBuild()
		introspection.AddIntrospectionToSchema(schema)
		as = &articleService{
			schema:  schema,
			execute: &execution.Executor{},
		}
	})
	return as
}
