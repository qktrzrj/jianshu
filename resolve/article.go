package resolve

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog"
	"github.com/shyptr/graphql/schemabuilder"
	"github.com/shyptr/jianshu/model"
	"github.com/shyptr/plugins/sqlog"
	"reflect"
	"strconv"
	"time"
)

type articleResolver struct{}

var ArticleResolver articleResolver

// TODO: 每天定时维护热门文章列表，

// 将文章推送到热门
func (r articleResolver) PutHots(ctx context.Context, article model.Article) {
	// 计算当前用户量
	row := model.PSql.Select("count(id)").
		From(`"user"`).
		Where("deleted_at is null").
		RunWith(model.DB).QueryRow()
	var count float64
	row.Scan(&count)

	count = count * 0.01
	score := float64(article.ViewNum)/count + float64(article.LikeNum)/(count*0.2) + float64(article.CmtNum)/(0.05*count)
	diff := float64(time.Now().Sub(article.CreatedAt))
	score /= diff*diff + 1

	if score >= 1/(diff+1)*2 {
		_, err := model.RedisClient.ZAdd("hots", &redis.Z{
			Score:  score,
			Member: article.Id,
		}).Result()
		logger := ctx.Value("logger").(zerolog.Logger)
		if err != nil {
			logger.Error().Caller().AnErr("热门文章发送redis失败", err).Int("文章ID", article.Id).Send()
			// TODO: 发送运维提醒邮件
		}
	}
}

type ArticlePage struct {
	*schemabuilder.PaginationInfo
	Result []model.Article
}

// 查询热门文章
func (r articleResolver) Hots(ctx context.Context, arg struct {
	*schemabuilder.ConnectionArgs
}) (*ArticlePage, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	var index int64

	var (
		cursor []byte
		err    error
	)
	if arg.After != nil {
		cursor, err = base64.StdEncoding.DecodeString(*arg.After)
	} else if arg.Before != nil {
		cursor, err = base64.StdEncoding.DecodeString(*arg.Before)
	}
	if err != nil {
		logger.Error().Caller().AnErr("分页索引解码错误", err).Send()
		return nil, errors.New("查询热门文章出错")
	}
	if len(cursor) > 0 {
		cursor = bytes.TrimPrefix(cursor, []byte(schemabuilder.PREFIX))
		// 从redis获取对应id索引值
		index, err = model.RedisClient.ZRank("hots", string(cursor)).Result()
		if err != nil {
			logger.Error().Caller().AnErr("从redis获取索引失败", err).Send()
			return nil, errors.New("查询热门文章出错")
		}
	}

	var (
		ids  []string
		page ArticlePage
	)
	// 总数
	tot, err := model.RedisClient.ZCard("hots").Result()
	if err != nil {
		logger.Error().Caller().AnErr("从redis获取热门文章数量失败", err).Send()
		return nil, errors.New("查询热门文章出错")
	}
	page.PaginationInfo = &schemabuilder.PaginationInfo{}
	page.TotalCount = int(tot)

	// 往前
	if arg.First != nil {
		var i int64
		if arg.After != nil {
			i = index
		}
		ids, err = model.RedisClient.ZRevRange("hots", i, i+*arg.First).Result()
		if err != nil {
			logger.Error().Caller().AnErr("从redis获取热门文章列表失败", err).Send()
			return nil, errors.New("查询热门文章出错")
		}
		if i != 0 {
			page.HasPrevPage = true
		}
		if i != tot {
			page.HasNextPage = true
		}
	} else if arg.Last != nil { // 往后
		i := tot
		if arg.Before != nil {
			i = index
		}
		ids, err = model.RedisClient.ZRevRange("hots", i-*arg.Last, i).Result()
		if err != nil {
			logger.Error().Caller().AnErr("从redis获取热门文章列表失败", err).Send()
			return nil, errors.New("查询热门文章出错")
		}
		if i != 0 {
			page.HasPrevPage = true
		}
		if i != tot {
			page.HasNextPage = true
		}
	} else { // 全部
		ids, err = model.RedisClient.ZRevRange("hots", 0, tot).Result()
		if err != nil {
			logger.Error().Caller().AnErr("从redis获取热门文章列表失败", err).Send()
			return nil, errors.New("查询热门文章出错")
		}
	}

	page.Result = make([]model.Article, len(ids))
	for index, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Error().Caller().AnErr("解析文章id出错", err).Send()
			return nil, errors.New("查询热门文章出错")
		}

		article, err := model.QueryArticle(tx, int(id))
		if err != nil {
			logger.Error().Caller().AnErr("查询文章出错", err).Send()
			return nil, errors.New("查询热门文章出错")
		}
		page.Result[index] = article
	}

	return &page, nil
}

