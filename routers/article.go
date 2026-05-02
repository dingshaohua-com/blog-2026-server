package routers

import (
	"blog-server/model"
	"blog-server/service"
	"blog-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var articleService = &service.ArticleService{}

// 定义一个统一的处理逻辑
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

func RegisterArticleRoutes(g *gin.RouterGroup) {
	router := g.Group("/article") // 自动变成 /api/lark
	{
		router.GET("", func(c *gin.Context) {
			res, err := articleService.GetList()
			if err != nil {
				utils.ResultFail(c, err.Error())
				return
			}
			utils.ResultOk(c, res)
		})

		router.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, _ := strconv.Atoi(idStr)
			res, err := articleService.GetOne(id)
			if err != nil {
				utils.ResultFail(c, err.Error())
				return
			}
			utils.ResultOk(c, res)
		})

		router.POST("", func(c *gin.Context) {
			handleArticleRequest(c, articleService.Save)
		})

		router.PUT("", func(c *gin.Context) {
			handleArticleRequest(c, articleService.Update)
		})

		router.DELETE("", func(c *gin.Context) {
			idStr := c.Param("id")
			id, _ := strconv.Atoi(idStr)
			err := articleService.Delete(id)
			if err != nil {
				utils.ResultFail(c, err.Error())
				return
			}
			utils.ResultOk(c, id)
		})
	}
}
