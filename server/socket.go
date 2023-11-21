package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/x14n/go-chat-x14n/global"
	"go.uber.org/zap"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	global.GLogger.Info("newUser", zap.Any("newUser", user))
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GLogger.Warn("upGrader error", zap.Any("upGrader error", err.Error()))
		return
	}
	client := &Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	MyServer.Register <- client
	go client.Read()
	go client.write()
}
