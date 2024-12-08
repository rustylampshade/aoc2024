package day03

import (
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type instruction struct {
	startingIdx int
	endingIdx   int
	command     op
	text        string
}

type op struct {
	name    string
	pattern *regexp.Regexp
}

func Solve() {
	dat, err := os.ReadFile("inputs/day03.txt")
	if err != nil {
		panic(err)
	}
	contents := string(dat)

	ops := []op{
		{name: "mul", pattern: regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)},
		{name: "do", pattern: regexp.MustCompile(`do\(\)`)},
		{name: "don't", pattern: regexp.MustCompile(`don't\(\)`)},
	}

	instructions := []instruction{}
	for _, op := range ops {
		matches := op.pattern.FindAllStringIndex(contents, -1)
		for _, m := range matches {
			instructions = append(instructions, instruction{
				startingIdx: m[0],
				endingIdx:   m[1],
				command:     op,
				text:        contents[m[0]:m[1]],
			})
		}
	}
	slices.SortFunc(instructions, func(i, j instruction) int {
		return cmp.Compare(i.startingIdx, j.startingIdx)
	})

	var partA, partB int = 0, 0
	mulEnabled := true
	for _, instruction := range instructions {
		switch instruction.command.name {
		case "mul":
			a, b := parseMul(instruction.text)
			partA += a * b
			if mulEnabled {
				partB += a * b
			}
		case "do":
			mulEnabled = true
		case "don't":
			mulEnabled = false
		}
	}
	fmt.Println(partA)
	fmt.Println(partB)
}

func parseMul(mulstr string) (a, b int) {
	inner := mulstr[4 : len(mulstr)-1]
	chunks := strings.Split(inner, ",")
	a, err := strconv.Atoi(chunks[0])
	if err != nil {
		panic(err)
	}
	b, err = strconv.Atoi(chunks[1])
	if err != nil {
		panic(err)
	}
	return a, b
}
