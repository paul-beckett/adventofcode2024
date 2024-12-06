package direction

import (
	"adventofcode2024/util/graph"
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var deltas = map[Direction]graph.Vector2{
	Up:    *graph.NewVector2(0, -1),
	Right: *graph.NewVector2(1, 0),
	Down:  *graph.NewVector2(0, 1),
	Left:  *graph.NewVector2(-1, 0),
}

func (d Direction) Delta() graph.Vector2 {
	return deltas[d]
}

func (d Direction) Clockwise() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic(fmt.Errorf("unknown direction %d", d))
	}
}
