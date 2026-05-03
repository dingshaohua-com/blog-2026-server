package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter 路由注册清单
func InitRouter() *gin.Engine {
	r := gin.Default()
	// 解决跨域，配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		// 允许的域名，如果不追求极致安全，可以先用 AllowAllOrigins: true
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 是否允许携带 Cookie
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	{
		RegisterRootRoutes(api)
		RegisterUserRoutes(api)
		RegisterArticleRoutes(api)
		RegisterMoodRoutes(api)
	}
	return r
}
