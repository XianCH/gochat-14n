package service

import (
	"errors"

	"github.com/x14n/go-chat-x14n/common/constant"
	"github.com/x14n/go-chat-x14n/common/protocol"
	"github.com/x14n/go-chat-x14n/common/request"
	"github.com/x14n/go-chat-x14n/common/response"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type messageService struct{}

const NULL_ID int32 = 0

var MessageService = new(messageService)

func (m *messageService) SaveMessage(msg protocol.Message) {
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

func (m *messageService) GetMessage(message request.MessageRequest) ([]response.MessageResponse, error) {
	db := global.DB
	if message.MessageType == constant.MESSAGE_TYPE_TOUSER {
		var query *model.User
		db.First(&query, "uuid = ?", message.Uuid)
		if NULL_ID == query.Id {
			return nil, errors.New("用户不存在")
		}

		var friend *model.User
		db.First(&friend, "username=?", friend.Username)
		if NULL_ID == query.Id {
			return nil, errors.New("用户不存在")
		}
		var messages []response.MessageResponse

		db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id LEFT JOIN users AS to_user ON m.to_user_id = to_user.id WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)",
			query.Id, friend.Id, query.Id, friend.Id).Scan(&messages)

		return messages, nil
	}

	if message.MessageType == constant.MESSAGE_TYPE_TOGROUP {
		msg, err := fetchGroupMessage(db, message.Uuid)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}
	return nil, errors.New("不支持查询类型")

}

func fetchGroupMessage(db *gorm.DB, toUuid string) ([]response.MessageResponse, error) {
	var query *model.Group
	db.First(&query, "uuid = ?", toUuid)
	if NULL_ID == query.ID {
		return nil, errors.New("群组不存在")
	}
	var messages []response.MessageResponse
	db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id WHERE m.message_type = 2 AND m.to_user_id = ?",
		query.ID).Scan(&messages)
	return messages, nil
}
