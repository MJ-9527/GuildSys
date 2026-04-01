package api

import (
	"github.com/MJ-9527/GulidSys/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// /health 健康检测路由
	r.GET("/health", handler.Health)

	// /user/register 用户注册路由
	r.POST("/user/register", handler.RegisterUserHandler)
	return r
}
