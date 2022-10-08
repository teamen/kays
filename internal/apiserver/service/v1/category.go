package v1

import (
	"context"

	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

type CategorySrv interface {
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

func (u *categoryService) Create(ctx context.Context, categry *v1.Category) error {
	return nil
}
