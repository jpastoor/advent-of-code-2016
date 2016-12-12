package aoc2016utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type AssignmentLoader struct {
}

func (*AssignmentLoader) Load(day int) string {
	wd, _ := os.Getwd()
	file, err := os.Open(wd + string(os.PathSeparator) + "src\\github.com\\jpastoor\\advent-of-code-2016\\day" + strconv.Itoa(day) + "\\input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(body)
}
