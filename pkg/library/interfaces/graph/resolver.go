package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

type Resolver struct {
	Repo   domain.Repository
	LastFM domain.InfoRepo
}
