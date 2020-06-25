package logger

import (
	"demo_api/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var sugarLogger *zap.SugaredLogger

// func InitLogger
func InitLogger() {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Cfg.LogFileLocation,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     1, //days
		}), zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	logger := zap.New(core)

	//logger := zap.New(core, zap.AddCaller())
	//zap.ReplaceGlobals(logger)
	sugarLogger = logger.Sugar()
}

// func SyncLogger
func SyncLogger() {
	defer sugarLogger.Sync()
}

// func Info
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// func Infof
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

// func Debug
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// func Debugf
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

// func Warn
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// func Warnf
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

// func Error
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// func Errorf
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

// func Panic
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

// func Panicf
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

// func Fatal
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

// func Fatalf
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
