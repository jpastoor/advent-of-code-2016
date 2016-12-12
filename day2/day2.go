package main

import (
	"strings"

	"fmt"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

type Step struct {
	direction string
	amount    int
}

func mainx() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(2)

	pos := 5
	for _, line := range strings.Split(input, "\n") {
		//fmt.Println(line)
		for _, ch := range line {

			switch string(ch) {
			case "L":
				if pos != 1 && pos != 4 && pos != 7 {
					pos -= 1
				}
			case "R":
				if pos != 3 && pos != 6 && pos != 9 {
					pos += 1
				}
			case "U":
				if pos != 1 && pos != 2 && pos != 3 {
					pos -= 3
				}
			case "D":
				if pos != 7 && pos != 8 && pos != 9 {
					pos += 3
				}
			}
		}
		fmt.Print(pos)
	}
}
