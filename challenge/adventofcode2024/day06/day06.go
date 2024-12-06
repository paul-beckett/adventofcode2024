package day06

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"maps"
)

type Day06 struct {
	maxX, minX, maxY, minY int
	obstacles              map[graph.Vector2]bool
	start                  graph.Vector2
}

func newDay06(data []string) *Day06 {
	obstacles := make(map[graph.Vector2]bool)
	var start graph.Vector2
	for y, line := range data {
		for x, c := range line {
			switch c {
			case '#':
				obstacles[*graph.NewVector2(x, y)] = true
			case '^':
				start = *graph.NewVector2(x, y)
			}
		}
	}
	return &Day06{
		minY:      0,
		maxY:      len(data) - 1,
		minX:      0,
		maxX:      len(data[0]) - 1,
		obstacles: obstacles,
		start:     start,
	}
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

func (d *Day06) isInMappedArea(v graph.Vector2) bool {
	return v.Y >= d.minY && v.Y <= d.maxY && v.X >= d.minX && v.X <= d.maxX
}

func (d *Day06) isWall(v graph.Vector2) bool {
	return d.obstacles[v]
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
func (d *Day06) walkPath(startPos graph.Vector2, startDir direction.Direction, v visitedPath) ([]pathSegment, bool) {
	path := []pathSegment{*newPathSegment(startPos, startDir)}
	currentPos := startPos
	currentDir := startDir
	for {
		v.visit(currentPos, currentDir)
		nextPosition := *currentPos.Add(currentDir.Delta())

		if !d.isInMappedArea(nextPosition) {
			return path, true
		} else if v.alreadyVisited(nextPosition, currentDir) {
			return path, false
		} else if d.isWall(nextPosition) {
			currentDir = currentDir.Clockwise()
		} else {
			path = append(path, *newPathSegment(nextPosition, currentDir))
			currentPos = nextPosition
		}
	}
}

func (d *Day06) part1() int {
	path, _ := d.walkPath(d.start, direction.Up, make(visitedPath))
	visited := make(map[graph.Vector2]bool)
	for _, p := range path {
		visited[p.position] = true
	}
	return len(visited)
}

func (d *Day06) part2() int {
	visited := make(visitedPath)
	path, _ := d.walkPath(d.start, direction.Up, maps.Clone(visited))

	loopPositions := make(map[graph.Vector2]bool)
	for i := 0; i < len(path)-1; i++ {
		obstacle := path[i+1].position
		if obstacle == d.start || visited[obstacle] != nil {
			continue
		}

		d.obstacles[obstacle] = true
		_, exit := d.walkPath(d.start, direction.Up, maps.Clone(visited))
		if !exit {
			loopPositions[obstacle] = true
		}
		delete(d.obstacles, obstacle)

	}
	return len(loopPositions)
}
