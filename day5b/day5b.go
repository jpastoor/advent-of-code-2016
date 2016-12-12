package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "ojvtpuvg"
	output := make([]string, 8)
	found := 0
	for i := 0; found < 8; i++ {
		hashStr := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if hashStr[:5] == "00000" {

			pos := string(hashStr[5])
			posInt, err := strconv.Atoi(pos)

			// Only allow valids
			if posInt < 8 && err == nil {
				// Skip if value is already
				if output[posInt] == "" {
					output[posInt] = string(hashStr[6])
					fmt.Println(output)
					found++
				}
			}
		}
	}

	fmt.Println(strings.Join(output, ""))
}
