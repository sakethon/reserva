package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/disintegration/imaging"
)

func getImg() (image.Image, error) {
	img, err := imaging.Open("resource/input/example.jpg")

	return img, err
}

func resizeImg(img image.Image) image.Image {
	resizedImg := imaging.Resize(img, 800, 0, imaging.Lanczos)
	return resizedImg
}

func writeImg(img image.Image) {
	err := imaging.Save(img, "resource/output/example.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func main() {
	img, err := getImg()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	resizedImg := resizeImg(img)
	writeImg(resizedImg)
}
