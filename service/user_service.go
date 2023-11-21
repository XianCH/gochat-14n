package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
	"go.uber.org/zap"
)

type userService struct{}

var UserService = new(userService)

func (u *userService) GetUserDetails(uuid string) model.User {
	var queryUser *model.User
	db := global.DB
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "uuid = ?", uuid)
	return *queryUser
}

func (u *userService) Login(user *model.User) bool {
	db := global.DB
	var queryUser *model.User

	db.First(&queryUser, "username = ?", user.Username)
	global.GLogger.Debug("queryUser", zap.Any("queryUser", user))

	user.Uuid = queryUser.Uuid

	return user.Password == queryUser.Password
}

func (u *userService) Register(user *model.User) error {
	db := global.DB
	var userCount int64
	db.Model(user).Where("username", user.Username).Count(&userCount)
	if userCount > 0 {
		return errors.New("user already exist")
	}
	user.Uuid = uuid.New().String()
	db.Create(&user)
	return nil
}
