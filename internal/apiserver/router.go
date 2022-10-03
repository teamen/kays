package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/teamen/kays/internal/apiserver/controller/v1/user"
	"github.com/teamen/kays/internal/apiserver/store/mysql"
)

func loadRouter(g *gin.Engine) {

	installController(g)
}

func installController(g *gin.Engine) {

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	userController := user.NewUserController(storeIns)

	g.POST("/login", userController.Login)

	v1 := g.Group("v1")
	{
		userV1 := v1.Group("users")
		{
			userV1.POST("", userController.Create)

		}
	}

}
