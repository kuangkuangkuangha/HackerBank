package main

import "github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"

func main() {
	request, err := httptool.NewRequest(
		"",
		"http://localhost:8080/api/v1/bank/iris_recognition_gate",
		"",
		httptool.DEFAULT)
	if err != nil {
		return
	}

	request.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiamFja3NpZSIsImlhdCI6MTYzMzMzODA3OCwibmJmIjoxNjMzMzM4MDc4fQ.L_x1S7aTnnFa3D0vpfphPr4TujRQXeiFqXGlcFLrIP8")

	res, err := httptool.SendRequest(request)
	if err != nil {
		return
	}

	res.ShowBody()
}
