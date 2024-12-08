package day06

import (
	"fmt"
	"os"

	ds "mhayden/aoc2024/pkg/datastructures"
)

func Solve() {
	xy := initGrid()

	startingGuardCoord := ds.Coord{}
	for j, row := range xy.Cells {
		for i := range row {
			if xy.Cells[j][i].Val == "^" {
				startingGuardCoord = ds.Coord{i, j}
			}
		}
	}
	simulateWalk(xy, startingGuardCoord)

	visited_coords := make(map[ds.Coord]bool)
	for j, row := range xy.Cells {
		for i := range row {
			if xy.Cells[j][i].Flagged {
				visited_coords[ds.Coord{i, j}] = true
			}
		}
	}
	fmt.Println(len(visited_coords))

	cyclic_obstructions := 0
	for possible_obstruction := range visited_coords {
		xy = initGrid()
		xy.Cells[possible_obstruction[1]][possible_obstruction[0]].Val = "#"
		if simulateWalk(xy, startingGuardCoord) {
			cyclic_obstructions++
		}
	}
	fmt.Println(cyclic_obstructions)
}

func initGrid() *ds.XYGrid {
	dat, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		panic(err)
	}
	return ds.NewXYGrid(dat)
}

func simulateWalk(xy *ds.XYGrid, startingGuardCoord ds.Coord) (cycleDetected bool) {
	xy.Cells[startingGuardCoord[1]][startingGuardCoord[0]].Flagged = true
	direction := ds.North
	sigil := direction.DirectionSymbol()
	movement := direction.Offset()
	currentCoord := startingGuardCoord
	nextCoord := currentCoord.NextAlongPath(movement)
	for xy.ContainsCoord(nextCoord) {
		if xy.Get(nextCoord) == "#" {
			direction = direction.TurnRelative(ds.Right)
			movement = direction.Offset()
			sigil = direction.DirectionSymbol()
		} else {
			currentCoord = nextCoord
			if xy.Cells[currentCoord[1]][currentCoord[0]].Flagged && xy.Get(currentCoord) == sigil {
				return true
			}
			xy.Cells[currentCoord[1]][currentCoord[0]].Flagged = true
			xy.Cells[currentCoord[1]][currentCoord[0]].Val = sigil
		}
		nextCoord = currentCoord.NextAlongPath(movement)
	}
	return false
}
