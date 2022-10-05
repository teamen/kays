package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/teamen/kays/internal/apiserver/controller/v1/user"
	"github.com/teamen/kays/internal/apiserver/store/mysql"
	"github.com/teamen/kays/internal/pkg/constant"
	"github.com/teamen/kays/pkg/token"
)

func loadRouter(g *gin.Engine) {

	installController(g)
}

func installController(g *gin.Engine) {

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	userController := user.NewUserController(storeIns)

	g.POST("/login", userController.Login)

	v1 := g.Group("v1")
	v1.Use(authMiddleware())
	{
		userV1 := v1.Group("users")
		{
			userV1.POST("", userController.Create)

		}
	}

}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, username, err := token.ParseRequest(ctx)

		if err != nil {
			ctx.JSON(403, gin.H{
				"errorCode":    "100403",
				"errorMessage": "未授权",
			})
			ctx.Abort()
			return
		}

		ctx.Set(constant.XUserIDKey, id)
		ctx.Set(constant.XUsernameKey, username)
		ctx.Next()
	}
}
