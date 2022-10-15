package setting

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/marmotedu/errors"

	"github.com/teamen/kays/internal/pkg/constant"
	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
	"github.com/teamen/kays/internal/pkg/validation"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"
)

type CreateSettingRequest struct {
	Name    string `json:"name" binding:"required,max=255" comment:"名称"`
	Slug    string `json:"slug" binding:"required,max=48" comment:"唯一标识"`
	Sort    int    `json:"sort,omitempty" binding:"numeric,min=0" comment:"排序"`
	Type    string `json:"type" binding:"required,oneof=number text" comment:"类别"`
	Value   string `json:"value" binding:"required,max=255" comment:"设置值"`
	Summary string `json:"summary" binding:"required,max=255" comment:"摘要"`
}

func (ctrl *SettingController) Create(ctx *gin.Context) {

	var request CreateSettingRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		var validationErrors validator.ValidationErrors

		fmt.Println(err.Error())

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

	var setting v1.Setting

	if err := copier.Copy(&setting, request); err != nil {
		core.WriteResponse(ctx, errors.WrapC(err, code.ErrInvalidCopy, ""), nil)
		return
	}

	setting.LastModifer = ctx.GetString(constant.XUsernameKey)
	setting.LastModifedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if err := ctrl.srv.Settings().Create(ctx, &setting); err != nil {
		core.WriteResponse(ctx, errors.WrapC(err, code.ErrDatabase, ""), nil)
		return
	}

	core.WriteResponse(ctx, nil, setting)
	return
}
