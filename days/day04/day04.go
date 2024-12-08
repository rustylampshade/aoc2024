package day04

import (
	"fmt"
	"os"

	ds "mhayden/aoc2024/pkg/datastructures"
)

func Solve() {
	dat, err := os.ReadFile("inputs/day04.txt")
	if err != nil {
		panic(err)
	}
	xy := ds.NewXYGrid(dat)

	word_count := 0
	for j, row := range xy.Cells {
		for i, cell := range row {
			if cell.Val == "X" {
				for _, d := range ds.AllDirections {
					if xy.StrInDirection(ds.Coord{i, j}, 4, d) == "XMAS" {
						word_count++
					}
				}
			}
		}
	}
	fmt.Printf("%d instances of XMAS\n", word_count)

	dumb_xcount := 0
	for j, row := range xy.Cells {
		if j == 0 || j == xy.Num_rows-1 {
			continue
		}
		for i, cell := range row {
			if i == 0 || i == xy.Num_cols-1 {
				continue
			}
			if cell.Val == "A" {
				southeast := xy.StrInDirection(ds.Coord{i - 1, j - 1}, 3, ds.Southeast)
				southwest := xy.StrInDirection(ds.Coord{i + 1, j - 1}, 3, ds.Southwest)
				if (southeast == "MAS" || southeast == "SAM") && (southwest == "MAS" || southwest == "SAM") {
					dumb_xcount++
				}
			}
		}
	}
	fmt.Printf("%d instances of X-MAS\n", dumb_xcount)
}
