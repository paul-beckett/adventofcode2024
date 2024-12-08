package day08

import (
	"adventofcode2024/util/graph"
	"unicode"
)

type Day08 struct {
	allAntennas map[rune]antennas
	inBounds    func(graph.Vector2) bool
}

func newDay08(data []string) *Day08 {
	lower := graph.NewVector2(0, 0)
	upper := graph.NewVector2(len(data[0])-1, len(data)-1)

	inBounds := func(v graph.Vector2) bool {
		return v.X >= lower.X && v.Y >= lower.Y && v.X <= upper.X && v.Y <= upper.Y
	}

	allAntennas := make(map[rune]antennas)
	for y, row := range data {
		for x, c := range row {
			if unicode.IsDigit(c) || unicode.IsLetter(c) {
				allAntennas[c] = append(allAntennas[c], *graph.NewVector2(x, y))
			}
		}
	}

	return &Day08{
		inBounds:    inBounds,
		allAntennas: allAntennas,
	}
}

type antennas []graph.Vector2

func (a *antennas) findAntiNodes(inBounds func(graph.Vector2) bool, single bool) []graph.Vector2 {
	var antiNodes []graph.Vector2
	for i, lhs := range *a {
		for j, rhs := range *a {
			if i == j {
				continue
			}
			delta := *lhs.Minus(rhs)
			antiNode := lhs
			for {
				antiNode = *antiNode.Add(delta)
				if !inBounds(antiNode) {
					break
				}
				antiNodes = append(antiNodes, antiNode)
				if single {
					break
				}
			}
		}
	}
	return antiNodes
}

func (d *Day08) allAntiNodes(single bool) map[graph.Vector2]bool {
	allAntiNodes := make(map[graph.Vector2]bool)
	for _, a := range d.allAntennas {
		antiNodes := a.findAntiNodes(d.inBounds, single)
		for _, node := range antiNodes {
			allAntiNodes[node] = true
		}
	}
	return allAntiNodes
}

func (d *Day08) part1() int {
	return len(d.allAntiNodes(true))
}

func (d *Day08) part2() int {
	antiNodes := d.allAntiNodes(false)
	for _, a := range d.allAntennas {
		for _, antenna := range a {
			antiNodes[antenna] = true
		}
	}
	return len(antiNodes)
}
