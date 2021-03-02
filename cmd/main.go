package main

import (
	"log"
	"net/http"

	"github.com/pajarraco93/graphql-test/src/shared/infra/adapters/graphql"
	"github.com/pajarraco93/graphql-test/src/shared/infra/adapters/mongo"
)

func main() {
	_, err := mongo.NewMongoRepository()
	if err != nil {
		log.Fatal("Error with MongoDB: ", err)
	}
	//defer mongoDB.disconnect()

	h := graphql.NewGraphQLHandler()
	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
