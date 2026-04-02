package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health 健康检查接口
func Health(c *gin.Context) {
	// TODO:后续可添加数据库,redis检查等
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
	})
}
