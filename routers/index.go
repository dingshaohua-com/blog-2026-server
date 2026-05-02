package routers

import "github.com/gin-gonic/gin"

// InitRouter 路由注册清单
func InitRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	{
		RegisterRootRoutes(api)
		RegisterUserRoutes(api)
		RegisterArticleRoutes(api)
	}
	return r
}