// 查询文章
// 根据文章名,作者名,文章内容进行查询
// 当所有查询条件为空时，默认查询最新文章
func (r articleResolver) Articles(ctx context.Context, arg struct {
	Condition string `graphql:"condition;;null"`
	Uid       int    `graphql:"uid;;null"`
}) ([]model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	var articles []model.Article
	if arg.Condition != "" {
		arg.Condition = "*" + arg.Condition + "*"
	} else {
		arg.Condition = "*"
	}
	searchResult, err := model.ESClient.Search().
		Index("article").
		Query(elastic.NewWildcardQuery("title", arg.Condition)).
		Query(elastic.NewWildcardQuery("author", arg.Condition)).
		Query(elastic.NewWildcardQuery("content", arg.Condition)).
		Sort("id", true).
		Pretty(true).
		Do(ctx)
	if err != nil {
		logger.Error().Caller().AnErr("从es查询文章失败", err).Send()
		if arg.Condition == "*" {
			arg.Condition = ""
		} else {
			arg.Condition = "%" + arg.Condition + "%"
		}
		goto DB
	}
	articles = make([]model.Article, searchResult.TotalHits())
	for index, item := range searchResult.Each(reflect.TypeOf(model.Article{})) {
		if a, ok := item.(model.Article); ok {
			a.Content = ""
			articles[index] = a
		} else {
			logger.Error().Caller().Msg("文章数据转换失败")
			goto DB
		}
	}
	goto RS
DB:
	articles, err = model.QueryArticles(tx, arg.Condition, arg.Uid, false)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章列表失败", err).Send()
		return nil, errors.New("获取文章列表失败")
	}
RS:
	return articles, nil
}

// 登录人文章
func (r articleResolver) CurArticles(ctx context.Context) ([]model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	articles, err := model.QueryArticles(tx, "", ctx.Value("userId").(int), true)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章列表失败", err).Send()
		return nil, errors.New("获取文章列表失败")
	}
	return articles, nil
}

// 获取文章详细信息
func (r articleResolver) Article(ctx context.Context, arg IdArgs) (model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	article, err := model.QueryArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章详细失败", err).Send()
		return model.Article{}, errors.New("获取文章内容失败")
	}

	return article, nil
}

// 草稿
func (r articleResolver) Draft(ctx context.Context, arg struct {
	Title string `graphql:"title"`
}) (model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	uid := ctx.Value("userId").(int)
	id, err := model.InsertArticle(tx, map[string]interface{}{
		"title":   arg.Title,
		"uid":     uid,
		"content": "",
		"state":   model.Draft,
	})
	if err != nil {
		logger.Error().Caller().AnErr("保存文章失败", err).Send()
		return model.Article{}, errors.New("草稿保存失败")
	}

	article, err := model.QueryArticle(tx, id)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章失败", err).Send()
		return model.Article{}, errors.New("草稿保存失败")
	}

	return article, nil
}

