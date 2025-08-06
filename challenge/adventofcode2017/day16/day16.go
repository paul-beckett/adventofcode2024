package day16

import (
	"adventofcode2024/util/ints"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Day16 struct {
	instructions []string
}

func newDay16(data []string) *Day16 {
	return &Day16{instructions: strings.Split(data[0], ",")}
}

func spin(programs []rune, i int) []rune {
	var spun []rune
	spun = append(spun, programs[len(programs)-i:]...)
	return append(spun, programs[:len(programs)-i]...)
}

func swap(programs []rune, i, j int) {
	programs[i], programs[j] = programs[j], programs[i]
}

func partner(programs []rune, a, b rune) {
	i := slices.Index(programs, a)
	j := slices.Index(programs, b)
	swap(programs, i, j)
}

func (d *Day16) part1Size(size int) string {
	var programs []rune
	for i := range size {
		programs = append(programs, rune('a'+i))
	}

	for _, instruction := range d.instructions {
		if instruction[0] == 's' {
			n, err := strconv.Atoi(instruction[1:])
			if err != nil {
				panic(err)
			}
			programs = spin(programs, n)
		} else if instruction[0] == 'x' {
			indices := ints.ToInts(instruction[1:], unicode.IsPunct)
			swap(programs, indices[0], indices[1])
		} else if instruction[0] == 'p' {
			partner(programs, rune(instruction[1]), rune(instruction[3]))
		}
	}
	return string(programs)
}

func (d *Day16) part1() string {
	return d.part1Size(16)
}

func (d *Day16) part2() int {
	panic("not implemented")
}
