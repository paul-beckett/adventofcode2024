package day06

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
)

type Day06 struct {
	maxX, minX, maxY, minY int
	obstacles              map[graph.Vector2]bool
	start                  pathSegment
}

func newDay06(data []string) *Day06 {
	obstacles := make(map[graph.Vector2]bool)
	var startPosition graph.Vector2
	for y, line := range data {
		for x, c := range line {
			switch c {
			case '#':
				obstacles[*graph.NewVector2(x, y)] = true
			case '^':
				startPosition = *graph.NewVector2(x, y)
			}
		}
	}
	return &Day06{
		minY:      0,
		maxY:      len(data) - 1,
		minX:      0,
		maxX:      len(data[0]) - 1,
		obstacles: obstacles,
		start:     *newPathSegment(startPosition, direction.Up),
	}
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
func (d *Day06) walkPath(start pathSegment) ([]pathSegment, bool) {
	path := []pathSegment{start}

	visited := make(map[pathSegment]bool)
	visited[start] = true
	//for _, p := range pathTo {
	//	visited[p] = true
	//}

	current := start
	for {
		nextDir := current.direction
		nextPosition := *current.position.Add(nextDir.Delta())
		for d.isWall(nextPosition) {
			nextDir = nextDir.Clockwise()
			nextPosition = *current.position.Add(nextDir.Delta())
		}
		next := *newPathSegment(nextPosition, nextDir)

		if !d.isInMappedArea(next.position) {
			return path, true
		} else if visited[next] {
			return path, false
		} else {
			visited[next] = true
			path = append(path, next)
			current = next
		}
	}
}

func (d *Day06) part1() int {
	path, _ := d.walkPath(d.start)
	visited := make(map[graph.Vector2]bool)
	for _, p := range path {
		visited[p.position] = true
	}
	return len(visited)
}

func (d *Day06) part2() int {
	path, _ := d.walkPath(d.start)

	startPositions := make(map[graph.Vector2]bool)
	addedObstacles := make(map[graph.Vector2]bool)
	for i := 0; i < len(path)-1; i++ {
		startFrom := path[i]
		startPositions[startFrom.position] = true
		obstacle := path[i+1]
		if startPositions[obstacle.position] {
			continue
		}

		d.obstacles[obstacle.position] = true
		_, exit := d.walkPath(startFrom)
		if !exit {
			addedObstacles[obstacle.position] = true
		}
		delete(d.obstacles, obstacle.position)

	}
	return len(addedObstacles)
}
