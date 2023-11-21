package server

import (
	"encoding/base64"
	"os"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/x14n/go-chat-x14n/common/constant"
	"github.com/x14n/go-chat-x14n/common/protocol"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/service"
	"go.uber.org/zap"
)

type Server struct {
	Clients    map[string]*Client
	Mutex      *sync.Mutex
	Register   chan *Client
	UnRegister chan *Client
	Broadcast  chan []byte
}

func NewServer() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		Mutex:      &sync.Mutex{},
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

var MyServer = NewServer()

func (s *Server) start() {
	global.GLogger.Info("server start", zap.Any("server start", "server start ..."))
	for {
		select {
		case conn := <-s.Register:
			global.GLogger.Info("login", zap.Any("login", "new user login"+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "system",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.UnRegister:
			global.GLogger.Info("login", zap.Any("login out", "user:"+conn.Name+"login out"))
			if _, ok := s.Clients[conn.Name]; ok {
				delete(s.Clients, conn.Name)
				close(conn.Send)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				//判断是什么类型的消息
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDIO {
					_, exits := s.Clients[msg.From]

					if exits {
						saveMessage(msg)
					}
					if msg.MessageType == constant.MESSAGE_TYPE_TOUSER {
						client, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err != nil {
								global.GLogger.Error("proto to byte error", zap.Any("proto error", err.Error()))
								return
							}
							client.Send <- msgByte
							close(client.Send)
						}
						global.GLogger.Warn("user is not online")
					} else if msg.MessageType == constant.MESSAGE_TYPE_TOGROUP {
						sendMessageToGroup(msg, s)
					}
				} else {
					global.GLogger.Error("msgtype is not right！")
					return
				}
			} else {
				global.GLogger.Info("no user to send")
				return
			}
		}
	}
}

//群发消息

// 保存消息的策略
func saveMessage(msg *protocol.Message) {
	//file 类型 file类型主要是文件压缩包tar zip
	if msg.ContentType == 2 {
		handleBase64Context(".zip", msg)
	}
	// 图片
	if msg.ContentType == 3 {
		handleBase64Context(".png", msg)
	}
	//音频
	if msg.ContentType == 4 {
		handleBase64Context(".audio", msg)
	}
	//视屏
	if msg.ContentType == 5 {
		handleBase64Context(".vido", msg)
	}

	service.MessageService.SaveMessage(*msg)
}

func handleBase64Context(extension string, msg *protocol.Message) {
	url := uuid.New().String() + extension
	index := strings.Index(msg.Content, "base64")
	index += 7

	context := msg.Content
	context = context[index:]

	dataBuffer, dataError := base64.StdEncoding.DecodeString(context)
	if dataError != nil {
		global.GLogger.Error("transfer base64 to file error", zap.Any("transfer base64 error", dataError.Error()))
		return
	}
	err := os.WriteFile(global.StaticFilePath+url, dataBuffer, 0666)
	if err != nil {
		global.GLogger.Error("writer file error", zap.Any("write file error", err.Error()))
		return
	}

	msg.Url = url
	msg.Content = ""
}

func sendMessageToGroup(m *protocol.Message, s *Server) {
	users := service.GroupService.GetUserIdByGroup(m.To)
	for _, user := range users {
		if user.Uuid == m.From {
			continue
		}

		client, ok := s.Clients[user.Uuid]
		if !ok {
			continue
		}

		fromUserDetail := service.UserService.GetUserDetails(m.From)

		msgSend := protocol.Message{
			Avatar:       fromUserDetail.Avatar,
			FromUsername: m.FromUsername,
			From:         m.To,
			To:           m.From,
			Content:      m.Content,
			ContentType:  m.ContentType,
			Type:         m.Type,
			MessageType:  m.MessageType,
			Url:          m.Url,
		}

		byteMessage, err := proto.Marshal(&msgSend)
		if err != nil {
			client.Send <- byteMessage
		}
	}
}
