package user

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}

func (u *UserController) Login(ctx *gin.Context) {

	var r LoginRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(200, &gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"message": "success",
	})
}
