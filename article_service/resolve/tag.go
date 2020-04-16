package resolve

import (
	"context"
	"github.com/shyptr/hello-world-web/article_service/model"
)

type tagResolve struct{}

var TagResolve = tagResolve{}

func (this tagResolve) AddTag(ctx context.Context, args struct {
	Name string `graphql:"name" validate:"max=10"`
}) error {
	return model.InsertTag(ctx, args.Name)
}

func (this tagResolve) Tag(ctx context.Context, args struct {
	Name string `graphql:"name" validate:"max=10"`
}) ([]string, error) {
	return model.GetTags(ctx, args.Name)
}
