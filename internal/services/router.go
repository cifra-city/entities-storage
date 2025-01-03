package service

import (
	"context"

	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) {
	_ = chi.NewRouter()
}
