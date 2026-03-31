package main

import (
	"log"

	"github.com/MJ-9527/GulidSys/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB("root", "123456", "127.0.0.1:3306", "guild_system")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	r := gin.Default()

	// 用户模块路由
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", registerHandler)
		userGroup.POST("/login", loginHandler)
	}

	r.Run(":8080")
}
