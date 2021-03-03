package main

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"

	"github.com/pajarraco93/graphql-test/pkg/library/application/usecases"
	graphqlLibrary "github.com/pajarraco93/graphql-test/pkg/library/interfaces/graphql"
	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/mysql"
)

func main() {

	e := echo.New()

	repo := mysql.NewMySQLRepository()

	uc := usecases.NewUseCases(repo)

	schema := graphqlLibrary.NewSchema(graphqlLibrary.NewResolver(uc))
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.Query(),
		Mutation: schema.Mutation(),
	})
	if err != nil {
		log.Fatal(err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	e.GET("/graphql", echo.WrapHandler(graphQLHandler))
	e.POST("/graphql", echo.WrapHandler(graphQLHandler))

	log.Fatal(e.Start(":8080"))
}
