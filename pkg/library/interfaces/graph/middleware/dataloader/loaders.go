package dataloader

import (
	"context"

	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

type contextKey string

const key = contextKey("dataloaders")

type Loaders struct {
	GetGroupByID *ComposedBy
	GetAlbumByID *AppearsIn
}

func newLoaders(ctx context.Context, repo domain.Repository) *Loaders {
	return &Loaders{
		GetGroupByID: newGetGroupByID(ctx, repo),
		GetAlbumByID: newGetAlbumByID(ctx, repo),
	}
}
