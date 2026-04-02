package repo

import (
	"errors"

	"github.com/MJ-9527/GulidSys/internal/model"
)

// User 模拟数据库模型

// 模拟数据库
var users []*model.User
var nextID int64 = 1

// CreateUser 添加新用户到数据库
func CreateUser(user *model.User) error {
	//检查用户名是否已存在
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("username already exists")
		}
	}
	user.ID = nextID
	nextID++
	users = append(users, user)
	return nil
}

func GetUserByUsername(username string) (*model.User, error) {
	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
