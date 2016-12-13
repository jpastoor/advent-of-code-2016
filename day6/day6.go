package main

import (
	"sort"
	"strings"

	"fmt"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(6)

	data := make([][]string, 8)
	for columnIndex := 0; columnIndex < 8; columnIndex++ {
		data[columnIndex] = make([]string, 0)
	}

	for _, line := range strings.Split(input, "\n") {

		for columnIndex, ch := range line {
			data[columnIndex] = append(data[columnIndex], string(ch))
		}
	}

	for columnIndex := 0; columnIndex < 8; columnIndex++ {
		// Count occurrences per character
		occ := map[string]int{}
		for _, ch := range data[columnIndex] {
			value, _ := occ[string(ch)]
			occ[string(ch)] = value + 1
		}

		sortedPairs := rankByWordCount(occ)
		fmt.Print(sortedPairs[0].Key)
	}

}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value != p[j].Value {
		return p[i].Value < p[j].Value
	} else {
		// with ties broken by alphabetization
		return p[i].Key > p[j].Key
	}
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	// FOR A: sort.Sort(sort.Reverse(pl))
	sort.Sort(pl)
	return pl
}
