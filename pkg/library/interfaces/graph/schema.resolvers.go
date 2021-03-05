package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/generated"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/model"
)

func (r *albumResolver) ComposedBy(ctx context.Context, obj *model.Album) (*model.Group, error) {
	group, err := r.GroupRepo.GetGroupByID(obj.ComposedBy)
	if err != nil {
		return nil, err
	}

	gqlGroup := &model.Group{
		ID:    group.ID,
		Name:  group.Name,
		Genre: &group.Genre,
	}

	return gqlGroup, nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAlbum(ctx context.Context, input model.NewAlbum) (*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSong(ctx context.Context, input model.NewSong) (*model.Song, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllGroups(ctx context.Context) ([]*model.Group, error) {
	groups, err := r.GroupRepo.AllGroups()
	if err != nil {
		return nil, err
	}

	var gqlGroups []*model.Group
	for _, group := range groups {
		gqlGroups = append(gqlGroups, &model.Group{
			ID:    group.ID,
			Name:  group.Name,
			Genre: &group.Genre,
		})
	}

	return nil, nil
}

func (r *queryResolver) AllAlbums(ctx context.Context) ([]*model.Album, error) {
	albums, err := r.GroupRepo.AllAlbums()
	if err != nil {
		return nil, err
	}

	var gqlAlbums []*model.Album
	for _, album := range albums {
		gqlAlbums = append(gqlAlbums, &model.Album{
			ID:         album.ID,
			Name:       album.Name,
			ComposedBy: album.ComposedBy.ID,
			Year:       &album.Year,
		})
	}

	return gqlAlbums, nil
}

func (r *queryResolver) AllSongs(ctx context.Context) ([]*model.Song, error) {
	panic(fmt.Errorf("not implemented"))
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type albumResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
