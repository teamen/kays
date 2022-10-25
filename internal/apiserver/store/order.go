package store

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type OrderStore interface {
	Create(ctx context.Context, order *v1.Order) error
}
