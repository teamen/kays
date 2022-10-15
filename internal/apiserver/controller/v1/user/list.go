package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"
	v1 "github.com/teamen/kays/pkg/meta/v1"
)

func (u *UserController) List(ctx *gin.Context) {

	var r v1.ListOptions
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, ""), nil)
		return
	}

	fmt.Println(r)

	users, err := u.srv.Users().List(ctx, r)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{
		"users": users,
	})
	return
}
