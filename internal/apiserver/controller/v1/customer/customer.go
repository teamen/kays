package customer

import (
	"log"

	"github.com/gin-gonic/gin"
	srvv1 "github.com/teamen/kays/internal/apiserver/service/v1"
)

type CustomerController struct {
	srv srvv1.Service
}

func (ctrl *CustomerController) Create(ctx *gin.Context) {

	var r struct {
	}
	log.Printf("%v", r)
}

func (ctrl *CustomerController) List(ctx *gin.Context) {}

func (ctrl *CustomerController) Get(ctx *gin.Context) {}

func (ctrl *CustomerController) Delete(ctx *gin.Context) {}
