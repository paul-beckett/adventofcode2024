package day21

import (
	"adventofcode2024/util/graph"
)

type Day21 struct {
	data []string
}

func newDay21(data []string) *Day21 {
	return &Day21{data: data}
}

var (
	numericLayout = layout{
		'7': *graph.NewVector2(0, 0), '8': *graph.NewVector2(1, 0), '9': *graph.NewVector2(2, 0),
		'4': *graph.NewVector2(0, 1), '5': *graph.NewVector2(1, 1), '6': *graph.NewVector2(2, 1),
		'1': *graph.NewVector2(0, 2), '2': *graph.NewVector2(1, 2), '3': *graph.NewVector2(2, 2),
		'0': *graph.NewVector2(1, 3), 'A': *graph.NewVector2(2, 3),
	}
	directionLayout = layout{
		'^': *graph.NewVector2(1, 0), 'A': *graph.NewVector2(2, 0),
		'<': *graph.NewVector2(0, 1), 'v': *graph.NewVector2(1, 1), '>': *graph.NewVector2(2, 1),
	}
)

type keypad struct {
	position rune
	layout   layout
	next     *keypad
}

type layout map[rune]graph.Vector2

func (d *Day21) part1() int {
	return -1
}

func (d *Day21) part2() int {
	return -1
}
