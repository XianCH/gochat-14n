package service

import (
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
)

type groupService struct{}

var GroupService = new(groupService)

func (g *groupService) GetUserIdByGroup(groupUUID string) []model.User {
	var group model.Group
	db := global.DB
	db.First(&group, "uuid = ?", groupUUID)
	if group.ID <= 0 {
		return nil
	}

	var users []model.User
	db.Raw("SELECT u.uuid, u.avatar, u.username FROM `groups` AS g JOIN group_members AS gm ON gm.group_id = g.id JOIN users AS u ON u.id = gm.user_id WHERE g.id = ?",
		group.ID).Scan(&users)
	return users
}
