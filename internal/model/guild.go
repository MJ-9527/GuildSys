package model

type Guild struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Leader int64  `json:"leader"` //用户ID
}

type GuildMember struct {
	GuildID int64  `json:"guild_id"`
	UserID  int64  `json:"user_id"`
	Role    string `json:"role"`
}
