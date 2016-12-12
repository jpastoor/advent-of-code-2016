package main

import (
	"fmt"

	"strings"

	"strconv"

	"log"

	"math"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

type Step struct {
	direction string
	amount    int
}

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(1)
	steps := parseInput(input)

	dir := 0
	x := 0
	y := 0

	secondX := 0
	secondY := 0

	visited := map[string]bool{}

	for _, step := range steps {
		// Change direction
		if step.direction == "R" {
			dir += 90
		} else {
			dir -= 90
		}

		dir = (dir + 360) % 360

		// Calculate new position
		for i := 0; i < step.amount; i++ {

			switch dir {
			case 0:
				y -= 1
			case 90:
				x += 1
			case 180:
				y += 1
			case 270:
				x -= 1
			}

			// Check if position is visited before
			if secondX == 0 && secondY == 0 {
				positionStr := strconv.Itoa(x) + "_" + strconv.Itoa(y)
				if visited[positionStr] {
					secondX = x
					secondY = y
				}
				visited[positionStr] = true
			}
		}
	}

	stepcount := int(math.Abs(float64(x)) + math.Abs(float64(y)))
	fmt.Printf("%d steps to HQ, [x:%d,y:%d]\n", stepcount, x, y)
	secondDistance := int(math.Abs(float64(secondX)) + math.Abs(float64(secondY)))
	fmt.Printf("%d steps to Second visit [x:%d,y:%d]\n", secondDistance, secondX, secondY)

}

/**
This function converts the input string to a list of steps
*/
func parseInput(input string) []Step {

	// First split by the comma
	parts := strings.Split(input, ",")

	// Reserve output memory
	output := make([]Step, len(parts))

	for i, part := range parts {
		// Split the first character off (direction) and the rest is the amount, store as Step struct
		trimmed := strings.TrimSpace(part)
		amount, err := strconv.Atoi(trimmed[1:])
		if err != nil {
			log.Fatal(err)
		}

		output[i] = Step{
			direction: trimmed[:1],
			amount:    amount,
		}
	}

	return output
}
