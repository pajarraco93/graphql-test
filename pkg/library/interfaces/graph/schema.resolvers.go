package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/generated"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/model"
)

func (r *albumResolver) ComposedBy(ctx context.Context, obj *model.Album) (*model.Group, error) {
	return r.DataLoaders.Retrieve(ctx).GetGroupByID.Load(obj.ComposedBy)
}

func (r *albumResolver) Songs(ctx context.Context, obj *model.Album) ([]*model.Song, error) {
	songs, err := r.Repo.GetSongsByAlbumID(obj.ID)
	if err != nil {
		return nil, err
	}

	var gqlSongs []*model.Song
	for _, song := range songs {
		gqlSongs = append(gqlSongs, &model.Song{
			ID:        song.ID,
			Name:      song.Name,
			AppearsIn: song.AppearsIn.ID,
		})
	}

	return gqlSongs, nil
}

func (r *groupResolver) GroupInfo(ctx context.Context, obj *model.Group) (*model.GroupInfo, error) {
	info, err := r.LastFM.GetGroupInfo(entities.Group{
		Name: obj.Name,
	})
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(info), &result)
	if err != nil {
		return nil, err
	}

	info = result["artist"].(map[string]interface{})["bio"].(map[string]interface{})["content"].(string)

	return &model.GroupInfo{
		Info: &info,
	}, nil
}

func (r *groupResolver) Albums(ctx context.Context, obj *model.Group) ([]*model.Album, error) {
	albums, err := r.Repo.GetAlbumsByGroupID(obj.ID)
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
	groups, err := r.Repo.AllGroups()
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

	return gqlGroups, nil
}

func (r *queryResolver) AllAlbums(ctx context.Context) ([]*model.Album, error) {
	albums, err := r.Repo.AllAlbums()
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
	songs, err := r.Repo.AllSongs()
	if err != nil {
		return nil, err
	}

	var gqlSongs []*model.Song
	for _, song := range songs {
		gqlSongs = append(gqlSongs, &model.Song{
			ID:        song.ID,
			Name:      song.Name,
			AppearsIn: song.AppearsIn.ID,
		})
	}

	return gqlSongs, nil
}

func (r *songResolver) AppearsIn(ctx context.Context, obj *model.Song) (*model.Album, error) {
	return r.DataLoaders.Retrieve(ctx).GetAlbumByID.Load(obj.AppearsIn)
}

// Album returns generated.AlbumResolver implementation.
func (r *Resolver) Album() generated.AlbumResolver { return &albumResolver{r} }

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Song returns generated.SongResolver implementation.
func (r *Resolver) Song() generated.SongResolver { return &songResolver{r} }

type albumResolver struct{ *Resolver }
type groupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type songResolver struct{ *Resolver }
