package day20

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"math"
)

type Day20 struct {
	track         map[graph.Vector2]bool
	start         graph.Vector2
	end           graph.Vector2
	minTimeSaving int
}

func newDay20(data []string) *Day20 {
	return newDay20WithMinTimeSaving(data, 100)
}
func newDay20WithMinTimeSaving(data []string, minTimeSaving int) *Day20 {
	track := make(map[graph.Vector2]bool)
	var start graph.Vector2
	var end graph.Vector2
	for y, row := range data {
		for x, c := range row {
			if c == '#' {
				continue
			}
			pos := *graph.NewVector2(x, y)
			track[pos] = true
			if c == 'S' {
				start = pos
			} else if c == 'E' {
				end = pos
			}
		}
	}
	return &Day20{track: track, start: start, end: end, minTimeSaving: minTimeSaving}
}

func (d *Day20) minCosts(from graph.Vector2) map[graph.Vector2]int {
	costs := make(map[graph.Vector2]int)

	queue := []graph.Vector2{from}
	steps := 0
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			current := queue[0]
			queue = queue[1:]
			_, visited := costs[current]
			if visited {
				continue
			}
			costs[current] = steps
			for _, dir := range direction.Cardinals {
				next := *current.Add(dir.Delta())
				_, visitedNext := costs[next]
				if d.track[next] && !visitedNext {
					queue = append(queue, next)
				}
			}
		}
		steps++
	}
	return costs
}

type cheat struct {
	start  graph.Vector2
	end    graph.Vector2
	saving int
}

func (d *Day20) findCheats(distFromStart map[graph.Vector2]int, maxTime int) map[cheat]bool {
	cheats := make(map[cheat]bool)

	for start, startCost := range distFromStart {
		for x := -maxTime; x <= maxTime; x++ {
			maxY := maxTime - int(math.Abs(float64(x)))
			for y := -maxY; y <= maxY; y++ {
				if x == 0 && y == 0 {
					continue
				}
				end := *graph.NewVector2(start.X+x, start.Y+y)
				endCost, isPath := distFromStart[end]
				if !isPath {
					continue
				}
				saving := endCost - startCost - start.ManhattanDistance(end)
				if saving >= 0 {
					cheats[cheat{
						start:  start,
						end:    end,
						saving: saving,
					}] = true
				}
			}
		}
	}
	return cheats
}

func (d *Day20) part1() int {
	distFromStart := d.minCosts(d.start)
	cheats := d.findCheats(distFromStart, 2)

	total := 0
	for c := range cheats {
		if c.saving >= d.minTimeSaving {
			total++
		}
	}
	return total
}

func (d *Day20) part2() int {
	distFromStart := d.minCosts(d.start)
	cheats := d.findCheats(distFromStart, 20)

	total := 0
	for c := range cheats {
		if c.saving >= d.minTimeSaving {
			total++
		}
	}
	return total
}
