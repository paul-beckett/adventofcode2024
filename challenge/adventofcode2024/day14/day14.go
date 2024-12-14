package day14

import (
	"adventofcode2024/util/graph"
	"adventofcode2024/util/ints"
	"fmt"
	"unicode"
)

type Day14 struct {
	robots []*robot
	width  int
	height int
}

func newDay14(data []string) *Day14 {
	return newDay14WithSize(data, 101, 103)
}

func newDay14WithSize(data []string, width, height int) *Day14 {
	var robots []*robot
	for _, line := range data {
		nums := ints.ToInts(line, func(r rune) bool {
			return !(unicode.IsDigit(r) || r == '-')
		})
		robot := newRobot(*graph.NewVector2(nums[0], nums[1]), *graph.NewVector2(nums[2], nums[3]))
		robots = append(robots, robot)
	}

	return &Day14{robots: robots, width: width, height: height}
}

type robot struct {
	position graph.Vector2
	velocity graph.Vector2
}

func newRobot(start graph.Vector2, velocity graph.Vector2) *robot {
	return &robot{
		position: start,
		velocity: velocity,
	}
}

func (r *robot) positionAfter(seconds int, width int, height int) graph.Vector2 {
	newPosition := r.position.Add(*r.velocity.Scale(seconds))
	mod(newPosition, width, height)
	return *newPosition
}

func mod(v *graph.Vector2, width int, height int) {
	v.X = ((v.X % width) + width) % width
	v.Y = ((v.Y % height) + height) % height
}

func (d *Day14) part1() int {
	seconds := 100
	midX := d.width / 2
	midY := d.height / 2
	var positions []graph.Vector2
	for _, r := range d.robots {
		newPosition := r.positionAfter(seconds, d.width, d.height)

		if newPosition.X != midX && newPosition.Y != midY {
			positions = append(positions, newPosition)
		}
	}

	quadrants := make([]int, 4)
	for _, p := range positions {
		if p.X < midX {
			if p.Y < midY {
				quadrants[0]++
			} else {
				quadrants[1]++
			}
		} else {
			if p.Y < midY {
				quadrants[2]++
			} else {
				quadrants[3]++
			}
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func (d *Day14) part2() int {
	i := 0
	for {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		positions := make(map[graph.Vector2]bool)
		for _, r := range d.robots {
			newPosition := r.positionAfter(i, d.width, d.height)
			positions[newPosition] = true
		}
		//if allNextToOther(positions) {
		//	break
		//}
		if len(positions) == len(d.robots) {
			fmt.Printf("\ni=%d\n%v\n\n", i, d.draw(positions))
			break
		}

		i++
		if i > 100_000 {
			break
		}
	}
	return i
}

func (d *Day14) draw(positions map[graph.Vector2]bool) string {
	runes := make([][]rune, d.height)
	for i := 0; i < len(runes); i++ {
		row := make([]rune, d.width)
		for j := 0; j < len(row); j++ {
			row[j] = ' '
		}
		runes[i] = row
	}

	for p := range positions {
		runes[p.Y][p.X] = 'O'
	}

	s := ""
	for _, line := range runes {
		s = s + string(line) + "\n"
	}
	return s
}
