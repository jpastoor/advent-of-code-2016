package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {

	input := "ojvtpuvg"
	output := ""

	for i := 0; len(output) < 8; i++ {
		hashStr := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if hashStr[:5] == "00000" {
			fmt.Println(hashStr)
			output += string(hashStr[5])
		}
	}

	fmt.Println(output)
}
