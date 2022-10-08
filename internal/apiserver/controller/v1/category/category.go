package category

import (
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
	"github.com/teamen/kays/internal/apiserver/store"
)

type CategoryController struct {
	srv srvv1.Service
}

func NewCategoryController(store store.Factory) *CategoryController {
	return &CategoryController{
		srv: srvv1.NewService(store),
	}
}
