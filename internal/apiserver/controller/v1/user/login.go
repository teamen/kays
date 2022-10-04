package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/auth"
	"github.com/teamen/kays/pkg/token"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" comment:"用户名"`
	Password string `json:"password" binding:"required" comment:"密码"`
}

func (u *UserController) Login(ctx *gin.Context) {

	var request LoginRequest

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

	user, err := u.srv.Users().FindByUsername(ctx, request.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode":    "110001",
			"errorMessage": "用户不存在",
		})
		return
	}

	if err := auth.Compare(user.Password, request.Password); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode":    "110002",
			"errorMessage": "密码错误",
		})
		return
	}

	tokenString, err := token.Sign(int(user.ID), user.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode":    "110003",
			"errorMessage": "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errorCode":    "0",
		"errorMessage": "OK",
		"data": gin.H{
			"token": tokenString,
		},
	})
}
