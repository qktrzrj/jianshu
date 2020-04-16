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
	"github.com/shyptr/hello-world-web/user_service/model"
	"github.com/shyptr/hello-world-web/user_service/resolve"
	"sync"
)

var (
	us   *userService
	once sync.Once
)

type userService struct {
	schema  *system.Schema
	execute *execution.Executor
}

func (s *userService) Subscription(ctx context.Context, req *federation.FederationRequest, stream federation.FederationService_SubscriptionStream) error {
	return nil
}

func (s *userService) Introspection(ctx context.Context, null *federation.Null, response *federation.FederationResponse) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	schemaJSON, err := introspection.ComputeSchemaJSON(s.schema)
	if err != nil {
		logger.Error().Caller().Err(err).Send()
		return err
	}
	response.Data = &any.Any{Value: schemaJSON}
	return nil
}

func (s *userService) Execute(ctx context.Context, r *federation.FederationRequest, resp *federation.FederationResponse) error {
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

func registerUserType(schema *schemabuilder.Schema) {
	schema.Enum("Gender", model.Gender(0), map[string]model.Gender{
		model.GenderString[model.Man]:     model.Man,
		model.GenderString[model.Woman]:   model.Woman,
		model.GenderString[model.Unknown]: model.Unknown,
	})
	schema.Enum("UserState", model.UserState(0), map[string]model.UserState{
		model.UserStateString[model.Unsign]:    model.Unsign,
		model.UserStateString[model.Normal]:    model.Normal,
		model.UserStateString[model.Forbidden]: model.Forbidden,
		model.UserStateString[model.Freeze]:    model.Freeze,
	})

	user := schema.Object("User", model.User{})
	user.FieldFunc("UserFollowList", resolve.UserResolve.FollowList)
	user.FieldFunc("UserFollowerList", resolve.UserResolve.FollowerList)

	schema.Object("UserCount", model.UserCount{})
	user.FieldFunc("userCount", resolve.UserResolve.GetUserCount)

	query := schema.Query()
	schemabuilder.RelayKey(model.User{}, "id")
	query.FieldFunc("GetUserList", resolve.UserResolve.GetUserList, schemabuilder.RelayConnection)
	query.FieldFunc("GetUser", resolve.UserResolve.GetUser)
	query.FieldFunc("CurrentUser", resolve.UserResolve.CurrentUser, middleware.LoginNeed())
	query.FieldFunc("CheckUsername", resolve.UserResolve.CheckUsername, middleware.NotLogin())
	query.FieldFunc("CheckEmail", resolve.UserResolve.CheckEmail, middleware.NotLogin())

	mutation := schema.Mutation()
	mutation.FieldFunc("SignUp", resolve.UserResolve.SignUp, middleware.NotLogin())
	mutation.FieldFunc("SignIn", resolve.UserResolve.SignIn, middleware.NotLogin())
	mutation.FieldFunc("Logout", resolve.UserResolve.Logout, middleware.LoginNeed())
	mutation.FieldFunc("UpdateUser", resolve.UserResolve.Update, middleware.LoginNeed())
	mutation.FieldFunc("Follow", resolve.UserResolve.Follow, middleware.LoginNeed())
	mutation.FieldFunc("UnFollow", resolve.UserResolve.UnFollow, middleware.LoginNeed())

}

func NewService() *userService {
	once.Do(func() {
		schemabuilder.NewValidate()
		builder := schemabuilder.NewSchema()
		registerUserType(builder)
		schema := builder.MustBuild()
		introspection.AddIntrospectionToSchema(schema)
		us = &userService{schema: schema, execute: &execution.Executor{}}
	})
	return us
}
