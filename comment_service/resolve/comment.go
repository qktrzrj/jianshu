package resolve

import (
	"context"
	"github.com/shyptr/hello-world-web/comment_service/model"
)

type commentResolve struct{}

var CommentResolve = commentResolve{}

func (this commentResolve) Comments(ctx context.Context, args struct {
	Aid int64 `graphql:"aid"`
}) ([]model.Comment, error) {
	return model.GetComments(ctx, args.Aid)
}

func (this commentResolve) AddComment(ctx context.Context, args struct {
	Aid     int64  `graphql:"aid"`
	Content string `graphql:"content" validate:"max=500"`
}) (model.Comment, error) {
	return model.InsertComment(ctx, map[string]interface{}{
		"uid":     ctx.Value("userId"),
		"aid":     args.Aid,
		"content": args.Content,
	})
}

func (this commentResolve) RemoveComment(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	return model.RemoveComment(ctx, args.Id)
}

func (this commentResolve) Replies(ctx context.Context, comment model.Comment) ([]model.CommentReply, error) {
	return model.GetReplies(ctx, comment.Id)
}

func (this commentResolve) Reply(ctx context.Context, args struct {
	Cid     int64  `graphql:"cid"`
	Content string `graphql:"content" validate:"max=100"`
}) (model.CommentReply, error) {
	return model.InsertReply(ctx, map[string]interface{}{
		"cid":     args.Cid,
		"content": args.Content,
	})
}

func (this commentResolve) RemoveReply(ctx context.Context, args struct {
	Id int64 `graphql:"id"`
}) error {
	return model.RemoveReply(ctx, args.Id)
}