// 发布
func (r articleResolver) NewArticle(ctx context.Context, arg IdArgs) (model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	uid := ctx.Value("userId").(int)

	// 校验文章归属
	article, err := model.QueryArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章失败", err).Send()
		return model.Article{}, errors.New("文章发布失败")
	}
	if article.Uid != uid {
		return model.Article{}, errors.New("你无权发布此文章")
	}
	if article.State == model.Deleted {
		return model.Article{}, errors.New("该文章已删除，无法发布")
	}

	if article.SubTitle == "" && article.Content != "" {
		content := []rune(article.Content)
		if len(content) > 100 {
			article.SubTitle = string(content[:100]) + "..."
		} else {
			article.SubTitle = article.Content
		}
	}

	if err := model.UpdateArticle(tx, arg.Id, map[string]interface{}{
		"sub_title": article.SubTitle,
		"state":     model.Unaudited,
	}); err != nil {
		logger.Error().Caller().AnErr("修改文章失败", err).Send()
		return model.Article{}, errors.New("文章发布失败")
	}

	user, _ := UserResolver.User(ctx, IdArgs{Id: uid})
	article.Author = user.Username
	article.State = model.Unaudited

	// 存入es
	_, err = model.ESClient.Index().
		Index("article").
		Id(fmt.Sprintf("%d", article.Id)).
		BodyJson(article).
		Do(ctx)
	if err != nil {
		logger.Error().Caller().AnErr("文章存入ES失败", err).Send()
		return model.Article{}, errors.New("文章发布失败")
	}

	return article, nil
}

// 更新
func (r articleResolver) UpdateArticle(ctx context.Context, arg struct {
	Id       int     `graphql:"id"`
	Title    string  `graphql:"title;;null"`
	Cover    string  `graphql:"cover;;null"`
	SubTitle string  `graphql:"subTitle;;null"`
	Content  *string `graphql:"content"`
}) (model.Article, error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	uid := ctx.Value("userId").(int)

	// 校验文章归属
	article, err := model.QueryArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章失败", err).Send()
		return model.Article{}, errors.New("文章更新失败")
	}
	if article.Uid != uid {
		return model.Article{}, errors.New("你无权更新此文章")
	}

	state := article.State
	if article.State != model.Draft {
		state = model.Updated
	}

	setMap := map[string]interface{}{
		"state": state,
	}

	if arg.Title != "" {
		setMap["title"] = arg.Title
	}
	if arg.Cover != "" {
		setMap["cover"] = arg.Cover
	}
	if arg.SubTitle != "" {
		setMap["sub_title"] = arg.SubTitle
	}
	if arg.Content != nil {
		setMap["content"] = arg.Content
	}

	if err := model.UpdateArticle(tx, arg.Id, setMap); err != nil {
		logger.Error().Caller().AnErr("修改文章失败", err).Send()
		return model.Article{}, errors.New("修改文章失败")
	}

	article, err = model.QueryArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("查询文章失败", err).Send()
		return model.Article{}, errors.New("修改文章失败")
	}

	return article, nil
}

// 删除文章
func (r articleResolver) Delete(ctx context.Context, arg IdArgs) error {
	logger := ctx.Value("logger").(zerolog.Logger)
	tx := ctx.Value("tx").(*sqlog.DB)

	// 校验
	article, err := model.QueryArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("校验文章归属失败", err).Send()
		return errors.New("删除文章失败")
	}
	if article.Uid != ctx.Value("userId").(int) {
		return errors.New("删除文章失败,无操作权限")
	}

	idStr := fmt.Sprintf("%d", arg.Id)
	result, err := model.RedisClient.ZScore("hots", idStr).Result()
	if err != nil && err != redis.Nil {
		logger.Error().Caller().AnErr("查询redis失败", err).Send()
	}
	if result > 0 {
		_, err := model.RedisClient.ZRem("hots", idStr).Result()
		if err != nil {
			logger.Error().Caller().AnErr("删除redis热门文章失败", err).Send()
			return errors.New("删除文章失败")
		}
	}

	_, err = model.ESClient.Delete().
		Index("article").
		Id(idStr).
		Do(ctx)
	if err != nil && !elastic.IsNotFound(err) {
		logger.Error().Caller().AnErr("删除es文章数据失败", err).Send()
		return errors.New("删除文章失败")
	}

	err = model.DeleteArticle(tx, arg.Id)
	if err != nil {
		logger.Error().Caller().AnErr("删除文章数据失败", err).Send()
		return errors.New("删除文章失败")
	}
	return nil
}
