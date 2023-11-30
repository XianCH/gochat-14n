package vi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/go-chat-x14n/common/request"
	"github.com/x14n/go-chat-x14n/common/response"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/service"
	"go.uber.org/zap"
)

func GetMessage(c *gin.Context) {
	global.GLogger.Info(c.Query("uuid"))
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if err != nil {
		global.GLogger.Error("bindQueryError", zap.Any("bindQueryEroor", err))
	}
	global.GLogger.Info("messageRequest params:", zap.Any("messageRequest", messageRequest))

	message, err := service.MessageService.GetMessage(messageRequest)
	if err != nil {
		global.GLogger.Error("GetMessage error", zap.Any("GetMessage error", err))
		c.JSON(http.StatusOK, response.FailWithMsg("获取消息失败"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithMsgData("获取消息成功", message))
}
