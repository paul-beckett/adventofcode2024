package day11

import (
	"math"
	"strings"
)

type Day11 struct {
	data []string
}

func newDay11(data []string) *Day11 {
	return &Day11{data: data}
}

// hexagon is based on the cube coords from https://www.redblobgames.com/grids/hexagons/#coordinates-cube
type hexagon struct {
	q int
	r int
	s int
}

func (h *hexagon) add(o hexagon) hexagon {
	return hexagon{q: h.q + o.q, r: h.r + o.r, s: h.s + o.s}
}

func (h *hexagon) subtract(o hexagon) hexagon {
	return hexagon{q: h.q - o.q, r: h.r - o.r, s: h.s - o.s}
}

func (h *hexagon) dist(o hexagon) int {
	abs := func(n int) int { return int(math.Abs(float64(n))) }
	v := h.subtract(o)
	return (abs(v.q) + abs(v.r) + abs(v.s)) / 2
}

var moves = map[string]hexagon{
	"ne": {q: 1, r: -1},
	"se": {q: 1, s: -1},
	"s":  {r: 1, s: -1},
	"sw": {q: -1, r: 1},
	"nw": {q: -1, s: 1},
	"n":  {r: -1, s: 1},
}

func (d *Day11) part1() int {
	start := hexagon{q: 0, r: 0, s: 0}
	dirs := strings.Split(d.data[0], ",")

	end, _ := followPath(start, dirs)
	return end.dist(start)
}

func followPath(start hexagon, dirs []string) (hexagon, hexagon) {
	current := start
	furthest := start
	for _, dir := range dirs {
		move, exists := moves[dir]
		if !exists {
			panic("invalid dir " + dir)
		}
		current = current.add(move)
		if start.dist(current) > start.dist(furthest) {
			furthest = current
		}
	}
	return current, furthest
}

func (d *Day11) part2() int {
	start := hexagon{q: 0, r: 0, s: 0}
	dirs := strings.Split(d.data[0], ",")

	_, furthest := followPath(start, dirs)
	return furthest.dist(start)
}
