package order

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
	"github.com/teamen/kays/internal/apiserver/store"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"
)

type OrderController struct {
	srv srvv1.Service
}

func NewOrderController(store store.Factory) *OrderController {
	return &OrderController{srv: srvv1.NewService(store)}
}

type CreateOrderRequest struct {
	CustomerID uint32 `binding:"required,numeric,min=0" json:"customer_id,omitempty" comment:"所属客户ID"`
	Type       string `binding:"required,oneof=process repair" json:"type,omitempty" comment:"订单类型"`
}

type UpdateOrderRequest struct {
	Owner           string `binding:"max=32" json:"owner,omitempty" comment:"姓名"`
	OwnerPhone      string `binding:"max=32" json:"owner_phone,omitempty" comment:"手机"`
	Status          uint8  `binding:"required" json:"status,omitempty" comment:"状态"`
	DiscountPercent uint16 `binding:"min=0,numeric,max=10000" json:"discount_percent,omitempty" comment:"折扣"`
	AdjustFee       int32  `binding:"numeric" json:"adjust_fee,omitempty" comment:"差价"`
	Memo            string `binding:"max=255" json:"memo,omitempty" comment:"订单备注"`
}

func (ctrl *OrderController) Create(ctx *gin.Context) {
	var r CreateOrderRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, ""), nil)
		return
	}

	// TODO

	core.WriteResponse(ctx, nil, r)
	return
}

func (ctrl *OrderController) Update(ctx *gin.Context) {
	var r UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, ""), nil)
		return
	}
	id := ctx.Param("id")

	// TODO
	core.WriteResponse(ctx, nil, gin.H{
		"r":  r,
		"id": id,
	})
	return
}

func (ctrl *OrderController) List(ctx *gin.Context) {

}

func (ctrl *OrderController) Get(ctx *gin.Context) {}

func (ctrl *OrderController) Delete(ctx *gin.Context) {}
