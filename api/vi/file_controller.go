package vi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/x14n/go-chat-x14n/common/response"
	"github.com/x14n/go-chat-x14n/global"
	"github.com/x14n/go-chat-x14n/service"
	"go.uber.org/zap"
)

func SaveAvatar(c *gin.Context) {
	filePrefix := uuid.New().String()
	userUuid := c.Param("uuid")
	file, _ := c.FormFile("file")
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]

	newFileName := filePrefix + suffix

	global.GLogger.Info("file", zap.Any("fiel name", global.StaticFilePath+newFileName))
	global.GLogger.Info("userUuid", zap.Any("userUuid name", userUuid))

	c.SaveUploadedFile(file, global.StaticFilePath+newFileName)
	err := service.UserService.ChangeUserAvatar(global.StaticFilePath+newFileName, userUuid)
	if err != nil {
		global.GLogger.Error("changeUserAvatar error", zap.Any("change user avatar error:", err))
		c.JSON(http.StatusOK, response.FailWithMsg(err.Error()))
	}
	c.JSON(http.StatusOK, response.Success())
}
