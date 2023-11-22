package initliza

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
	"github.com/x14n/go-chat-x14n/global"
)

func InitViper() {
	var configFIle string
	flag.StringVar(&configFIle, "c", global.ConfigFile, "配置文件位置")

	if len(configFIle) == 0 {
		panic("配置文件不存在")
	}
	v := viper.New()
	v.SetConfigFile(configFIle)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件出错 :%s\n", err))
	}

	if err := v.Unmarshal(&global.GConfig); err != nil {
		panic(fmt.Errorf("序列化配置文件失败:%s\n", err))
	}
	global.GConfig.App.ConfigFile = configFIle
}
