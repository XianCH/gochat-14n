package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/x14n/go-chat-x14n/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ConfigFile     = "../config.yml"
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHHIISS = "2006-01-02 15:04:05"
	StaticFilePath = "web/static/file/"
)

var (
	GLogger *zap.Logger
	DB      *gorm.DB
	GConfig config.ServerConfig
	GRedis  *redis.Client
)
