package main

import b64 "encoding/base64"
import "fmt"
import "os"

func main() {

    file, _ := os.Open("image.jpg")
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

    uEnc := b64.URLEncoding.EncodeToString(data)
    fmt.Println(uEnc)
}

