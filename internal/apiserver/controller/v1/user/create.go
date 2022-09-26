package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email,max=255"`
	Status          int    `json:"status" binding:"required,oneof=0 1"`
	Password        string `json:"password" binding:"required,alphanum,min=6,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (u *UserController) Create(ctx *gin.Context) {

	var r CreateUserRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Validate pass")
}
