package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"mhayden/aoc2024/days/day01"
	"mhayden/aoc2024/days/day02"
	"mhayden/aoc2024/days/day03"
	"mhayden/aoc2024/days/day04"
	"mhayden/aoc2024/days/day05"
	"mhayden/aoc2024/days/day06"
	"mhayden/aoc2024/days/day07"
	"mhayden/aoc2024/days/day08"
	"mhayden/aoc2024/days/day09"
	"mhayden/aoc2024/days/day10"
	"mhayden/aoc2024/days/day11"
	"mhayden/aoc2024/days/day12"
	"mhayden/aoc2024/days/day13"
	"mhayden/aoc2024/days/day14"
	"mhayden/aoc2024/days/day15"
	"mhayden/aoc2024/days/day16"
	"mhayden/aoc2024/days/day17"
	"mhayden/aoc2024/days/day18"
	"mhayden/aoc2024/days/day19"
	"mhayden/aoc2024/days/day20"
	"mhayden/aoc2024/days/day21"
	"mhayden/aoc2024/days/day22"
	"mhayden/aoc2024/days/day23"
	"mhayden/aoc2024/days/day24"
	"mhayden/aoc2024/days/day25"
)

func main() {
	cpufile, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	defer pprof.StopCPUProfile()

	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc24 <day>")
		os.Exit(1)
	}

	day := os.Args[1]

	//fmt.Println("Allocs:", int(testing.AllocsPerRun(1, func() {
	switch day {
	case "01":
		day01.Solve()
	case "02":
		day02.Solve()
	case "03":
		day03.Solve()
	case "04":
		day04.Solve()
	case "05":
		day05.Solve()
	case "06":
		day06.Solve()
	case "07":
		day07.Solve()
	case "08":
		day08.Solve()
	case "09":
		day09.Solve()
	case "10":
		day10.Solve()
	case "11":
		day11.Solve()
	case "12":
		day12.Solve()
	case "13":
		day13.Solve()
	case "14":
		day14.Solve()
	case "15":
		day15.Solve()
	case "16":
		day16.Solve()
	case "17":
		day17.Solve()
	case "18":
		day18.Solve()
	case "19":
		day19.Solve()
	case "20":
		day20.Solve()
	case "21":
		day21.Solve()
	case "22":
		day22.Solve()
	case "23":
		day23.Solve()
	case "24":
		day24.Solve()
	case "25":
		day25.Solve()
	default:
		fmt.Println("Invalid day")
		os.Exit(1)
	}
	//})))
}
