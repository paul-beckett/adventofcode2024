package day10

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"strconv"
)

type Day10 struct {
	topo        map[graph.Vector2]int
	topLeft     graph.Vector2
	bottomRight graph.Vector2
}

var directions = []direction.Direction{direction.Up, direction.Down, direction.Left, direction.Right}

func newDay10(data []string) *Day10 {
	topo := make(map[graph.Vector2]int)
	topLeft := *graph.NewVector2(0, 0)
	maxY := len(data) - 1
	maxX := len(data[maxY]) - 1
	bottomRight := *graph.NewVector2(maxX, maxY)
	for y, line := range data {
		for x, v := range line {
			height, _ := strconv.Atoi(string(v))
			topo[*graph.NewVector2(x, y)] = height
		}
	}
	return &Day10{
		topo:        topo,
		topLeft:     topLeft,
		bottomRight: bottomRight,
	}
}

func (d *Day10) canTraverse(from, to graph.Vector2) bool {
	return to.X >= d.topLeft.X && to.Y >= d.topLeft.Y &&
		to.X <= d.bottomRight.X && to.Y <= d.bottomRight.Y &&
		d.topo[to]-d.topo[from] == 1
}

func (d *Day10) reachableEnds(current graph.Vector2, visited map[graph.Vector2]bool, score *int, allowRepeats bool) {
	visited[current] = true
	if d.topo[current] == 9 {
		*score++
		return
	}
	for _, dir := range directions {
		next := *current.Add(dir.Delta())
		if (allowRepeats || !visited[next]) && d.canTraverse(current, next) {
			d.reachableEnds(next, visited, score, allowRepeats)
		}
	}
}

func (d *Day10) getScore(allowRepeats bool) int {
	totalScore := 0
	for pos, height := range d.topo {
		if height != 0 {
			continue
		}
		visited := make(map[graph.Vector2]bool)
		d.reachableEnds(pos, visited, &totalScore, allowRepeats)
	}
	return totalScore
}

func (d *Day10) part1() int {
	return d.getScore(false)
}

func (d *Day10) part2() int {
	return d.getScore(true)
}

// An alternative solution for part2 using a bit of dp where we keep the score for each height to get to 9.
// Any height 9 fields score is 1
// Any height n field is the sum of all reachable height n+1 fields
// Then simply sum the scores from the height 0 fields
func (d *Day10) part2Alternate() int {
	byHeight := make([]map[graph.Vector2]bool, 10)
	for i := range byHeight {
		byHeight[i] = make(map[graph.Vector2]bool)
	}
	for pos, height := range d.topo {
		byHeight[height][pos] = true
	}

	scores := make(map[graph.Vector2]int)
	for pos := range byHeight[len(byHeight)-1] {
		scores[pos] = 1
	}

	for i := len(byHeight) - 2; i >= 0; i-- {
		for pos := range byHeight[i] {
			for _, dir := range directions {
				next := *pos.Add(dir.Delta())
				if byHeight[i+1][next] {
					scores[pos] += scores[next]
				}
			}
		}
	}
	total := 0
	for trailhead := range byHeight[0] {
		total += scores[trailhead]
	}
	return total
}
