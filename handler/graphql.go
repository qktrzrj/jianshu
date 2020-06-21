package handler

import (
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/introspection"
	"github.com/shyptr/graphql/schemabuilder"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Register(mux *http.ServeMux) {
	builder := schemabuilder.NewSchema()
	registerUser(builder)
	registerReply(builder)
	registerComment(builder)
	registerArticle(builder)
	registerLike(builder)
	registerFileHandler(builder)
	registerMsg(builder)
	schema, err := builder.Build()
	if err != nil {
		log.Fatalln(err)
	}

	introspection.AddIntrospectionToSchema(schema)

	mux.Handle("/graphiql", graphql.GraphiQLHandler("/graphql"))
	mux.Handle("/graphql", graphql.HTTPHandler(schema))
	mux.Handle("/image/", func() http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			split := strings.Split(request.RequestURI, "/")
			file, err := os.Open("static/image/" + split[len(split)-1])

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
		}
	}())
	mux.Handle("/", func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.Header.Get("Accept"), "html") {
				if strings.HasSuffix(r.RequestURI, "reset.css") {
					r.URL.Path = "/css/reset.css"
				}
				http.FileServer(http.Dir("static/jianshu/build/")).ServeHTTP(w, r)
				return
			}
			t, err := template.ParseFiles("static/jianshu/build/index.html")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			t.Execute(w, nil)
		}
	}())
}
