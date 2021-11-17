package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("hello")
	logger.Error("dsfsdf")
	//logger.With("")

	logger.With().Debug("msg info")

}
func getEncoder() zapcore.Encoder {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "line",
		NameKey:        "logger",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	})
	return encoder
}
