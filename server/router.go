package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/go-chat-x14n/api/vi"
	"github.com/x14n/go-chat-x14n/global"
	"go.uber.org/zap"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(Core())
	server.Use(Recovery)

	socket := RunSocket

	group := server.Group("")
	{
		group.GET("/user", vi.Login)
		group.PUT("/register", vi.Register)
		group.PUT("/changeAvatar")
		group.GET("/socket.io", socket)
	}
	return server
}

func Core() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTISONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				global.GLogger.Error("httpError", zap.Any("httpError", err))
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			global.GLogger.Error("gin catch error", zap.Any("gin catch error", r))
			c.JSON(http.StatusOK, "系统内部出错！")
		}
	}()
	c.Next()
}
