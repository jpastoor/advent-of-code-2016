package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jpastoor/advent-of-code-2016/aoc2016utils"
)

var roomRegex = regexp.MustCompile(`^([\w-]+)-(\d+)\[(\w+)\]$`)

func main() {
	loader := aoc2016utils.AssignmentLoader{}
	input := loader.Load(4)

	sumSectors := 0

	for _, line := range strings.Split(input, "\n") {
		room := parseRoom(line)

		if room.isRealRoom() {
			sumSectors += room.sectorId

			if room.NameDecrypted() == "northpole object storage" {
				fmt.Printf("%s has sector %d\n", room.NameDecrypted(), room.sectorId)
			}
		}
	}

	fmt.Println(sumSectors)

}

func parseRoom(input string) Room {
	matches := roomRegex.FindStringSubmatch(input)
	sectorId, _ := strconv.Atoi(matches[2])
	return Room{
		name:     matches[1],
		sectorId: sectorId,
		checksum: matches[3],
	}
}

func (r Room) isRealRoom() bool {

	// Count occurrences per character
	occ := map[string]int{}
	for _, ch := range r.NameWithoutDashes() {
		value, _ := occ[string(ch)]
		occ[string(ch)] = value + 1
	}

	sortedPairs := rankByWordCount(occ)

	sortedPairKeys := ""
	for _, pair := range sortedPairs {
		sortedPairKeys += pair.Key
	}

	cs := sortedPairKeys[:5]

	return cs == r.checksum
}

func (r Room) NameDecrypted() string {

	newName := ""
	for _, ch := range r.name {
		if ch >= 97 && ch <= 97+26 {
			chNew := ((int(ch) + r.sectorId - 97) % 26) + 97
			newName += string(chNew)
			// Replaces dashes with spaces
		} else if string(ch) == "-" {
			newName += " "
		}

	}
	return newName
}

func (r Room) NameWithoutDashes() string {
	return strings.Replace(r.name, "-", "", -1)
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Room struct {
	name     string
	sectorId int
	checksum string
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
