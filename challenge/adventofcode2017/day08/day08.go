package day08

import (
	"maps"
	"slices"
	"strconv"
	"strings"
)

type Day08 struct {
	data []string
}

var tests = map[string]func(int, int) bool{
	">":  func(a, b int) bool { return a > b },
	"<":  func(a, b int) bool { return a < b },
	">=": func(a, b int) bool { return a >= b },
	"<=": func(a, b int) bool { return a <= b },
	"==": func(a, b int) bool { return a == b },
	"!=": func(a, b int) bool { return a != b },
}
var modifiers = map[string]func(int, int) int{
	"inc": func(a, b int) int { return a + b },
	"dec": func(a, b int) int { return a - b },
}

func newDay08(data []string) *Day08 {
	return &Day08{data: data}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (d *Day08) part1() int {
	maxRegister, _ := d.maxValues()
	return maxRegister
}

func (d *Day08) maxValues() (int, int) {
	registers := make(map[string]int)
	maxValue := 0
	for _, line := range d.data {
		fields := strings.Fields(line)
		test := tests[fields[5]]
		if test(registers[fields[4]], parseInt(fields[6])) {
			modifier := modifiers[fields[1]]
			newValue := modifier(registers[fields[0]], parseInt(fields[2]))
			if newValue > maxValue {
				maxValue = newValue
			}
			registers[fields[0]] = newValue
		}
	}

	var s []int
	s = slices.AppendSeq(s, maps.Values(registers))
	return slices.Max(s), maxValue
}

func (d *Day08) part2() int {
	_, maxValue := d.maxValues()
	return maxValue
}
