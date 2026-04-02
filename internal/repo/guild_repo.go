package repo

import (
	"errors"

	"github.com/MJ-9527/GulidSys/internal/model"
)

// 模拟数据库
var guilds []model.Guild
var nextGuildID int64 = 1
var guildMembers []model.GuildMember

// CreateGuild 创建新工会
func CreateGuild(name string, leaderID int64) (*model.Guild, error) {
	//检测同名工会
	for _, g := range guilds {
		if g.Name == name {
			return nil, errors.New("guild already exists")
		}
	}

	guild := &model.Guild{
		ID:     nextGuildID,
		Name:   name,
		Leader: leaderID,
	}
	nextGuildID++
	guilds = append(guilds, *guild)
	return guild, nil
}

// GetGuildByID 根据ID查找工会
func GetGuildByID(guildID int64) (*model.Guild, error) {
	for _, g := range guilds {
		if g.ID == guildID {
			return &g, nil
		}
	}
	return nil, errors.New("guild not found")
}

// AddMember 添加工会成员
func AddMember(guildID, userID int64, role string) (*model.GuildMember, error) {
	for _, m := range guildMembers {
		if m.GuildID == guildID && m.UserID == userID {
			return nil, errors.New("member already exists")
		}
	}

	member := &model.GuildMember{
		GuildID: guildID,
		UserID:  userID,
		Role:    role,
	}
	guildMembers = append(guildMembers, *member)
	return member, nil
}

// GetMembersByGuild 查询工会所有成员
func GetMembersByGuild(guildID int64) ([]model.GuildMember, error) {
	var members []model.GuildMember
	for _, m := range guildMembers {
		if m.GuildID == guildID {
			members = append(members, m)
		}
	}
	return members, nil
}
