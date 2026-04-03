package service

import (
	"errors"

	"github.com/MJ-9527/GulidSys/internal/model"
	"github.com/MJ-9527/GulidSys/internal/repo"
)

var ErrGuildNotFound = errors.New("guild not found")

// CreateGuild 创建工会
func CreateGuild(name string, leaderID int64) (*model.Guild, error) {
	guild, err := repo.CreateGuild(name, leaderID)
	if err != nil {
		return nil, err
	}

	// 2.自动将创建者加入工会(工会长)
	_, err = repo.AddMember(guild.ID, leaderID, "leader")
	if err != nil {
		return nil, errors.New("failed to add leader to the guild")
	}

	return guild, nil
}

// JoinGuild 加入工会（只有Leader/Admin可以邀请成员）
func JoinGuild(inviterID, guildID, userID int64) error {
	// 检测工会是否存在
	guild, err := repo.GetGuildByID(guildID)
	if err != nil {
		return ErrGuildNotFound
	}

	// 获取邀请人角色
	members, _ := repo.GetMembersByGuild(guild.ID)
	var inviterRole string
	for _, m := range members {
		if m.UserID == inviterID {
			inviterRole = m.Role
			break
		}
	}

	// 权限判断
	if inviterRole != model.RoleLeader && inviterRole != model.RoleAdmin {
		return errors.New("permission denied")
	}

	// 检测是否已加入工会
	for _, m := range members {
		if m.UserID == userID {
			return errors.New("already joined the guild")
		}
	}

	// 添加工会成员
	_, err = repo.AddMember(guild.ID, userID, "member")
	if err != nil {
		return err
	}

	return nil
}
