package main

import (
    "fmt"
    "io/ioutil"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

func getImg() {
    img, err := ioutil.ReadFile("resource/input/example.jpg")
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
    fmt.Printf("File contents: %s", img)
}

func main() {
	getImg()
}
