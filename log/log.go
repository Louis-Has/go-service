package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Logger *zap.Logger

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 记录准确的调用函数信息
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder      // 使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   "log/logs/test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	ws := io.MultiWriter(file, os.Stdout) // 支持文件和终端两个输出目标
	return zapcore.AddSync(ws)
}

func Debug(msg string, fields ...zap.Field) {
	defer Logger.Sync()
	Logger.Debug(msg, fields...)
}
func Info(msg string, fields ...zap.Field) {
	defer Logger.Sync()
	Logger.Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	defer Logger.Sync()
	Logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	defer Logger.Sync()
	Logger.Error(msg, fields...)
}
