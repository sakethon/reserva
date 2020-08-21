package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
)

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
	img, err := getImg()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	writeImg(img)
}
