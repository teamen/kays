package category

import (
	"github.com/gin-gonic/gin"
)

func (c *CategoryController) Create(ctx *gin.Context) {
	// core.WriteResponse(ctx, nil, gin.H{})
	// core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, "用户名必须", ""), gin.H{
	// 	"paramErrors": []string{},
	// })

	// err := errors.WithCode(code.ErrDatabase, "分类创建失败:%s", "")
	// core.WriteResponse(ctx, err, nil)
	// fmt.Printf("%#+v\n", err)

}
