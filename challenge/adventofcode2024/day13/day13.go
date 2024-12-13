package day13

import (
	"adventofcode2024/util/graph"
	"adventofcode2024/util/ints"
	"adventofcode2024/util/map_reduce"
	"slices"
	"strings"
	"unicode"
)

type Day13 struct {
	machines []machine
}

func newDay13(data []string) *Day13 {

	parseButton := func(s string) graph.Vector2 {
		coords := strings.Split(s, ": ")
		xAndY := ints.ToInts(coords[1], func(r rune) bool {
			return !unicode.IsDigit(r) && r != '+' && r != '-'
		})

		return *graph.NewVector2(xAndY[0], xAndY[1])
	}
	parsePrize := func(s string) graph.Vector2 {
		coords := strings.Split(s, ": ")
		xAndY := ints.ToInts(coords[1], func(r rune) bool {
			return !unicode.IsDigit(r)
		})

		return *graph.NewVector2(xAndY[0], xAndY[1])
	}

	var machines []machine
	for game := range slices.Chunk(data, 4) {
		m := *newMachine(parseButton(game[0]), parseButton(game[1]), parsePrize(game[2]))
		machines = append(machines, m)
	}
	return &Day13{machines: machines}
}

type machine struct {
	buttonA graph.Vector2
	buttonB graph.Vector2
	prize   graph.Vector2
}

func newMachine(buttonA, buttonB, prize graph.Vector2) *machine {
	return &machine{buttonA: buttonA, buttonB: buttonB, prize: prize}
}

func (m *machine) validCombination() (a, b int, valid bool) {
	b = (m.prize.Y*m.buttonA.X - m.prize.X*m.buttonA.Y) / (m.buttonA.X*m.buttonB.Y - m.buttonA.Y*m.buttonB.X)
	a = (m.prize.X - m.buttonB.X*b) / m.buttonA.X

	valid = a*m.buttonA.X+b*m.buttonB.X == m.prize.X &&
		a*m.buttonA.Y+b*m.buttonB.Y == m.prize.Y
	return a, b, valid
}

func totalCost(machines []machine) int {
	aCost := 3
	bCost := 1
	total := 0
	for _, m := range machines {
		a, b, valid := m.validCombination()
		if valid {
			total += aCost*a + bCost*b
		}
	}
	return total
}

func (d *Day13) part1() int {
	return totalCost(d.machines)
}

func (d *Day13) part2() int {
	modifier := 10_000_000_000_000
	machines := map_reduce.Map(d.machines, func(m machine) machine {
		m.prize = *m.prize.Add(*graph.NewVector2(modifier, modifier))
		return m
	})
	return totalCost(machines)
}
