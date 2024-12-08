package day08

import (
	"fmt"
	"os"
	"strings"

	ds "mhayden/aoc2024/pkg/datastructures"
)

func Solve() {
	dat, err := os.ReadFile("inputs/day08.txt")
	if err != nil {
		panic(err)
	}
	xy, antennas := newXY(dat)

	partA := make(map[ds.Coord]bool)
	equations := make(map[equation]bool)
	for _, locations := range antennas {
		for left := 0; left < len(locations)-1; left++ {
			for right := left + 1; right < len(locations); right++ {
				a1, a2, e := antinodes(locations[left], locations[right])
				if xy.ContainsCoord(a1) {
					partA[a1] = true
				}
				if xy.ContainsCoord(a2) {
					partA[a2] = true
				}
				equations[e] = true
			}
		}
	}

	partB := make(map[ds.Coord]bool)
	for j, row := range xy.Cells {
		for i := range row {
			for equation := range equations {
				if equation.fallsOnLine(ds.Coord{i, j}) {
					xy.Cells[j][i].Val = "X"
					partB[ds.Coord{i, j}] = true
					break
				}
			}
		}
	}
	fmt.Println(xy)
	fmt.Println(len(partA))
	fmt.Println(len(partB))
}

func newXY(data []byte) (*ds.XYGrid, map[string][]ds.Coord) {
	xy := ds.XYGrid{}
	s := string(data)
	lines := strings.Split(s, "\n")
	xy.Num_rows = len(lines)
	xy.Num_cols = len(lines[0])
	xy.Cells = make([][]ds.Cell, len(lines))
	antennas := make(map[string][]ds.Coord)
	for j, line := range lines {
		chars := strings.Split(line, "")
		xy.Cells[j] = make([]ds.Cell, len(chars))
		for i, char := range chars {
			xy.Cells[j][i] = ds.Cell{Val: char}
			if char != "." {
				if _, exists := antennas[char]; !exists {
					antennas[char] = []ds.Coord{}
				}
				antennas[char] = append(antennas[char], ds.Coord{i, j})
			}
		}
	}
	return &xy, antennas
}

func antinodes(p1, p2 ds.Coord) (ds.Coord, ds.Coord, equation) {
	ax, ay := p1[0], p1[1]
	bx, by := p2[0], p2[1]
	dx, dy := ax-bx, ay-by
	return ds.Coord{ax + dx, ay + dy}, ds.Coord{bx - dx, by - dy}, equation{x1: ax, x2: bx, y1: ay, y2: by}
}

type equation struct {
	x1, y1, x2, y2 int
}

func (e equation) fallsOnLine(loc ds.Coord) bool {
	x, y := loc[0], loc[1]
	return float64(y-e.y1) == float64((x-e.x1)*(e.y2-e.y1))/float64(e.x2-e.x1)
}
