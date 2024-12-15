package day15

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"fmt"
	"maps"
	"strings"
)

type Day15 struct {
	warehouse  warehouse
	robotStart graph.Vector2
	directions []direction.Direction
}

func newDay15(data []string) *Day15 {
	wh := make(warehouse)

	var robotStart graph.Vector2

	i := 0
	for i < len(data) {
		if strings.TrimSpace(data[i]) == "" {
			break
		}
		for x, c := range data[i] {
			pos := *graph.NewVector2(x, i)
			switch c {
			case '#':
				wh[pos] = newWall(pos)
			case 'O':
				wh[pos] = newBox(pos)
			case '@':
				robotStart = pos
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
		warehouse:  wh,
		robotStart: robotStart,
		directions: directions,
	}
}

type warehouse map[graph.Vector2]obstacle

type obstacle interface {
	canMove(dir direction.Direction, w warehouse) bool
	move(dir direction.Direction, w warehouse)
	gpsScore() int
}

type wall struct {
	position graph.Vector2
}

func newWall(position graph.Vector2) *wall {
	return &wall{
		position: position,
	}
}

func (w *wall) canMove(_ direction.Direction, _ warehouse) bool {
	return false
}
func (w *wall) move(_ direction.Direction, _ warehouse) {
	panic(fmt.Sprintf("cannot move wall %v", w))
}

func (w *wall) gpsScore() int { return 0 }

type box struct {
	position graph.Vector2
}

func newBox(position graph.Vector2) *box {
	return &box{
		position: position,
	}
}

func (b *box) canMove(dir direction.Direction, wh warehouse) bool {
	next := *b.position.Add(dir.Delta())
	obstacle, hasObstacle := wh[next]
	if !hasObstacle {
		return true
	} else {
		return obstacle.canMove(dir, wh)
	}
}
func (b *box) move(dir direction.Direction, wh warehouse) {
	next := *b.position.Add(dir.Delta())
	obstacle, hasObstacle := wh[next]
	if hasObstacle {
		obstacle.move(dir, wh)
	}
	delete(wh, b.position)
	b.position = next
	wh[next] = b
}

func (b *box) gpsScore() int {
	return b.position.Y*100 + b.position.X
}

type bigBox struct {
	l, r box
}

func newBigBox(position graph.Vector2) *bigBox {
	return &bigBox{
		l: *newBox(position),
		r: *newBox(*position.Add(direction.Right.Delta())),
	}
}

func (b *bigBox) canMove(dir direction.Direction, wh warehouse) bool {
	switch dir {
	case direction.Left:
		return b.l.canMove(dir, wh)
	case direction.Right:
		return b.r.canMove(dir, wh)
	default:
		return b.l.canMove(dir, wh) && b.r.canMove(dir, wh)
	}
}

func (b *bigBox) move(dir direction.Direction, wh warehouse) {
	var moveablePositions []graph.Vector2
	switch dir {
	case direction.Left:
		moveablePositions = append(moveablePositions, *b.l.position.Add(dir.Delta()))
	case direction.Right:
		moveablePositions = append(moveablePositions, *b.r.position.Add(dir.Delta()))
	default:
		moveablePositions = append(moveablePositions, *b.l.position.Add(dir.Delta()))
		moveablePositions = append(moveablePositions, *b.r.position.Add(dir.Delta()))
	}
	for _, pos := range moveablePositions {
		obstacle, hasObstacle := wh[pos]
		if hasObstacle {
			obstacle.move(dir, wh)
		}
	}
	delete(wh, b.l.position)
	delete(wh, b.r.position)
	b.l.position = *b.l.position.Add(dir.Delta())
	b.r.position = *b.r.position.Add(dir.Delta())
	wh[b.l.position] = b
	wh[b.r.position] = b
}

func (b *bigBox) gpsScore() int {
	return b.l.gpsScore()
}

func moveRobot(robot graph.Vector2, directions []direction.Direction, wh warehouse) {
	current := robot
	for _, d := range directions {
		next := *current.Add(d.Delta())
		obstacle, hasObstacle := wh[next]
		if !hasObstacle || obstacle.canMove(d, wh) {
			if hasObstacle {
				obstacle.move(d, wh)
			}
			current = next
		}
	}
}

func sumGps(wh warehouse) int {
	total := 0
	visited := make(map[obstacle]bool)
	for _, t := range wh {
		if !visited[t] {
			total += t.gpsScore()
			visited[t] = true
		}
	}
	return total
}

func (d *Day15) part1() int {
	wh := maps.Clone(d.warehouse)
	moveRobot(d.robotStart, d.directions, wh)
	return sumGps(wh)
}

func (d *Day15) part2() int {
	wh := biggerWarehouse(d.warehouse)
	newStart := *graph.NewVector2(d.robotStart.X*2, d.robotStart.Y)
	moveRobot(newStart, d.directions, wh)
	return sumGps(wh)
}

func biggerWarehouse(wh warehouse) warehouse {
	bigger := make(warehouse)
	for pos, t := range wh {
		newPos := *graph.NewVector2(pos.X*2, pos.Y)
		nextPos := *newPos.Add(direction.Right.Delta())
		switch t.(type) {
		case *wall:
			bigger[newPos] = newWall(newPos)
			bigger[nextPos] = newWall(nextPos)
		case *box:
			bb := newBigBox(newPos)
			bigger[bb.l.position] = bb
			bigger[bb.r.position] = bb
		}
	}
	return bigger
}
