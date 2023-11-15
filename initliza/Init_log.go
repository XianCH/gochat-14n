package initliza

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/x14n/go-chat-x14n/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog(logPath string, logLevel string) {
	var jumberjacklogger = global.GConfig.Log.LumberJack

	hook := lumberjack.Logger{
		Filename:   getLogFile(),
		MaxSize:    jumberjacklogger.MaxSize,
		MaxAge:     jumberjacklogger.MaxAge,
		MaxBackups: jumberjacklogger.MaxBackups,
		Compress:   jumberjacklogger.Compress,
	}

	write := zapcore.AddSync(&hook)

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     getEncodeTime,                  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	atom := zap.NewAtomicLevel()
	atom.SetLevel(level)

	var writes = []zapcore.WriteSyncer{write}

	//如果是开发环境
	if level == zap.DebugLevel {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		level,
	)
	//栈跟踪
	caller := zap.AddCaller()
	//开启文件和行号
	development := zap.Development()
	//初始化字段
	filed := zap.Fields(zap.String("application", "chat-room"))

	//构建日志
	global.GLogger = zap.New(core, caller, development, filed)
	global.GLogger.Info("logger init success!")
}

func getLogFile() string {
	fileFormat := time.Now().Format(global.GConfig.Log.FileFormat)
	fileName := strings.Join([]string{
		global.GConfig.Log.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(global.GConfig.Log.Path, fileName)
}

func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
