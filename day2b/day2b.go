package main

import (
	"strings"

	"fmt"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(2)

	keypad := [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	x := 0
	y := 2

	// Split input by line
	for _, line := range strings.Split(input, "\n") {
		// Split string in characters
		for _, ch := range line {

			switch string(ch) {
			case "L":
				if x != 0 && keypad[y][x-1] != "" {
					x--
				}

			case "R":
				if x != 4 && keypad[y][x+1] != "" {
					x++
				}
			case "U":
				if y != 0 && keypad[y-1][x] != "" {
					y--
				}
			case "D":
				if y != 4 && keypad[y+1][x] != "" {
					y++
				}
			}
		}

		fmt.Print(keypad[y][x])
	}
}
