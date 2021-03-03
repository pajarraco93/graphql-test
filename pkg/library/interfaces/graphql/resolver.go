package graphql

import (
	"log"

	"github.com/graphql-go/graphql"

	"github.com/pajarraco93/graphql-test/pkg/library/application/usecases"
	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
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
	log.Println("Listing all groups...")

	groups, err := r.libraryService.AllGroups()
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r resolver) CreateGroup(params graphql.ResolveParams) (interface{}, error) {
	log.Println("Creating group...")

	group := entities.Group{
		Name:  "",
		Genre: "",
	}

	if name, ok := params.Args["input"].(map[string]interface{})["name"].(string); ok {
		group.Name = name
	}

	if genre, ok := params.Args["input"].(map[string]interface{})["genre"].(string); ok {
		group.Genre = genre
	}

	err := r.libraryService.CreateGroup(group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (r resolver) CreateAlbum(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r resolver) CreateSong(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
