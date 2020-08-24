package main

import (
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger() (*zap.Logger, error) {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			TimeKey:     "time",
			MessageKey:  "msg",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}

	return logConfig.Build()
}

func getImg() ([]byte, error) {
	return ioutil.ReadFile("resource/input/example.jpg")
}

func writeImg(img []byte) error {
	return ioutil.WriteFile("resource/output/example.jpg", img, 0644)
}

func main() {
	logger, err := createLogger()
	if err != nil {
		return
	}
	defer logger.Sync() // nolint

	logger.Info("convert started")

	logger.Info("read image")
	img, err := getImg()
	if err != nil {
		logger.Warn("failed to read image", zap.Error(err))
		return
	}

	logger.Info("write image")
	err = writeImg(img)
	if err != nil {
		logger.Warn("failed to write image", zap.Error(err))
		return
	}

	logger.Info("convert completed")
}
