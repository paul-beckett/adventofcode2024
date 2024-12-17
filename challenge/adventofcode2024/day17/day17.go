package day17

import (
	"adventofcode2024/util/ints"
	"adventofcode2024/util/map_reduce"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Day17 struct {
	computer computer
}

func newDay17(data []string) *Day17 {
	notInt := func(r rune) bool { return !unicode.IsDigit(r) }
	var instructions []instruction
	for c := range slices.Chunk(ints.ToInts(data[4], notInt), 2) {
		i := instruction{
			opcode:  opcode(c[0]),
			operand: c[1],
		}
		instructions = append(instructions, i)
	}

	c := computer{
		a:            ints.ToInts(data[0], notInt)[0],
		b:            ints.ToInts(data[1], notInt)[0],
		c:            ints.ToInts(data[2], notInt)[0],
		instructions: instructions,
	}
	return &Day17{computer: c}
}

func (d *Day17) part1() string {
	c := d.computer
	stdOut := map_reduce.Map(c.run(), func(i int) string {
		return strconv.Itoa(i)
	})
	return strings.Join(stdOut, ",")
}

func (d *Day17) part2() int {
	instructions := d.computer.instructions

	var codes []int
	for _, i := range instructions {
		codes = append(codes, int(i.opcode))
		codes = append(codes, i.operand)
	}
	a := 0
	for i := len(codes) - 1; i >= 0; i-- {
		for {
			c := &computer{a: a, instructions: instructions}
			c.a = a
			stdOut := c.run()
			if i < len(stdOut) && slices.Equal(stdOut[i:], codes[i:]) {
				break
			}
			a += power(8, i)
		}
	}

	return a
}
