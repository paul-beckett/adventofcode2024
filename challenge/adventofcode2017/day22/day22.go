package day22

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"fmt"
)

type Day22 struct {
	data []string
}

func newDay22(data []string) *Day22 {
	return &Day22{data: data}
}

func (d *Day22) part1() int {
	cluster := make(map[graph.Vector2]bool)
	for y, line := range d.data {
		for x, c := range line {
			if c == '#' {
				cluster[*graph.NewVector2(x, y)] = true
			}
		}
	}

	infectionCount := 0
	current := graph.NewVector2(len(d.data[0])/2, len(d.data)/2)
	dir := direction.Up
	for range 10_000 {
		if cluster[*current] {
			dir = dir.Clockwise()
			delete(cluster, *current)
		} else {
			dir = dir.AntiClockwise()
			cluster[*current] = true
			infectionCount++
		}
		current = current.Add(dir.Delta())
	}
	return infectionCount
}

type nodeState int

const (
	clean nodeState = iota
	weakened
	infected
	flagged
)

func (d nodeState) next() nodeState {
	switch d {
	case clean:
		return weakened
	case weakened:
		return infected
	case infected:
		return flagged
	case flagged:
		return clean
	default:
		panic(fmt.Errorf("unknown state: %d", d))
	}
}

func (d *Day22) part2() int {
	cluster := make(map[graph.Vector2]nodeState)
	for y, line := range d.data {
		for x, c := range line {
			if c == '#' {
				cluster[*graph.NewVector2(x, y)] = infected
			}
		}
	}

	infectionCount := 0
	current := graph.NewVector2(len(d.data[0])/2, len(d.data)/2)
	dir := direction.Up
	for range 10_000_000 {
		state, exists := cluster[*current]
		if !exists {
			state = clean
		}
		switch state {
		case clean:
			dir = dir.AntiClockwise()
		case weakened:
		case infected:
			dir = dir.Clockwise()
		case flagged:
			dir = dir.Clockwise().Clockwise()
		}
		nextState := state.next()
		if nextState == clean {
			delete(cluster, *current)
		} else {
			cluster[*current] = nextState
		}
		if nextState == infected {
			infectionCount++
		}
		current = current.Add(dir.Delta())
	}
	return infectionCount
}
