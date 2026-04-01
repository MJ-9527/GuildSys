package repo

import "errors"

// User 模拟数据库模型
type User struct {
	ID       int64
	Username string
	Password string
	Role     string
}

// 模拟数据库
var users = []*User{}
var nextID int64 = 1

// CreateUser 添加新用户到数据库
func CreateUser(user *User) error {
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

func GetUserByUsername(username string) (*User, error) {
	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
