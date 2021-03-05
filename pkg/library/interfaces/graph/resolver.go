package graph

import (
	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GroupRepo domain.Repository
}
