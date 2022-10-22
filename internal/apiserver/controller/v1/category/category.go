package category

import (
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"github.com/marmotedu/errors"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
	"github.com/teamen/kays/internal/apiserver/store"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"
)

type CategoryController struct {
	srv srvv1.Service
}

func NewCategoryController(store store.Factory) *CategoryController {
	return &CategoryController{
		srv: srvv1.NewService(store),
	}
}

type CreateCategoryRequest struct {
	ID       int64  `binding:"numeric" json:"id,omitempty" comment:"分类ID"`
	Name     string `binding:"required,max=32" json:"name,omitempty" comment:"分类名称"`
	Slug     string `binding:"required,max=64" json:"slug,omitempty" comment:"分类标识"`
	ParentID int64  `binding:"numeric" json:"parent_id,omitempty" comment:"父分类ID"`
}

func (c *CategoryController) Create(ctx *gin.Context) {

	var r CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			out, _ := validation.ParseValidationErrors(validationErrors, r)
			core.WriteResponse(ctx, errors.WrapC(err, code.ErrValidation, ""), gin.H{
				"errors": out,
			})
			return
		}

		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	core.WriteResponse(ctx, nil, r)

	return

}
