package initliza

import (
	"fmt"
	"log"
	"os"

	"github.com/x14n/go-chat-x14n/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql() {
	mysqlConfig := global.GConfig.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset,
		mysqlConfig.ParseTime, mysqlConfig.TimeZone)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: mysqlConfig.Gorm.SkipDefaultTx, //是否跳过默认事务
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConfig.Gorm.TablePrefix,
			SingularTable: mysqlConfig.Gorm.SingularTable,
		},
		// 执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		PrepareStmt: mysqlConfig.Gorm.PreparedStmt,
		//在AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: mysqlConfig.Gorm.CloseForeignKey,
	}

	if global.GConfig.Mysql.Gorm.CoverLogger {
		setNewLogger(gormConfig)
	}

	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         mysqlConfig.DefaultStringSize,
		DisableDatetimePrecision:  mysqlConfig.DisableDatetimePrecision,
		SkipInitializeWithVersion: mysqlConfig.SkipInitializeWithVersion,
	}), gormConfig)

	if err != nil {
		panic(fmt.Sprintf("mysql客户端启动失败:%s", err))
	}
	global.GLogger.Info("mysql 初始化完成")
	global.DB = client
}

func setNewLogger(gConfig *gorm.Config) {
	logPath := global.GConfig.Log.Path
	file, _ := os.OpenFile(logPath+"/sql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	// 日志级别映射 error、info、warn
	logLevelMap := map[string]logger.LogLevel{
		"error": logger.Error,
		"info":  logger.Info,
		"warn":  logger.Warn,
	}
	var logLevel logger.LogLevel
	var ok bool
	if logLevel, ok = logLevelMap[global.GConfig.Mysql.LogLevel]; !ok {
		logLevel = logger.Error
	}
	newLogger := logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             global.GConfig.Mysql.SlowSql,                   //慢SQL时间
		LogLevel:                  logLevel,                                       // 记录日志级别
		IgnoreRecordNotFoundError: global.GConfig.Mysql.IgnoreRecordNotFoundError, // 是否忽略ErrRecordNotFound(未查到记录错误)
		Colorful:                  false,                                          // 开关颜色
	})
	gConfig.Logger = newLogger
}
