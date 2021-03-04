package main

import (
	"log"

	"github.com/pajarraco93/graphql-test/pkg/library/application/usecases"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/echo"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graphql"
	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/mysql"
)

func main() {
	repo := mysql.NewMySQLRepository()

	uc := usecases.NewUseCases(repo)

	graphQLHandler, err := graphql.NewGraphQL(uc)

	echoServer, err := echo.NewEcho(
		echo.WithPort(8080),
		echo.WithGraphQLServer(graphQLHandler),
	)
	if err != nil {
		log.Fatal("Error initializing echo server: ", err)
	}

	echoServer.Start()
}
