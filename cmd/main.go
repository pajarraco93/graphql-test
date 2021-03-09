package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/joho/godotenv"

	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/generated"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/middleware/dataloader"
	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/lastfm"
	"github.com/pajarraco93/graphql-test/pkg/library/shared/infra/adapters/mysql"
)

func main() {
	err := godotenv.Load("./dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	repo := mysql.NewMySQLRepository()
	lfm := lastfm.NewLastFMAPI(
		os.Getenv("APIKEY"),
	)

	dl := dataloader.NewRetriever()
	dlMiddleware := dataloader.Middleware(repo)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Repo:        repo,
					LastFM:      lfm,
					DataLoaders: dl,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dlMiddleware(srv))

	color.Green("connect to http://localhost:%s/ for GraphQL playground", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
