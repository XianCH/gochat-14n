package initliza

func InitServer() {
	InitViper()
	InitLog("logs/chat.log", "debug")
	Redis_Init()
	InitMysql()
}
