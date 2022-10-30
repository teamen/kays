package product

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/marmotedu/errors"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
	"github.com/teamen/kays/internal/apiserver/store"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"

	metav1 "github.com/teamen/kays/pkg/meta/v1"
)

type ProductController struct {
	srv srvv1.Service
}

func NewProductController(store store.Factory) *ProductController {
	return &ProductController{srv: srvv1.NewService(store)}
}

type CreateOrUpdateProductRequest struct {
	Name             string `binding:"required" json:"name"`
	Fullname         string `binding:"required" json:"fullname"`
	Code             string `binding:"required" json:"code"`
	CustomerID       uint32 `binding:"required,numeric,min=0" json:"customer_id"`
	TypeCategoryID   int64  `binding:"omitempty,numeric,min=0" json:"type_category_id"`
	SeriesCategoryID int64  `binding:"omitempty,numeric,min=0" json:"series_category_id"`
	FrameType        string `binding:"omitempty,oneof=half full less" json:"frame_type"`
	Status           int8   `binding:"omitempty,oneof=0 1" json:"status"`
	Stock            uint32 `binding:"omitempty,numeric,min=0" json:"stock"`
	MaterialFee      uint32 `binding:"omitempty,numeric,min=0" json:"material_fee"`
	ExtraFee         uint32 `binding:"omitempty,numeric,min=0" json:"extra_fee"`
}

func (ctrl *ProductController) Create(ctx *gin.Context) {

	var r CreateOrUpdateProductRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		var validationErrors validator.ValidationErrors

		fmt.Println(err.Error())

		if errors.As(err, &validationErrors) {
			out, _ := validation.ParseValidationErrors(validationErrors, r)
			core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, ""), gin.H{
				"errors": out,
			})
			return
		}

		core.WriteResponse(ctx, errors.WrapC(err, code.ErrBind, ""), nil)
		return
	}

	customer, err := ctrl.srv.Customers().Get(ctx, r.CustomerID)
	if err != nil {
		core.WriteResponse(ctx, errors.WrapC(err, code.ErrCustomerNotFound, ""), nil)
		return
	}
	fmt.Printf("%+v\n", customer)

	var product v1.Product
	if err := copier.Copy(&product, &r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	fmt.Printf("%+v\n", product)

	if err := ctrl.srv.Products().Create(ctx, &product); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, product)
	return

}

func (ctrl *ProductController) Update(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, ""), gin.H{
			"errors": "invalid param id",
		})
		return
	}

	product, err := ctrl.srv.Products().Get(ctx, uint32(id))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	fmt.Printf("%+v\n", product)

	var r CreateOrUpdateProductRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		var validationErrors validator.ValidationErrors

		fmt.Println(err.Error())

		if errors.As(err, &validationErrors) {
			out, _ := validation.ParseValidationErrors(validationErrors, r)
			core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, ""), gin.H{
				"errors": out,
			})
			return
		}

		core.WriteResponse(ctx, errors.WrapC(err, code.ErrBind, ""), nil)
		return
	}

	if err := copier.Copy(&product, &r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := ctrl.srv.Products().Update(ctx, product); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, product)
}

func (ctrl *ProductController) Get(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, ""), gin.H{
			"errors": "invalid param id",
		})
		return
	}

	product, err := ctrl.srv.Products().Get(ctx, uint32(id))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	fmt.Printf("%+v\n", product)
	core.WriteResponse(ctx, nil, product)
}

func (ctrl *ProductController) List(ctx *gin.Context) {
	var r metav1.ListOptions
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, ""), nil)
		return
	}

	fmt.Println(r)

	produts, err := ctrl.srv.Products().List(ctx, r)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{
		"produts": produts,
	})
	return
}

func (ctrl *ProductController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	_, err := ctrl.srv.Products().Get(ctx, uint32(id))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := ctrl.srv.Products().Delete(ctx, uint32(id)); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, gin.H{"id": id})
}
