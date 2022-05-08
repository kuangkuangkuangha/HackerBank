package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
)

func main() {
	secretKey := "c2VjcmV0X2tleTogTXV4aVN0dWRpbzIwMzMwNCwgZXJyb3JfY29kZTogZm9yIHtnbyBmdW5jKCl7dGltZS5TbGVlcCgxKnRpbWUuSG91cil9KCl9"

	decode, err := encrypt.Base64Decode(secretKey)
	if err != nil {
		return
	}

	fmt.Println(decode)
}
