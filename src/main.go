package main

import (
    "fmt"
    "io/ioutil"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

func getImg()([]byte, error) {
    img, err := ioutil.ReadFile("resource/input/example.jpg")

    return img, err
}

func writeImg(img []byte) {
    ioutil.WriteFile("resource/output/example.jpg", img, 0644)
}

func main() {
	img, err := getImg()
	 if err != nil {
         fmt.Fprintln(os.Stderr, err)
         return
     }
	writeImg(img)
}
