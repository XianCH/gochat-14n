package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/x14n/go-chat-x14n/common/request"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
	"go.uber.org/zap"
)

type userService struct{}

var UserService = new(userService)

// 获取用户信息
func (u *userService) GetUserDetails(uuid string) model.User {
	var queryUser *model.User
	db := global.DB
	db.Select("uuid", "username", "nickname", "avatar").First(&queryUser, "uuid = ?", uuid)
	return *queryUser
}

// 登录service
func (u *userService) Login(user *model.User) bool {
	db := global.DB
	var queryUser *model.User

	db.First(&queryUser, "username = ?", user.Username)
	global.GLogger.Debug("queryUser", zap.Any("queryUser", user))

	user.Uuid = queryUser.Uuid

	return user.Password == queryUser.Password
}

// 注册service
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

func (u *userService) ChangeUserAvatar(avatar string, uuid string) error {
	var queryUser *model.User
	db := global.DB
	db.First(&queryUser, "uuid = ?", uuid)
	if NULL_ID == queryUser.Id {
		return errors.New("用户不存在")
	}
	db.Model(&queryUser).Update("avatar", avatar)
	return nil
}

// todo :好友是否同意申请
func (u *userService) AddFriend(userFriendRequest *request.FriendRequest) error {
	var queryUser *model.User
	db := global.DB
	db.First(&queryUser, "uuid", userFriendRequest.Uuid)
	if NULL_ID == queryUser.Id {
		return errors.New("本用户不存在")
	}

	var friend *model.User
	db.First(&friend, "username = ?", userFriendRequest.FriendName)
	if NULL_ID == friend.Id {
		return errors.New("friend用户不存在")
	}
	userFriend := model.UserFriend{
		UserId:   queryUser.Id,
		FriendId: friend.Id,
	}

	var userFriendQuery *model.UserFriend
	db.First(&userFriendQuery, "user_id = ? and firend_id =?", queryUser.Id, friend.Id)
	if NULL_ID != userFriendQuery.ID {
		return errors.New("用户已经是你好友")
	}

	db.AutoMigrate(&userFriend)
	db.Save(&userFriend)
	global.GLogger.Info("userFriend", zap.Any("userFriend", userFriend))
	return nil
}
