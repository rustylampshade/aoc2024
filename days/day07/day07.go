package day07

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	dat, err := os.ReadFile("inputs/day07.txt")
	if err != nil {
		panic(err)
	}
	contents := string(dat)
	calibration_result := 0
	for _, line := range strings.Split(contents, "\n") {
		words := strings.Split(line, " ")
		target, err := strconv.Atoi(strings.Trim(words[0], ":"))
		if err != nil {
			panic(err)
		}
		numbers := []int{}
		for _, numberString := range words[1:] {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}

		fmt.Printf("%v -> %v\n", target, numbers)
		if solveable(target, numbers) {
			calibration_result += target
		}
	}
	fmt.Println(calibration_result)
}

func solveable(target int, numbers []int) bool {
	for mulMask := 0; mulMask < int(math.Pow(2, float64(len(numbers)-1))); mulMask++ {
		for catMask := 0; catMask < int(math.Pow(2, float64(len(numbers)-1))); catMask++ {

			subtotal := numbers[0]
			s := fmt.Sprintf("%d", numbers[0])
			for gap := 0; gap <= len(numbers)-2; gap++ {
				if mulMask&(1<<gap) != 0 {
					if catMask&(1<<gap) != 0 {
						s += fmt.Sprintf(" | %d", numbers[gap+1])
						new, err := strconv.Atoi(fmt.Sprintf("%d%d", subtotal, numbers[gap+1]))
						subtotal = new
						if err != nil {
							panic(err)
						}
					} else {
						s += fmt.Sprintf(" * %d", numbers[gap+1])
						subtotal *= numbers[gap+1]
					}
				} else {
					s += fmt.Sprintf(" + %d", numbers[gap+1])
					subtotal += numbers[gap+1]
				}
			}
			s += fmt.Sprintf(" => %d\n", subtotal)
			fmt.Print(s)
			if subtotal == target {
				return true
			}
		}
	}
	return false
}
