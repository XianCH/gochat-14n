package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type UserFriend struct {
	ID        int32                 `json:"id",gorm:"primarykey"`
	CreatedAt time.Time             `json:"createdAt"`
	DeleteAt  soft_delete.DeletedAt `json:"deletedAt"`
	UpdateAt  time.Time             `json:"updateAt"`
	UserId    int32                 `json:"userId,gorm:"index"`
	FriendId  int32                 `json:"friendId",gorm:"index"`
}
