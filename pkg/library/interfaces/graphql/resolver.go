package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/pajarraco93/graphql-test/pkg/library/application/usecases"
)

type Resolver interface {
	AllGroups(pararms graphql.ResolveParams) (interface{}, error)

	CreateGroup(pararms graphql.ResolveParams) (interface{}, error)
	CreateAlbum(pararms graphql.ResolveParams) (interface{}, error)
	CreateSong(pararms graphql.ResolveParams) (interface{}, error)
}

type resolver struct {
	libraryService usecases.UseCasesInterface
}

func NewResolver(uc usecases.UseCasesInterface) Resolver {
	return &resolver{
		libraryService: uc,
	}
}

func (r resolver) AllGroups(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r resolver) CreateGroup(pararms graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r resolver) CreateAlbum(pararms graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r resolver) CreateSong(pararms graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
