package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teamen/kays/internal/pkg/validation"
)

type CreateUserRequest struct {
	Username        string `json:"username" binding:"required" comment:"用户名"`
	Email           string `json:"email" binding:"required,email,max=255" comment:"电子邮箱"`
	Status          int    `json:"status" binding:"required,oneof=0 1" comment:"状态"`
	Password        string `json:"password" binding:"required,alphanum,min=6,eqfield=ConfirmPassword" comment:"密码"`
	ConfirmPassword string `json:"confirm_password" binding:"required" comment:"确认密码"`
}

func (u *UserController) Create(ctx *gin.Context) {

	var request CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			out, _ := validation.ParseValidationErrors(validationErrors, request)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": out,
				"errorCode":    100004,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorCode":    100003,
			"errorMessage": "参数错误",
		})
		return
	}

}
