package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/auth"
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

	var user v1.User
	var err error

	copier.Copy(&user, &request)
	user.Password, err = auth.Encrypt(user.Password)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode":    100005,
			"errorMessage": "加密用户密码时发生错误",
		})
		return
	}

	if err := u.srv.Users().Create(ctx, &user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode":    102001,
			"errorMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
