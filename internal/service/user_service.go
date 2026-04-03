package service

import (
	"errors"

	"github.com/MJ-9527/GulidSys/internal/model"
	"github.com/MJ-9527/GulidSys/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser 注册用户
func RegisterUser(username, password string) (*model.User, error) {
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
	user := &model.User{
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

// Login 用户登录
func Login(username, password string) (*model.User, error) {
	// 查用户
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
