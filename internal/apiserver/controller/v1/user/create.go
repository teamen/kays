package user

import (
	// "errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/auth"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"

	"github.com/marmotedu/errors"
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
			core.WriteResponse(ctx, errors.WrapC(err, code.ErrValidation, ""), gin.H{
				"errors": out,
			})
			return
		}

		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, ""), nil)
		return
	}

	var user v1.User
	var err error

	copier.Copy(&user, &request)
	user.Password, err = auth.Encrypt(user.Password)

	if err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrEncrypt, ""), nil)
		return
	}

	if err := u.srv.Users().Create(ctx, &user); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, user)
	return

}
