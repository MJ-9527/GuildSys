package repo

import "fmt"

// PushMessage 存储消息
func PushMessage(guildID int64, message string) error {
	key := fmt.Sprintf("guild:%d:message", guildID)
	// RPUSH 心消息
	if err := Rdb.RPush(Ctx, key, message).Err(); err != nil {
		return err
	}

	// LTRIM 保留最近50条消息
	return Rdb.LTrim(Ctx, key, -50, -1).Err()
}

// GetMassage 获取消息
func GetMassage(guildID int64) ([]string, error) {
	key := fmt.Sprintf("guild:%d:massage", guildID)
	return Rdb.LRange(Ctx, key, 0, -1).Result()
}
