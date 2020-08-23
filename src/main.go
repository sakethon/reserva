package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

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
	img, err := ioutil.ReadFile("resource/input/example.jpg")

	return img, err
}

func writeImg(img []byte) {
	err := ioutil.WriteFile("resource/output/example.jpg", img, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func main() {
	logger, err := createLogger()
	if err != nil {
		return
	}

	img, err := getImg()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	writeImg(img)
}
