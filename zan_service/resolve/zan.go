package resolve

import (
	"context"
	"github.com/shyptr/hello-world-web/zan_service/model"
)

type zanResolve struct{}

var ZanResolve = zanResolve{}

func (this zanResolve) AddZan(ctx context.Context, args struct {
	Objtype model.ObjType `graphql:"objtype"`
	Objid   int64         `graphql:"objid"`
}) (model.Zan, error) {
	return model.InsertZan(ctx, map[string]interface{}{
		"uid":     ctx.Value("userId"),
		"objtype": args.Objtype,
		"objid":   args.Objid,
	})
}

func (this zanResolve) RemoveZan(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	return model.DeleteZan(ctx, args.Id)
}
