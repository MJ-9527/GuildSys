package api

import (
	"github.com/MJ-9527/GulidSys/internal/api/handler"
	"github.com/MJ-9527/GulidSys/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// /health 健康检测路由
	r.GET("/health", handler.Health)

	// 用户模块
	r.POST("/user/register", handler.RegisterUserHandler)
	r.POST("/user/login", handler.LoginUserHandler)

	// 工会模块
	guild := r.Group("/guild")
	guild.Use(middleware.JETAuthMiddleware())
	{
		guild.POST("/create", handler.CreateGuildHandler)
		guild.POST("/join", handler.JoinGuildHandler)
		// 后续可以添加踢人、解散等权限
	}

	return r
}
