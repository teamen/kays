package category

import "github.com/gin-gonic/gin"

func (c *CategoryController) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Here'is Category Ctrl Create",
	})
}
