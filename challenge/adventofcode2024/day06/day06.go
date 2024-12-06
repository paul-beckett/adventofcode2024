package day06

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
)

type Day06 struct {
	data []string
}

func newDay06(data []string) *Day06 {
	return &Day06{data: data}
}

type visitedPath map[graph.Vector2]map[direction.Direction]bool

func (v *visitedPath) visit(pos graph.Vector2, d direction.Direction) {
	dirsAtPosition, ok := (*v)[pos]
	if !ok {
		dirsAtPosition = make(map[direction.Direction]bool)
		(*v)[pos] = dirsAtPosition
	}
	dirsAtPosition[d] = true
}

func (v *visitedPath) alreadyVisited(pos graph.Vector2, d direction.Direction) bool {
	dirsAtPosition, ok := (*v)[pos]
	if !ok {
		return false
	}
	return dirsAtPosition[d]
}

func (d *Day06) findStart() *graph.Vector2 {
	for y, row := range d.data {
		for x, v := range row {
			if v == '^' {
				return graph.NewVector2(x, y)
			}
		}
	}
	panic("start not found")
}

func (d *Day06) isInMappedArea(v graph.Vector2) bool {
	return v.Y >= 0 && v.Y < len(d.data) && v.X >= 0 && v.X < len(d.data[v.Y])
}

func (d *Day06) isWall(v graph.Vector2) bool {
	return d.data[v.Y][v.X] == '#'
}

type pathSegment struct {
	position  graph.Vector2
	direction direction.Direction
}

func newPathSegment(position graph.Vector2, direction direction.Direction) *pathSegment {
	return &pathSegment{position: position, direction: direction}
}

// walks the path until it exits the graph or a loop is found returns the number
//
// returns the path and true if the graph is exited, false if a loop was found
func (d *Day06) walkPath(position graph.Vector2, direction direction.Direction, v visitedPath) ([]pathSegment, bool) {
	var path []pathSegment
	for {
		if v.alreadyVisited(position, direction) {
			return path, false
		}
		v.visit(position, direction)
		path = append(path, *newPathSegment(position, direction))
		nextPosition := *position.Add(direction.Delta())

		if !d.isInMappedArea(nextPosition) {
			return path, true
		} else if d.isWall(nextPosition) {
			direction = direction.Clockwise()
		} else {
			position = nextPosition
		}
	}
}

func (d *Day06) part1() int {
	path, _ := d.walkPath(*d.findStart(), direction.Up, make(visitedPath))
	visited := make(map[graph.Vector2]bool)
	for _, p := range path {
		visited[p.position] = true
	}
	return len(visited)
}

func (d *Day06) part2() int {
	start := d.findStart()
	path, _ := d.walkPath(*start, direction.Up, make(visitedPath))
	loopPositions := make(map[graph.Vector2]bool)
	for _, p := range path {
		if p.position != *start {
			row := d.data[p.position.Y]
			d.data[p.position.Y] = row[:p.position.X] + "#" + row[p.position.X+1:]
			_, exit := d.walkPath(*start, direction.Up, make(visitedPath))
			if !exit {
				loopPositions[p.position] = true
			}
			d.data[p.position.Y] = row
		}
	}
	return len(loopPositions)
}
