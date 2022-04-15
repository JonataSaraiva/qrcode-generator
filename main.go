package main

import (
"fmt"
"github.com/skip2/go-qrcode"
"os"
)

func main() {
	host := "http://domain.s3-website-us-east-1.amazonaws.com/"
	id := 1000000
	filePath := "/home/user/Documentos/QR.png"

	err := qrcode.WriteFile(fmt.Sprintf(host,id), qrcode.Low, 150, filePath)
	validate(err)
}

func validate(err error)  {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
