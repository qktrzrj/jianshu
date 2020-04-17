package handler

import (
	"github.com/shyptr/graphql"
	"github.com/shyptr/graphql/introspection"
	"github.com/shyptr/graphql/schemabuilder"
	"log"
	"net/http"
)

func Register(mux *http.ServeMux) {
	builder := schemabuilder.NewSchema()
	registerUser(builder)
	schema, err := builder.Build()
	if err != nil {
		log.Fatalln(err)
	}

	introspection.AddIntrospectionToSchema(schema)

	mux.Handle("/", graphql.GraphiQLHandler("/graphql"))
	mux.Handle("/graphql", graphql.HTTPHandler(schema))
}
