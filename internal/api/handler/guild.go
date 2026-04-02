package handler

import (
	"net/http"

	"github.com/MJ-9527/GulidSys/internal/service"
	"github.com/gin-gonic/gin"
)

// CreateGuildHandler 创建工会
func CreateGuildHandler(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Leader int64  `json:"leader"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
		return
	}

	guild, err := service.CreateGuild(req.Name, req.Leader)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"id":     guild.ID,
			"name":   guild.Name,
			"leader": guild.Leader,
		},
	})
}

func JoinGuildHandler(c *gin.Context) {
	var req struct {
		GuildID int64 `json:"guild_id"`
		UserID  int64 `json:"user_id"`
	}

	// 解析JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
		return
	}

	// 参数校验
	if req.GuildID == 0 || req.UserID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "guild_id and user_id are required",
		})
		return
	}

	// 调用 service
	err := service.JoinGuild(req.GuildID, req.UserID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}
