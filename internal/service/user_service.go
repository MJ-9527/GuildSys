package service

import (
	"errors"

	"github.com/MJ-9527/GulidSys/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) (*repo.User, error) {
	// 检查用户是否存在
	if _, err := repo.GetUserByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}

	//密码hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	//创建用户对象
	user := &repo.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     "member", //默认普通成员
	}

	// 保存到repo
	if err := repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
