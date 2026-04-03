package handler

import (
	"net/http"

	"github.com/MJ-9527/GulidSys/internal/service"
	"github.com/MJ-9527/GulidSys/internal/utils"
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
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

// LoginUserHandler 用户登录
func LoginUserHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 解析请求
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "invalid request",
		})
		return
	}

	// 调用service
	user, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"token": token,
		},
	})
}
