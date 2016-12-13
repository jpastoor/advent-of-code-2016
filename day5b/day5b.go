package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := "ojvtpuvg"
	output := make([]string, 8)
	found := 0
	for i := 0; found < 8; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))

		// Speeding up a bit
		if sum[0] != 0 {
			continue
		}

		hashStr := fmt.Sprintf("%x", sum)
		if hashStr[:1] == "0" && hashStr[:5] == "00000" {
			fmt.Print()
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

	end := time.Now()

	fmt.Printf("Took: %v", end.Sub(start))
}
