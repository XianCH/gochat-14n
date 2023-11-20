package service

import (
	"github.com/x14n/go-chat-x14n/common/constant"
	"github.com/x14n/go-chat-x14n/common/protocol"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
	"go.uber.org/zap"
)

type MessageService struct{}

const NULL_ID int32 = 0

func (m *MessageService) SaveMessage(msg protocol.Message) {
	db := global.DB
	var fromUser model.User

	db.Find(&fromUser, "uuid = ?", msg.From)
	if NULL_ID == fromUser.Id {
		global.GLogger.Error("Save Message not find from user", zap.Any("save Message not find from user", fromUser.Id))
		return
	}

	var toUserId int32 = 0

	if msg.MessageType == constant.MESSAGE_TYPE_TOUSER {
		var toUser model.User
		db.Find(&toUser, "uuid = ?", msg.To)
		if toUserId == toUser.Id {
			global.GLogger.Error("Save Message not find to User", zap.Any("Save Message not find to User", toUser.Id))
			return
		}
		toUserId = toUser.Id
	}

	if msg.MessageType == constant.MESSAGE_TYPE_TOGROUP {
		var group model.Group
		db.Find(&group, "uuid = ?", msg.To)
		if NULL_ID == group.ID {
			global.GLogger.Error("Save Message not find group", zap.Any("Save Message not find group", group.ID))
			return
		}
	}

	saveMessage := model.Message{
		FromUserId:  fromUser.Id,
		ToUserId:    toUserId,
		Content:     msg.Content,
		ContentType: int16(msg.ContentType),
		MessageType: int16(msg.MessageType),
		Url:         msg.Url,
	}
	db.Save(&saveMessage)
}
