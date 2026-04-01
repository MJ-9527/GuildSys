package handler

import (
	"net/http"

	"github.com/MJ-9527/GulidSys/internal/service"
	"github.com/gin-gonic/gin"
)

// RegisterUserHandler 用户注册接口
func RegisterUserHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 解析 JSON 请求
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "invalid request",
		})
		return
	}

	// 调用 service 层注册用户
	user, err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}
