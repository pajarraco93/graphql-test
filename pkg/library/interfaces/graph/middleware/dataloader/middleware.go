package dataloader

import (
	"context"
	"net/http"

	"github.com/pajarraco93/graphql-test/pkg/library/domain"
)

func Middleware(repo domain.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(ctx, repo)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}
