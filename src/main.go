package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/disintegration/imaging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DEFAULT_IMAGE_SIZE = 800
)

func createS3Client(awsSession client.ConfigProvider) *s3.S3 {
	return s3.New(awsSession)
}

func getBytesFromS3(s3Svc *s3.S3, bucket string, key string) ([]byte, error) {
	resp, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

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

func getImg(imgBytes []byte) (image.Image, string, error) {
	return image.Decode(bytes.NewReader(imgBytes))
}

func resizeImg(img image.Image) image.Image {
	resizedImg := imaging.Resize(img, DEFAULT_IMAGE_SIZE, 0, imaging.Lanczos)

	return resizedImg
}

func writeImg(s3Svc *s3.S3, bucket string, key string, img image.Image) error {
	buff := new(bytes.Buffer)

	err := png.Encode(buff, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buff.Bytes()),
	})
	if err != nil {
		return err
	}
	return nil
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

	logger.Info("resize image")
	resizedImg := resizeImg(img)

	logger.Info("write image")
	err = writeImg(resizedImg)
	if err != nil {
		logger.Warn("failed to write image", zap.Error(err))
		return
	}

	logger.Info("convert completed")
}

func main() {
	lambda.Start(Handler)
}
