package service

import (
	"fmt"

	"github.com/MJ-9527/GulidSys/internal/repo"
)

func CallAI(guildID int64, userMassage string) (string, error) {
	// 添加用户消息到缓存
	if err := repo.PushMessage(guildID, fmt.Sprintf("UserL: %s", userMassage)); err != nil {
		return "", err
	}

	// 读取最近消息作为上下文
	messages, err := repo.GetMassage(guildID)
	if err != nil {
		return "", err
	}

	context := ""
	for _, msg := range messages {
		context = context + msg + "\n"
	}

	// 调用真实AI
	aiReply := fmt.Sprintf("AI reply based on context: \n%s", context)

	// 存AI回复到
	_ = repo.PushMessage(guildID, fmt.Sprintf("AI: %s", aiReply))

	return aiReply, nil
}
