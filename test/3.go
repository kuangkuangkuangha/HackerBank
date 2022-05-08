package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	errorCode := "for {go func(){time.Sleep(1*time.Hour)}()}"
	secretKey := "MuxiStudio203304"

	data, err := encrypt.AESEncryptOutInBase64([]byte(errorCode), []byte(secretKey))
	if err != nil {
		fmt.Println(err)
	}

	req, _ := httptool.NewRequest(
		"PUT",
		"http://localhost:8080/api/v1/bank/gate",
		string(data),
		httptool.DEFAULT)

	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiamFja3NpZSIsImlhdCI6MTYzMzMzODA3OCwibmJmIjoxNjMzMzM4MDc4fQ.L_x1S7aTnnFa3D0vpfphPr4TujRQXeiFqXGlcFLrIP8")

	res, _ := httptool.SendRequest(req)

	res.ShowBody()
}
