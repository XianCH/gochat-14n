package initliza

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/x14n/go-chat-x14n/global"
)

func Redis_Init() {

	if !global.GConfig.Redis.Enable {
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.GConfig.Redis.Addr,
		Password: global.GConfig.Redis.Password,
		DB:       global.GConfig.Redis.DefaultDB,
	})

	//使用超时上下文，校验redis
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), global.GConfig.Redis.DialTimeOut)
	defer cancelFunc()

	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis 初始化失败" + err.Error())
	}

	global.GRedis = redisClient
	global.GLogger.Info("redis启动成功")
}
