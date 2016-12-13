package main

import (
	"fmt"
	"strings"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(7)

	counter := 0
	for _, line := range strings.Split(input, "\n") {
		if checkTLSSupport(line) {
			counter++
		}
	}

	fmt.Println(counter)

}

func checkTLSSupport(ip string) bool {

	// Split string in sequences
	normal := []string{}
	hypernet := []string{}

	current := ""
	for _, ch := range ip {
		str := string(ch)

		switch str {
		case "[":
			if current != "" {
				normal = append(normal, current)
				current = ""
			}
		case "]":
			if current != "" {
				hypernet = append(hypernet, current)
				current = ""
			}
		default:
			current += str
		}
	}

	// Clean buffer at the end
	if current != "" {
		normal = append(normal, current)
		current = ""
	}

	// Loop over all normal parts to see if there is an abba
	normalContainsABBA := false
	for _, input := range normal {
		if containsABBA(input) {
			normalContainsABBA = true
			break
		}
	}

	// Loop over all hypernet parts to see if there is an abba
	hypernetContainsABBA := false
	for _, input := range hypernet {
		if containsABBA(input) {
			hypernetContainsABBA = true
			break
		}
	}

	return normalContainsABBA && !hypernetContainsABBA
}

/**
Check if string contains an ABBA
*/
func containsABBA(input string) bool {
	// We loop over every char in groups of 4, changing 1 position at a time
	for i := 0; i < len(input)-3; i++ {
		// The ABBA check
		if input[i] == input[i+3] && input[i+1] == input[i+2] && input[i+2] != input[i+3] {
			return true
		}
	}

	return false
}
