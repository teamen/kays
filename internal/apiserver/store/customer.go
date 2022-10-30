package store

import (
	"context"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type CustomerStore interface {
	Create(ctx context.Context, customer *v1.Customer) error
	Get(ctx context.Context, ID uint32) (*v1.Customer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CustomerList, error)
	Update(ctx context.Context, customer *v1.Customer) error
	Delete(ctx context.Context, ID uint32) error
}
