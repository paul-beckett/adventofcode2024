package day12

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
)

type Day12 struct {
	garden map[graph.Vector2]rune
}

func newDay12(data []string) *Day12 {
	garden := make(map[graph.Vector2]rune)
	for y, row := range data {
		for x, c := range row {
			garden[*graph.NewVector2(x, y)] = c
		}
	}
	return &Day12{garden: garden}

}

type region map[graph.Vector2]bool

func (r *region) area() int {
	return len(*r)
}

func (r *region) perimeter() int {
	perimeter := 0
	for gp := range *r {
		for _, dir := range direction.Cardinals {
			side := *gp.Add(dir.Delta())
			if !(*r)[side] {
				perimeter++
			}
		}
	}
	return perimeter
}

func (r *region) sides() int {
	corners := 0
	for gp := range *r {
		checks := [][]direction.Direction{
			{direction.Up, direction.Right},
			{direction.Right, direction.Down},
			{direction.Down, direction.Left},
			{direction.Left, direction.Up},
		}
		for _, check := range checks {
			lInGrid := (*r)[*gp.Add(check[0].Delta())]
			rInGrid := (*r)[*gp.Add(check[1].Delta())]
			diagInGrid := (*r)[*gp.Add(*check[0].Delta().Add(check[1].Delta()))]
			if (!lInGrid && !rInGrid) || (lInGrid && rInGrid && !diagInGrid) {
				corners++
			}
		}
	}
	return corners
}

func (d *Day12) findRegions() []region {
	var regions []region
	visited := make(map[graph.Vector2]bool)
	for plot := range d.garden {
		if visited[plot] {
			continue
		}
		r := make(region)
		queue := []graph.Vector2{plot}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			if visited[current] {
				continue
			}
			visited[current] = true
			r[current] = true
			for _, dir := range direction.Cardinals {
				next := *current.Add(dir.Delta())
				if !visited[next] && d.garden[current] == d.garden[next] {
					queue = append(queue, next)
				}
			}
		}
		regions = append(regions, r)
	}
	return regions
}

func (d *Day12) cost(f func(r region) int) int {
	totalCost := 0
	for _, r := range d.findRegions() {
		totalCost += f(r)
	}
	return totalCost
}

func (d *Day12) part1() int {
	return d.cost(func(r region) int {
		return r.area() * r.perimeter()
	})
}

func (d *Day12) part2() int {
	return d.cost(func(r region) int {
		return r.area() * r.sides()
	})
}
