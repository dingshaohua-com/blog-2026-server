package routers

import "github.com/gin-gonic/gin"

func RegisterRootRoutes(router *gin.RouterGroup) {
	router.GET("", func(c *gin.Context) {
		htmlContent := "<div>你好,  Galaxy Wisdom PO </div>"
		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "i am hello"})
	})

	router.GET("/app-info", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "i am hello"})
	})
}
