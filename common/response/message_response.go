package response

import "time"

type MessageResponse struct {
	ID           int32     `json:"id"`
	FromUserId   int32     `json:"fromUserId"`
	ToUserId     int32     `json:"toUserId"`
	Content      string    `json:"content"`
	ContentType  int16     `json:"contentType"`
	CreatedAt    time.Time `json:"createAt"`
	FromUsername string    `json:"fromUsername"`
	ToUsername   string    `json:"toUsername"`
	Avatar       string    `json:"avatar"`
	Url          string    `json:"url"`
}


