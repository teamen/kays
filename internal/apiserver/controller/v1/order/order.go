package order

import (
	"github.com/gin-gonic/gin"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
)

type OrderController struct {
	srv *srvv1.Service
}

func (ctrl *OrderController) Create(ctx *gin.Context) {}

func (ctrl *OrderController) List(ctx *gin.Context) {}

func (ctrl *OrderController) Get(ctx *gin.Context) {}

func (ctrl *OrderController) Update(ctx *gin.Context) {}

func (ctrl *OrderController) Delete(ctx *gin.Context) {}
