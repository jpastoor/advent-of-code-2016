package main

import (
	"strings"

	"strconv"

	"sort"

	"fmt"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(3)

	valid1 := 0

	// Split input by line
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Extract values from line
		a := cutVal(line, 0, 5)
		b := cutVal(line, 5, 10)
		c := cutVal(line, 10, 15)

		sides := []int{a, b, c}
		if isValid(sides) {
			valid1++
		}
	}

	fmt.Println(valid1)

	valid2 := 0

	// Split lines in batches of 3
	for i := 0; i < len(lines); i += 3 {
		batch := lines[i:min(i+3, len(lines))]

		// Extract values from line
		a1 := cutVal(batch[0], 0, 5)
		b1 := cutVal(batch[0], 5, 10)
		c1 := cutVal(batch[0], 10, 15)

		a2 := cutVal(batch[1], 0, 5)
		b2 := cutVal(batch[1], 5, 10)
		c2 := cutVal(batch[1], 10, 15)

		a3 := cutVal(batch[2], 0, 5)
		b3 := cutVal(batch[2], 5, 10)
		c3 := cutVal(batch[2], 10, 15)

		sidesA := []int{a1, a2, a3}
		if isValid(sidesA) {
			valid2++
		}

		sidesB := []int{b1, b2, b3}
		if isValid(sidesB) {
			valid2++
		}

		sidesC := []int{c1, c2, c3}
		if isValid(sidesC) {
			valid2++
		}
	}

	fmt.Println(valid2)
}

func cutVal(line string, start, end int) int {
	val, _ := strconv.Atoi(strings.TrimSpace(line[start:end]))
	return val
}

func isValid(sides []int) bool {

	// Sort on size
	sort.Ints(sides)

	// In a valid triangle, the sum of any two sides must be larger than the remaining side.
	return sides[0]+sides[1] > sides[2]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
