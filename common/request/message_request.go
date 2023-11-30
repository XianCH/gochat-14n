package request

type MessageRequest struct {
	MessageType     int32  `json:"messageType"`
	Uuid            string `json:"uuid"`
	FridendUsername string `json:"friendUsername"`
}
