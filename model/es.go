package model

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/shyptr/jianshu/setting"
	"log"
	"sync"
)

var (
	ESClient *elastic.Client
	esOnce   sync.Once
)

func ESInit() {
	esOnce.Do(func() {
		var err error
		ESClient, err = elastic.NewClient(elastic.SetURL(setting.GetESUrl()), elastic.SetSniff(false))
		if err != nil {
			log.Fatalf("连接elasticsearch出错:%s", err)
		}

		_, _, err = ESClient.Ping(setting.GetESUrl()).Do(context.Background())
		if err != nil {
			log.Fatalf("连接elasticsearch出错:%s", err)
		}
		exist, _ := ESClient.IndexExists("article").Do(context.Background())
		if !exist {
			ESClient.CreateIndex("article").Do(context.Background())
		}
	})
}
