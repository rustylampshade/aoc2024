package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	dat, err := os.ReadFile("inputs/day02.txt")
	if err != nil {
		panic(err)
	}

	contents := string(dat)
	length := strings.Count(contents, "\n") + 1
	reports := make([][]int, length)

	for lineno, line := range strings.Split(contents, "\n") {
		words := strings.Split(line, " ")
		levels := make([]int, len(words))
		for i, word := range words {
			levels[i], err = strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
		}
		reports[lineno] = levels
	}

	var safe_count int = 0
	var relaxed_safe_count int = 0
	for _, levels := range reports {
		if is_increasing(&levels, 3) || is_decreasing(&levels, 3) {
			safe_count++
		} else {
			for i := range levels {
				minus_one_level := remove(levels, i)
				// fmt.Printf("Checking %v, which is index %d with chunk %v and chunk %v\n", minus_one_level, i, levels[:i], levels[i+1:])
				if is_increasing(&minus_one_level, 3) || is_decreasing(&minus_one_level, 3) {
					relaxed_safe_count++
					break
				}
			}
		}

	}
	fmt.Println(safe_count)
	fmt.Println(safe_count + relaxed_safe_count)
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func is_increasing(levels_r *[]int, max_step int) bool {
	levels := *levels_r
	for i := 1; i < len(levels); i++ {
		if levels[i]-levels[i-1] <= 0 {
			return false
		}
		if levels[i]-levels[i-1] > max_step {
			return false
		}
	}
	return true
}

func is_decreasing(levels_r *[]int, max_step int) bool {
	levels := *levels_r
	for i := 1; i < len(levels); i++ {
		if levels[i-1]-levels[i] <= 0 {
			return false
		}
		if levels[i-1]-levels[i] > max_step {
			return false
		}
	}
	return true
}
