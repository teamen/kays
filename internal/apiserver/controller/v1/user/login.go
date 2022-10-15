package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/marmotedu/errors"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/auth"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"
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
			core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, ""), gin.H{
				"errors": out,
			})
			return
		}

		core.WriteResponse(ctx, errors.WrapC(err, code.ErrBind, ""), nil)
		return
	}

	user, err := u.srv.Users().FindByUsername(ctx, request.Username)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := auth.Compare(user.Password, request.Password); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrPasswordIncorrect, ""), nil)
		return
	}

	tokenString, err := token.Sign(int(user.ID), user.Username)
	if err != nil {

		core.WriteResponse(ctx, errors.WrapC(err, code.ErrEncrypt, ""), nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{"token": tokenString})
}
