package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/pajarraco93/graphql-test/pkg/library/application/usecases"
)

func NewGraphQL(uc usecases.UseCasesInterface) (*handler.Handler, error) {
	schema := NewSchema(NewResolver(uc))
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.Query(),
		Mutation: schema.Mutation(),
	})
	if err != nil {
		return nil, err
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: false,
		Pretty:   true,
	})

	return graphQLHandler, nil
}
