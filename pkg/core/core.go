package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/teamen/kays/pkg/code"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		coder := errors.ParseCoder(err)
		ctx.JSON(coder.HTTPStatus(), Response{
			Code:    coder.Code(),
			Message: coder.String(),
			Data:    data,
		})
		return
	}
	ctx.JSON(http.StatusOK, Response{
		Code:    code.OK,
		Message: "OK",
		Data:    data,
	})
	return
}
