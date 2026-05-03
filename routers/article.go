package routers

import (
	"blog-2026-server/model"
	"blog-2026-server/service"
	"blog-2026-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 😊实例化文章服务层，供路由处理函数调用
var articleService = &service.ArticleService{}

// 😊 handleArticleRequest 统一处理文章相关的写操作请求（POST/PUT）
// 封装了 JSON 绑定、异常处理和结果返回的重复逻辑 ,action: 具体的业务处理函数（如 Save 或 Update）
func handleArticleRequest(c *gin.Context, action func(*model.Article) error) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	if err := action(&article); err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, article)
}

// GetArticleList 获取文章列表
// @Summary      获取文章列表
// @Description  分页获取文章列表，包含分类名称和是否有更多
// @Tags         Article
// @Param        current  query  int  false  "当前页码 (默认1)"
// @Param        size     query  int  false  "每页条数 (默认10)"
// @Success      200  {object}  utils.JsonResult{data=utils.PageResult[model.ArticleVO]}
// @Router       /article [get]
func GetArticleList(c *gin.Context) {
	currentStr := c.Query("current")
	sizeStr := c.Query("size")
	current, _ := strconv.Atoi(currentStr)
	size, _ := strconv.Atoi(sizeStr)

	res, err := articleService.GetList(current, size)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// GetArticle 获取单篇文章详情
// @Summary      获取文章详情
// @Description  根据文章 ID 查询文章的详细内容及分类信息
// @Tags         Article
// @Param        id   path      int  true  "文章ID"
// @Success      200  {object}  utils.JsonResult{data=model.ArticleVO}
// @Router       /article/{id} [get]
func GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	res, err := articleService.GetOne(id)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, res)
}

// InsertArticle 创建新文章
// @Summary      新增文章
// @Description  提交文章标题、内容等信息创建新篇章
// @Tags         Article
// @Param        article  body      model.Article  true  "文章内容对象"
// @Success      200      {object}  utils.JsonResult{data=model.ArticleVO} "返回创建成功的文章详情"
// @Router       /article [post]
func InsertArticle(c *gin.Context) {
	handleArticleRequest(c, articleService.Save)
}

// UpdateArticle 更新文章
// @Summary      更新文章
// @Description  根据文章对象中的 ID 更新标题、内容、分类等信息
// @Tags         Article
// @Accept       json
// @Produce      json
// @Param        article  body      model.Article  true  "需要更新的文章对象"
// @Success      200      {object}  utils.JsonResult{data=model.ArticleVO} "更新成功后的文章详情"
// @Router       /article [put]
func UpdateArticle(c *gin.Context) {
	handleArticleRequest(c, articleService.Update)
}

// DeleteArticle 删除文章
// @Summary      删除文章
// @Description  根据文章 ID 永久删除文章
// @Tags         Article
// @Param        id   query     int  true  "文章 ID"
// @Success      200  {object}  utils.JsonResult{data=int} "返回被删除的文章 ID"
// @Router       /article [delete]
func DeleteArticle(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)

	err := articleService.Delete(id)
	if err != nil {
		utils.ResultFail(c, err.Error())
		return
	}
	utils.ResultOk(c, id)
}

// RegisterArticleRoutes 注册文章模块相关的路由组，对应 URL 前缀: /api/article (基于父路由组)
func RegisterArticleRoutes(g *gin.RouterGroup) {
	// 创建 `/article` 路由分组
	router := g.Group("/article")
	{
		router.GET("", GetArticleList)
		router.GET("/:id", GetArticle)
		router.POST("", InsertArticle)
		router.PUT("", UpdateArticle)
		router.DELETE("", DeleteArticle)
	}
}
