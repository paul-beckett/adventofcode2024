package day19

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"strings"
	"unicode"
)

type Day19 struct {
	data []string
}

func newDay19(data []string) *Day19 {
	return &Day19{data: data}
}

func (d *Day19) navigate() (string, int) {
	current := graph.NewVector2(0, 0)
	for i, c := range d.data[0] {
		if c == '|' {
			current.X = i
			break
		}
	}

	dir := direction.Down
	var letters []string
	inMap := func(v graph.Vector2) bool {
		return v.Y >= 0 && v.Y < len(d.data) && v.X >= 0 && v.X < len(d.data[v.Y])
	}

	value := func(v graph.Vector2) rune {
		return rune(d.data[v.Y][v.X])
	}

	steps := 0
	for inMap(*current) && value(*current) != ' ' {
		steps++
		r := value(*current)
		if unicode.IsLetter(r) {
			letters = append(letters, string(r))
		} else if r == '+' {
			ignoreDir := dir.Clockwise().Clockwise()
			for {
				next := current.Add(dir.Delta())
				if inMap(*next) && value(*next) != ' ' {
					break
				}
				dir = dir.Clockwise()
				if dir == ignoreDir {
					dir = dir.Clockwise()
				}
			}
		}
		current = current.Add(dir.Delta())
	}

	return strings.Join(letters, ""), steps
}

func (d *Day19) part1() string {
	letters, _ := d.navigate()
	return letters
}

func (d *Day19) part2() int {
	_, steps := d.navigate()
	return steps
}
