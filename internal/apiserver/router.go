package apiserver

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/teamen/kays/internal/apiserver/controller/v1/category"
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
	categoryController := category.NewCategoryController(storeIns)

	g.POST("/login", userController.Login)

	v1 := g.Group("v1")

	e := initializeCasbin()

	// e.AddRoleForUser("Wayne Tse", "super-admin")

	v1.Use(authMiddleware())
	{
		userV1 := v1.Group("users").Use(permissionMiddleware(e))
		{
			userV1.POST("", userController.Create)

		}

		categoryV1 := v1.Group("categories")
		{
			categoryV1.POST("", categoryController.Create)
		}
	}

}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, username, err := token.ParseRequest(ctx)

		if err != nil {
			ctx.JSON(401, gin.H{
				"errorCode":    "100401",
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

func permissionMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		username, exists := ctx.Get(constant.XUsernameKey)

		if !exists {
			ctx.JSON(401, gin.H{
				"errorCode":    "100401",
				"errorMessage": "未授权",
			})
			ctx.Abort()
			return
		}

		path := ctx.Request.URL.RequestURI()
		action := ctx.Request.Method
		hasPermission, err := enforcer.Enforce(username, path, action)
		if err != nil {
			ctx.JSON(500, gin.H{
				"errorCode":    "100500",
				"errorMessage": fmt.Sprintf("权限校验出错: %s", err.Error()),
			})
			ctx.Abort()
			return
		}

		if !hasPermission {
			ctx.JSON(403, gin.H{
				"errorCode":    "100403",
				"errorMessage": "无权限",
			})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
