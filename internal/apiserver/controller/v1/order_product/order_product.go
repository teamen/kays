package orderproduct

import (
	"github.com/gin-gonic/gin"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
)

type OrderProductController struct {
	srv *srvv1.Service
}

func (ctrl *OrderProductController) Create(ctx *gin.Context) {}

func (ctrl *OrderProductController) Update(ctx *gin.Context) {}

func (ctrl *OrderProductController) Delete(ctx *gin.Context) {

}
