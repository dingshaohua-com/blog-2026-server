package routers

import (
	"blog-2026-server/service"
	"blog-2026-server/utils"

	"github.com/gin-gonic/gin"
)

var rootService = &service.RootService{}

// AppInfo 获取应用基础信息
// @Summary      获取系统信息
// @Description  获取当前后端应用的版本、环境、运行状态等基础信息
// @Tags         基础模块
// @Success      200  {object}  utils.JsonResult{data=model.AppInfo} "返回应用配置信息"
// @Router       /app-info [get]
func AppInfo(c *gin.Context) {
	res, err := rootService.GetAppInfo()
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

func RegisterRootRoutes(router *gin.RouterGroup) {
	router.GET("", func(c *gin.Context) {
		htmlContent := "<div>你好,  Galaxy Wisdom PO </div>"
		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "i am hello"})
	})
	router.GET("/app-info", AppInfo)
}
