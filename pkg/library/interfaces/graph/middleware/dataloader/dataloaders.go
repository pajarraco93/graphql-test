package dataloader

import (
	"context"
	"time"

	"github.com/pajarraco93/graphql-test/pkg/library/domain"
	"github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/model"
	graph "github.com/pajarraco93/graphql-test/pkg/library/interfaces/graph/model"
)

type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

func NewRetriever() Retriever {
	return &retriever{key: key}
}

func newGetGroupByID(ctx context.Context, repo domain.Repository) *ComposedBy {
	return NewComposedBy(ComposedByConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(groupIDs []int) ([]*graph.Group, []error) {
			res, err := repo.GetGroupsByIDs(groupIDs)
			if err != nil {
				return nil, []error{err}
			}

			groupsByIds := make(map[int]*graph.Group, len(groupIDs))
			for _, r := range res {
				groupsByIds[r.ID] = &model.Group{
					ID:    r.ID,
					Name:  r.Name,
					Genre: &r.Genre,
				}
			}

			result := make([]*graph.Group, len(groupIDs))
			for i, groupID := range groupIDs {
				result[i] = groupsByIds[groupID]
			}

			return result, nil
		},
	})
}

func newGetAlbumByID(ctx context.Context, repo domain.Repository) *AppearsIn {
	return NewAppearsIn(AppearsInConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(albumIDs []int) ([]*graph.Album, []error) {
			res, err := repo.GetAlbumsByIDs(albumIDs)
			if err != nil {
				return nil, []error{err}
			}

			albumsByIds := make(map[int]*graph.Album, len(albumIDs))
			for _, r := range res {
				albumsByIds[r.ID] = &model.Album{
					ID:         r.ID,
					Name:       r.Name,
					Year:       &r.Year,
					ComposedBy: r.ComposedBy.ID,
				}
			}

			result := make([]*graph.Album, len(albumIDs))
			for i, albumID := range albumIDs {
				result[i] = albumsByIds[albumID]
			}

			return result, nil
		},
	})
}
