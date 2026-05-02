package routers

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(g *gin.RouterGroup) {
	router := g.Group("/user") // 自动变成 /api/lark
	{
		router.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "i am user"})
		})
	}
}
