package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	dat, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	contents := string(dat)
	length := strings.Count(contents, "\n") + 1
	left := make([]int, length)
	right := make([]int, length)

	for lineno, line := range strings.Split(contents, "\n") {
		words := strings.Split(line, " ")

		left[lineno], err = strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}

		right[lineno], err = strconv.Atoi(words[len(words)-1])
		if err != nil {
			panic(err)
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	var total_distance int = 0
	var similarity_score int = 0
	for i := 0; i < length; i++ {
		total_distance += distance(left[i], right[i])
		similarity_score += similarity(left[i], &right)
	}

	fmt.Println(total_distance)
	fmt.Println(similarity_score)
}

// Return the absolute value of the distance between two values
func distance(a, b int) int {
	var distance int = a - b
	if distance < 0 {
		distance = distance * -1
	}
	return distance
}

// Return the similarity score, which is `a` times the number of times `a` occurs in the given slice.
func similarity(a int, right *[]int) int {
	var matches int = 0
	for _, elem := range *right {
		if elem == a {
			matches++
		}
	}
	return a * matches
}
