package response

type ResponseMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"Msg"`
	Data any    `json:"Data"`
}

func Success() *ResponseMessage {
	return &ResponseMessage{
		Code: 0,
	}
}

func SuccessWithMsg(msg string) *ResponseMessage {
	return &ResponseMessage{
		Code: 0,
		Msg:  msg,
	}
}

func SuccessWithData(data any) *ResponseMessage {
	return &ResponseMessage{
		Code: 0,
		Msg:  "Success",
		Data: data,
	}
}

func SuccessWithMsgData(msg string, data any) *ResponseMessage {
	return &ResponseMessage{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}

func Fail() *ResponseMessage {
	return &ResponseMessage{
		Code: -1,
	}
}

func FailWithMsg(msg string) *ResponseMessage {
	return &ResponseMessage{
		Code: -1,
		Msg:  msg,
	}
}

func FailWithMsgData(msg string, data any) *ResponseMessage {
	return &ResponseMessage{
		Code: -1,
		Msg:  msg,
		Data: data,
	}
}
