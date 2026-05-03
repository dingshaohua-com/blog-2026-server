package routers

import (
	"blog-2026-server/model"
	"blog-2026-server/service"
	"blog-2026-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var moodService = &service.MoodService{}

// CreateMood 添加心情
// @Summary      添加心情
// @Tags         Mood
// @Param        data  body      model.Mood  true  "心情内容"
// @Success      200   {object}  utils.JsonResult{data=model.Mood} "返回新增的心情"
// @Router       /mood [post]
func CreateMood(c *gin.Context) {
	var mood model.Mood
	if err := c.ShouldBindJSON(&mood); err != nil {
		utils.ResultFail(c, "参数解析失败")
		return
	}
	if err := moodService.Create(&mood); err != nil {
		utils.ResultFail(c, "添加失败")
		return
	}
	utils.ResultOk(c, mood)
}

// DeleteMood 删除心情
// @Summary      删除心情
// @Tags         Mood
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  utils.JsonResult{data=int} "返回被删除的心情 ID"
// @Router       /mood/{id} [delete]
func DeleteMood(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := moodService.Delete(id); err != nil {
		utils.ResultFail(c, "删除失败")
		return
	}
	utils.ResultOk(c, id)
}

// UpdateMood 修改心情
// @Summary      修改心情
// @Tags         Mood
// @Param        data  body      model.Mood  true  "心情对象"
// @Success      200   {object}  utils.JsonResult{data=model.Mood} "返回更新后的心情"
// @Router       /mood [put]
func UpdateMood(c *gin.Context) {
	var mood model.Mood
	if err := c.ShouldBindJSON(&mood); err != nil {
		utils.ResultFail(c, "参数错误")
		return
	}
	if err := moodService.Update(mood); err != nil {
		utils.ResultFail(c, "更新失败")
		return
	}
	utils.ResultOk(c, mood)
}

// GetMoodList 获取列表
// @Summary      获取心情列表
// @Tags         Mood
// @Param        current  query  int  false  "当前页"
// @Param        size     query  int  false  "每页条数"
// @Success      200      {object}  utils.JsonResult{data=model.PageResult[model.Mood]}
// @Router       /mood [get]
func GetMoodList(c *gin.Context) {
	current, _ := strconv.Atoi(c.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := moodService.GetList(current, size)
	if err != nil {
		utils.ResultFail(c, "查询失败")
		return
	}
	utils.ResultOk(c, res)
}

func RegisterMoodRoutes(g *gin.RouterGroup) {
	router := g.Group("/mood")
	{
		router.POST("", CreateMood)       // 新增：POST /mood
		router.DELETE("/:id", DeleteMood) // 删除：DELETE /mood/1
		router.PUT("", UpdateMood)        // 修改：PUT /mood
		router.GET("", GetMoodList)       // 列表：GET /mood
	}
}
