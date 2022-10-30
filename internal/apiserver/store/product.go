package store

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type ProductStore interface {
	Create(ctx context.Context, product *v1.Product) error
	Get(ctx context.Context, ID uint32) (*v1.Product, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProductList, error)
	Update(ctx context.Context, product *v1.Product) error
	Delete(ctx context.Context, ID uint32) error
}
