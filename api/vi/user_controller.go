package vi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/go-chat-x14n/common/request"
	"github.com/x14n/go-chat-x14n/common/response"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/model"
	"github.com/x14n/go-chat-x14n/service"
	"go.uber.org/zap"
)

//todo:
//1.密码加密
//2.token认证
//3.分布式id方案

func Login(c *gin.Context) {
	var user model.User

	c.ShouldBindJSON(&user)
	global.GLogger.Info("user Login", zap.Any("user Login", user))

	if service.UserService.Login(&user) {
		c.JSON(http.StatusOK, response.SuccessWithMsgData("登录成功", user))
		return
	}

	c.JSON(http.StatusOK, response.FailWithMsg("用户名或者密码错误"))
	// TODO: JWT TOKEN
}

func Register(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	err := service.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusOK, response.FailWithMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessWithMsg("注册成功"))
}

func GetUserDetails(c *gin.Context) {
	uuid := c.Param("uuid")
	u := service.UserService.GetUserDetails(uuid)
	c.JSON(http.StatusOK, response.SuccessWithData(u))
}

func AddFriend(c *gin.Context) {
	var userFriendRequest request.FriendRequest
	c.ShouldBindJSON(&userFriendRequest)

	err := service.UserService.AddFriend(&userFriendRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailWithMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.Success())
}
