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
					_, exit := s.Clients[msg.From]
					_, exits := s.Clients[msg.To]

					if exit && exits {

					}
				} else {

				}
				//检查clients chan中是否有user
				//检查是群聊还是什么类型
			} else {
				global.GLogger.Info("no user to send")
			}
			//case conn := <-s.Broadcast:
			//从message中查看发送给谁
			//有接受人
			//是发送到个人还是群
			// 个人
			//如果是普通文件消息的话直接转发
			//如果是文件等消息的话 先保存 再转发
			//如果client chan 中没有这个人
			//查看mysql是否有这个用户,
			//	Y: 将消息发送到kafka brooker 保存起来 等用户登录的时候再获取
			// N : 输出错误日志 断开连接
			// 群发消息
			//无接受人
			//直接中断连接
		}

	}
}

//群发消息

// 保存消息的策略
func saveMessage(msg *protocol.Message) {
	// file 类型 file类型主要是文件压缩包tar zip
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
