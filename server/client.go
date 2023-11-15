package server

import (
	"github.com/x14n/go-chat-x14n/common/constant"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/x14n/go-chat-x14n/common/protocol"
	"github.com/x14n/go-chat-x14n/global"
	"go.uber.org/zap"
)

type Client struct {
	Name string
	Send chan []byte
	Conn *websocket.Conn
}

// client向通道中读取数据
func (client *Client) Read() {
	defer func() {
		MyServer.UnRegister <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			global.GLogger.Error("reader message error", zap.Any("client read message error", err.Error()))
			MyServer.UnRegister <- client
			client.Conn.Close()
		}
		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)

		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err2 := proto.Marshal(pong)
			if nil != err2 {
				global.GLogger.Error("client marshal message error", zap.Any("client marshal message error", err2.Error()))
			}
			client.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			MyServer.Broadcast <- message
		}
	}
}

func (client *Client) write() {
	defer client.Conn.Close()
	for message := range client.Send {
		client.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
