package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/teamen/kays/pkg/code"
	"github.com/teamen/kays/pkg/core"

	"github.com/teamen/kays/internal/pkg/validation"
)

func (ctrl *SettingController) Get(ctx *gin.Context) {
	slug := ctx.Request.URL.Query().Get("slug")

	if len(slug) == 0 {
		err := errors.WithCode(code.ErrValidation, "")
		core.WriteResponse(ctx, err, gin.H{
			"errors": []validation.APIError{{Param: "slug", Message: "slug should not be empty"}},
		})
		return
	}

	setting, err := ctrl.srv.Settings().Get(ctx, slug)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, setting)
	return

}
