package store

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type UserStore interface {
	Create(ctx context.Context, user *v1.User) error
	FindByUsername(ctx context.Context, username string) (*v1.User, error)
}
