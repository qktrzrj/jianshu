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
	"github.com/shyptr/hello-world-web/comment_service/model"
	"github.com/shyptr/hello-world-web/comment_service/resolve"
	"github.com/shyptr/hello-world-web/middleware"
	"sync"
)

var (
	cs   *commentService
	once sync.Once
)

type commentService struct {
	schema  *system.Schema
	execute *execution.Executor
}

func (s *commentService) Subscription(ctx context.Context, req *federation.FederationRequest, stream federation.FederationService_SubscriptionStream) error {
	return nil
}

func (s *commentService) Introspection(ctx context.Context, null *federation.Null, response *federation.FederationResponse) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	schemaJSON, err := introspection.ComputeSchemaJSON(s.schema)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	response.Data = &any.Any{Value: schemaJSON}
	return nil
}

func (s *commentService) Execute(ctx context.Context, r *federation.FederationRequest, resp *federation.FederationResponse) error {
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

func registerCommentType(schema *schemabuilder.Schema) {
	schema.Enum("CommentState", model.CommentState(0), model.CommentStateString)

	comment := schema.Object("Comment", model.Comment{})
	comment.FieldFunc("replies", resolve.CommentResolve.Replies)

	schema.Object("CommentReply", model.CommentReply{})

	query := schema.Query()
	schemabuilder.RelayKey(model.Comment{}, "id")
	query.FieldFunc("Comments", resolve.CommentResolve.Comments, schemabuilder.RelayConnection)

	mutation := schema.Mutation()
	mutation.FieldFunc("Comment", resolve.CommentResolve.AddComment, middleware.LoginNeed())
	mutation.FieldFunc("DeleteComment", resolve.CommentResolve.RemoveComment, middleware.LoginNeed())
	mutation.FieldFunc("Reply", resolve.CommentResolve.Reply, middleware.LoginNeed())
	mutation.FieldFunc("DeleteReply", resolve.CommentResolve.RemoveReply, middleware.LoginNeed())
}

func NewService() *commentService {
	once.Do(func() {
		schemabuilder.NewValidate()
		builder := schemabuilder.NewSchema()
		registerCommentType(builder)
		schema := builder.MustBuild()
		introspection.AddIntrospectionToSchema(schema)
		cs = &commentService{schema: schema, execute: &execution.Executor{}}
	})
	return cs
}
