package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var group = graphql.NewInterface(graphql.InterfaceConfig{
	Name:        "Group",
	Description: "A group registered in our DB",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The id of the group.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the group.",
		},
	},
},
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"groups": &graphql.Field{
			Type: group,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func NewGraphQLHandler() http.Handler {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h
}
