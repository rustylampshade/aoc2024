package xygrid

import (
	"fmt"
	"strings"
)

// A grid of single-char cells that have the origin 0,0 in the upper left, counting
// the rows as Y goes from 0..num_rows and the columns as X goes from 0..num_cols.
type XYGrid struct {
	Num_rows int
	Num_cols int
	Cells    [][]Cell
}

type Cell struct {
	Val     string
	Flagged bool
}

type Offset struct {
	X int
	Y int
}

type Coord [2]int
type Direction int
type RelativeDirection int

const (
	North     Direction = iota
	Northeast           = iota
	East                = iota
	Southeast           = iota
	South               = iota
	Southwest           = iota
	West                = iota
	Northwest           = iota
)

const (
	Left  RelativeDirection = iota
	Right                   = iota
)

var AllDirections [8]Direction = [8]Direction{North, Northeast, East, Southeast, South, Southwest, West, Northwest}

func NewXYGrid(data []byte) *XYGrid {
	xy := XYGrid{}
	s := string(data)
	lines := strings.Split(s, "\n")
	xy.Num_rows = len(lines)
	xy.Num_cols = len(lines[0])
	xy.Cells = make([][]Cell, len(lines))
	for j, line := range lines {
		chars := strings.Split(line, "")
		xy.Cells[j] = make([]Cell, len(chars))
		for i, char := range chars {
			xy.Cells[j][i] = Cell{Val: char}
		}
	}
	return &xy
}

func (xy XYGrid) String() string {
	lines := []string{}
	for j, row := range xy.Cells {
		s := fmt.Sprintf("%02d: ", j)
		for _, c := range row {
			s += c.Val
		}
		lines = append(lines, s)
	}
	return strings.Join(lines, "\n")
}

// Starting at the given coord, what is the string formed by walking strlen positions
// in the given direction? Can return a shorter string than strlen if the edge of the grid
// is hit.
func (xy XYGrid) StrInDirection(loc Coord, strlen int, direction Direction) (s string) {
	return xy.strInDirection(loc, direction.Offset(), strlen)
}

func (xy XYGrid) strInDirection(loc Coord, offset Offset, strlen int) (stringAlongPath string) {
	stringAlongPath = xy.Get(loc)
	for count := 1; count < strlen; count++ {
		offsetCoord := loc.NextAlongPath(offset)
		if !xy.ContainsCoord(offsetCoord) {
			break
		}
		stringAlongPath += xy.Get(offsetCoord)
		loc = offsetCoord
	}
	return stringAlongPath
}

// Untested
func (xy XYGrid) Neighbors(x, y int) []Coord {
	coords := []Coord{}
	for _, j := range []int{y - 1, y + 1} {
		for _, i := range []int{x - 1, x + 1} {
			if !xy.ContainsCoord(Coord{i, j}) {
				continue
			}
			coords = append(coords, Coord{i, j})
		}
	}
	return coords
}

func (d Direction) Offset() Offset {
	switch d {
	case North:
		return Offset{0, -1}
	case Northeast:
		return Offset{1, -1}
	case East:
		return Offset{1, 0}
	case Southeast:
		return Offset{1, 1}
	case South:
		return Offset{0, 1}
	case Southwest:
		return Offset{-1, 1}
	case West:
		return Offset{-1, 0}
	case Northwest:
		return Offset{-1, -1}
	}
	panic("Impossible direction")
}

func (xy XYGrid) ContainsCoord(loc Coord) bool {
	i := loc[0]
	j := loc[1]
	return !(i < 0 || i >= xy.Num_cols || j < 0 || j >= xy.Num_rows)
}

func (d Direction) TurnRelative(rd RelativeDirection) Direction {
	switch d {
	case North:
		switch rd {
		case Right:
			return East
		case Left:
			return West
		}
	case East:
		switch rd {
		case Right:
			return South
		case Left:
			return North
		}
	case South:
		switch rd {
		case Right:
			return West
		case Left:
			return East
		}
	case West:
		switch rd {
		case Right:
			return North
		case Left:
			return South
		}
	}
	panic("Impossible direction")
}

func (c Coord) NextAlongPath(o Offset) Coord {
	i := c[0] + o.X
	j := c[1] + o.Y
	return Coord{i, j}
}

func (xy XYGrid) Get(c Coord) string {
	return xy.Cells[c[1]][c[0]].Val
}

func (d Direction) DirectionSymbol() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	}
	panic("Invalid direction to get symbol")
}
