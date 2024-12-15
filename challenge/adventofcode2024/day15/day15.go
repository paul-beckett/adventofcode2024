package day15

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"maps"
	"strings"
)

type Day15 struct {
	boxes      map[graph.Vector2]box
	walls      map[graph.Vector2]bool
	directions []direction.Direction
	robotStart graph.Vector2
}

func newDay15(data []string) *Day15 {
	boxes := make(map[graph.Vector2]box)
	walls := make(map[graph.Vector2]bool)

	var robotStart graph.Vector2

	i := 0
	for i < len(data) {
		if strings.TrimSpace(data[i]) == "" {
			break
		}
		for x, c := range data[i] {
			switch c {
			case '#':
				walls[*graph.NewVector2(x, i)] = true
			case 'O':
				b := *graph.NewVector2(x, i)
				boxes[b] = box{b}
			case '@':
				robotStart = *graph.NewVector2(x, i)
			default:
				continue
			}
		}
		i++
	}
	i++

	dirByC := map[rune]direction.Direction{
		'^': direction.Up,
		'>': direction.Right,
		'v': direction.Down,
		'<': direction.Left,
	}
	var directions []direction.Direction
	for i < len(data) {
		for _, c := range data[i] {
			directions = append(directions, dirByC[c])
		}
		i++
	}

	return &Day15{
		boxes:      boxes,
		walls:      walls,
		directions: directions,
		robotStart: robotStart,
	}
}

type box []graph.Vector2

func canMove(pos graph.Vector2, dir direction.Direction, boxes map[graph.Vector2]box, walls map[graph.Vector2]bool) bool {
	next := *pos.Add(dir.Delta())
	if walls[next] {
		return false
	}
	for _, b := range boxes[next] {
		if b == pos {
			continue
		}
		if !canMove(b, dir, boxes, walls) {
			return false
		}
	}
	return true
}

func moveBox(b box, dir direction.Direction, boxes map[graph.Vector2]box) {
	var pushes []graph.Vector2
	switch dir {
	case direction.Up, direction.Down:
		for _, pos := range b {
			pushes = append(pushes, *pos.Add(dir.Delta()))
		}
	case direction.Left:
		pushes = append(pushes, *b[0].Add(dir.Delta()))
	case direction.Right:
		pushes = append(pushes, *b[len(b)-1].Add(dir.Delta()))
	}
	for _, p := range pushes {
		b := boxes[p]
		if len(b) > 0 {
			moveBox(b, dir, boxes)
		}
	}
	var newBox box
	for _, pos := range b {
		newBox = append(newBox, *pos.Add(dir.Delta()))
		delete(boxes, pos)
	}
	for _, pos := range newBox {
		boxes[pos] = newBox
	}
}

func moveRobot(robot graph.Vector2, directions []direction.Direction, boxes map[graph.Vector2]box, walls map[graph.Vector2]bool) {
	current := robot
	for _, dir := range directions {
		if canMove(current, dir, boxes, walls) {
			next := *current.Add(dir.Delta())
			b := boxes[next]
			if b != nil {
				moveBox(b, dir, boxes)
			}
			current = next
		}
	}
}

func sumGps(boxes map[graph.Vector2]box) int {
	total := 0
	visited := make(map[graph.Vector2]bool)
	for p, bo := range boxes {
		if !visited[p] {
			for _, b := range bo {
				visited[b] = true
			}
			total += bo[0].Y * 100
			total += bo[0].X
		}
	}
	return total
}

func (d *Day15) part1() int {
	boxes := maps.Clone(d.boxes)
	moveRobot(d.robotStart, d.directions, boxes, d.walls)
	return sumGps(boxes)
}

func (d *Day15) part2() int {
	wideBoxes := make(map[graph.Vector2]box)
	for _, b := range d.boxes {
		var newBox box
		for _, pos := range b {
			pos.X = pos.X * 2
			newBox = append(newBox, pos)
			newBox = append(newBox, *pos.Add(direction.Right.Delta()))
		}
		for _, pos := range newBox {
			wideBoxes[pos] = newBox
		}
	}
	wideWalls := make(map[graph.Vector2]bool)
	for pos := range d.walls {
		pos.X = pos.X * 2
		wideWalls[pos] = true
		wideWalls[*pos.Add(direction.Right.Delta())] = true
	}
	newStart := *graph.NewVector2(d.robotStart.X*2, d.robotStart.Y)
	moveRobot(newStart, d.directions, wideBoxes, wideWalls)
	return sumGps(wideBoxes)
}
