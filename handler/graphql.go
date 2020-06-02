package handler

import (
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/introspection"
	"github.com/shyptr/graphql/schemabuilder"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Register(mux *http.ServeMux) {
	builder := schemabuilder.NewSchema()
	registerUser(builder)
	registerArticle(builder)
	registerFileHandler(builder)
	schema, err := builder.Build()
	if err != nil {
		log.Fatalln(err)
	}

	introspection.AddIntrospectionToSchema(schema)

	mux.Handle("/", graphql.GraphiQLHandler("/graphql"))
	mux.Handle("/graphql", graphql.HTTPHandler(schema))

	mux.HandleFunc("/image/", func(writer http.ResponseWriter, request *http.Request) {
		split := strings.Split(request.RequestURI, "/")
		file, err := os.Open("./static/image/" + split[len(split)-1])

		defer writer.WriteHeader(200)
		if err != nil {
			writer.Write([]byte("文件下载失败"))
			return
		}
		defer file.Close()

		writer.Header().Add("Content-Type", "application/octet-stream")
		writer.Header().Add("content-disposition", "attachment; filename=\""+split[len(split)-1]+"\"")
		_, err = io.Copy(writer, file)
		if err != nil {
			writer.Write([]byte("文件下载失败"))
		}
	})
}
