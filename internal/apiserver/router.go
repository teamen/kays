package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/teamen/kays/internal/apiserver/controller/v1/user"
)

func loadRouter(g *gin.Engine) {

	installController(g)
}

func installController(g *gin.Engine) {

	userController := user.NewUserController()

	g.POST("/login", userController.Login)
}
