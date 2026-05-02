package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JsonResult 统一响应结构体
type JsonResult struct {
	Code int         `json:"code"` // 业务约定：0 为成功，-1 为失败
	Msg  string      `json:"msg"`  // 错误描述或提示信息
	Data interface{} `json:"data"` // 业务数据
}

// ResultOk 业务成功 (Code: 0)
func ResultOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JsonResult{
		Code: 0,
		Msg:  "ok",
		Data: data,
	})
}

// ResultFail 业务失败 (Code: -1)
func ResultFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, JsonResult{
		Code: -1,
		Msg:  msg,
		Data: nil, // 或者返回 struct{}{}
	})
}
