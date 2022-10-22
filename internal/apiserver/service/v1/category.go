package v1

import (
	"context"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type CategorySrv interface {
	Create(ctx context.Context, categry *v1.Category) error
	Get(ctx context.Context, ID int64) (*v1.Category, error)
	GetBySlug(ctx context.Context, slug string) (category *v1.Category, err error)
}

type categoryService struct {
	store store.Factory
}

var _ CategorySrv = (*categoryService)(nil)

func newCategories(srv *service) *categoryService {
	return &categoryService{
		store: srv.store,
	}
}

func (c *categoryService) Create(ctx context.Context, categry *v1.Category) error {
	// c.store.Categories()
	return nil
}

func (c *categoryService) Get(ctx context.Context, ID int64) (category *v1.Category, err error) {
	return nil, nil
}

func (c *categoryService) GetBySlug(ctx context.Context, slug string) (category *v1.Category, err error) {
	return nil, nil

}
