package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {

	request, err := httptool.NewRequest(
		"",
		"http://127.0.0.1:8080/api/v1/organization/code",
		"",
		httptool.DEFAULT)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := httptool.SendRequest(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.ShowBody()
	resp.ShowHeader()
}
