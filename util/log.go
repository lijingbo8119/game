package util

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func LogDebug(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Debug(msg, fields...)
}

func LogInfo(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Info(msg, fields...)
}

func LogPanic(msg string, fields ...zap.Field) {
	defer logger.Sync()
	logger.Panic(msg, fields...)
}

func loggerEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func loggerWriter() zapcore.WriteSyncer {
	file, err := os.Create("./test.log")
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func init() {
	core := zapcore.NewCore(loggerEncoder(), loggerWriter(), zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
}
