package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/generated"
	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/mysql"
)

func main() {
	repo := mysql.NewMySQLRepository()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					GroupRepo: repo,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
