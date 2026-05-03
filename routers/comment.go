package routers

import (
	"blog-2026-server/model"
	"blog-2026-server/service"
	"blog-2026-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 实例化评论服务
var commentService = &service.CommentService{}

// handleCommentRequest 统一处理评论的写操作（POST/PUT），封装 JSON 绑定与异常返回
func handleCommentRequest(c *gin.Context, action func(*model.Comment) error) {
	var cm model.Comment
	if err := c.ShouldBindJSON(&cm); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	if err := action(&cm); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, cm)
}

// GetCommentList 分页获取某篇文章下的评论（含子评论）
// @Summary      获取评论列表
// @Description  按文章 ID 分页获取根评论，子评论会通过 children 字段一起返回
// @Tags         Comment
// @Param        articleId  query  int  true   "文章ID"
// @Param        current    query  int  false  "当前页码 (默认1)"
// @Param        size       query  int  false  "每页条数 (默认10)"
// @Success      200  {object}  utils.JsonResult{data=model.PageResult[model.Comment]}
// @Router       /comment [get]
func GetCommentList(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Query("articleId"))
	current, _ := strconv.Atoi(c.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := commentService.GetListByArticle(articleId, current, size)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// GetComment 根据 ID 获取单条评论
// @Summary      获取评论详情
// @Tags         Comment
// @Param        id   path      int  true  "评论ID"
// @Success      200  {object}  utils.JsonResult{data=model.Comment}
// @Router       /comment/{id} [get]
func GetComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := commentService.GetOne(id)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// InsertComment 新增评论（也可作为回复使用，传 replyCmId 即可）
// @Summary      新增评论
// @Description  根评论不传 replyCmId；回复某条评论时传上对应的 replyCmId
// @Tags         Comment
// @Param        comment  body      model.Comment  true  "评论对象"
// @Success      200      {object}  utils.JsonResult{data=model.Comment}
// @Router       /comment [post]
func InsertComment(c *gin.Context) {
	handleCommentRequest(c, commentService.Save)
}

// UpdateComment 更新评论
// @Summary      更新评论
// @Tags         Comment
// @Param        comment  body      model.Comment  true  "评论对象"
// @Success      200      {object}  utils.JsonResult{data=model.Comment}
// @Router       /comment [put]
func UpdateComment(c *gin.Context) {
	handleCommentRequest(c, commentService.Update)
}

// DeleteComment 删除评论（连同子评论一并删除）
// @Summary      删除评论
// @Tags         Comment
// @Param        id   query     int  true  "评论ID"
// @Success      200  {object}  utils.JsonResult{data=int}
// @Router       /comment [delete]
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := commentService.Delete(id); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, id)
}

// RegisterCommentRoutes 注册评论模块路由，对应 URL 前缀: /api/comment
func RegisterCommentRoutes(g *gin.RouterGroup) {
	router := g.Group("/comment")
	{
		router.GET("", GetCommentList)
		router.GET("/:id", GetComment)
		router.POST("", InsertComment)
		router.PUT("", UpdateComment)
		router.DELETE("", DeleteComment)
	}
}
