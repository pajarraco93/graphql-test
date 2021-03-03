package main

import (
	"net/http"

	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/graphql"
)

func main() {
	h := graphql.NewGraphQLHandler()
	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
