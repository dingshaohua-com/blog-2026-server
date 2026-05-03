package routers

import (
	"blog-2026-server/model"
	"blog-2026-server/service"
	"blog-2026-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var typeService = &service.TypeService{}

// handleTypeRequest 统一处理分类的写操作（POST/PUT），封装 JSON 绑定与异常返回
func handleTypeRequest(c *gin.Context, action func(*model.Type) error) {
	var t model.Type
	if err := c.ShouldBindJSON(&t); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	if err := action(&t); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, t)
}

// GetTypeList 分页获取分类列表
// @Summary      获取分类列表
// @Description  分页获取文章分类
// @Tags         Type
// @Param        current  query  int  false  "当前页码 (默认1)"
// @Param        size     query  int  false  "每页条数 (默认10)"
// @Success      200  {object}  utils.JsonResult{data=model.PageResult[model.Type]}
// @Router       /type [get]
func GetTypeList(c *gin.Context) {
	current, _ := strconv.Atoi(c.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := typeService.GetList(current, size)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// GetTypeAll 获取所有分类（不分页）
// @Summary      获取全部分类
// @Description  通常用于前端下拉框
// @Tags         Type
// @Success      200  {object}  utils.JsonResult{data=[]model.Type}
// @Router       /type/all [get]
func GetTypeAll(c *gin.Context) {
	res, err := typeService.GetAll()
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// GetType 根据 ID 获取分类详情
// @Summary      获取分类详情
// @Tags         Type
// @Param        id   path      int  true  "分类ID"
// @Success      200  {object}  utils.JsonResult{data=model.Type}
// @Router       /type/{id} [get]
func GetType(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := typeService.GetOne(id)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// InsertType 新增分类
// @Summary      新增分类
// @Tags         Type
// @Param        type  body      model.Type  true  "分类对象"
// @Success      200   {object}  utils.JsonResult{data=model.Type}
// @Router       /type [post]
func InsertType(c *gin.Context) {
	handleTypeRequest(c, typeService.Save)
}

// UpdateType 修改分类
// @Summary      修改分类
// @Tags         Type
// @Param        type  body      model.Type  true  "分类对象"
// @Success      200   {object}  utils.JsonResult{data=model.Type}
// @Router       /type [put]
func UpdateType(c *gin.Context) {
	handleTypeRequest(c, typeService.Update)
}

// DeleteType 删除分类
// @Summary      删除分类
// @Tags         Type
// @Param        id   query     int  true  "分类ID"
// @Success      200  {object}  utils.JsonResult{data=int}
// @Router       /type [delete]
func DeleteType(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	if err := typeService.Delete(id); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, id)
}

// RegisterTypeRoutes 注册分类模块路由，对应 URL 前缀: /api/type
func RegisterTypeRoutes(g *gin.RouterGroup) {
	router := g.Group("/type")
	{
		router.GET("", GetTypeList)
		router.GET("/all", GetTypeAll)
		router.GET("/:id", GetType)
		router.POST("", InsertType)
		router.PUT("", UpdateType)
		router.DELETE("", DeleteType)
	}
}
