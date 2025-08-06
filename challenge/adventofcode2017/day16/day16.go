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
	return string(dance(programs, d.instructions))
}

func dance(programs []rune, instructions []string) []rune {
	var result []rune
	result = append(result, programs...)
	for _, instruction := range instructions {
		if instruction[0] == 's' {
			n, err := strconv.Atoi(instruction[1:])
			if err != nil {
				panic(err)
			}
			result = spin(result, n)
		} else if instruction[0] == 'x' {
			indices := ints.ToInts(instruction[1:], unicode.IsPunct)
			swap(result, indices[0], indices[1])
		} else if instruction[0] == 'p' {
			partner(result, rune(instruction[1]), rune(instruction[3]))
		}
	}
	return result
}

func (d *Day16) part1() string {
	return d.part1Size(16)
}

func (d *Day16) part2() string {
	var programs []rune
	for i := range 16 {
		programs = append(programs, rune('a'+i))
	}
	seenAtIteration := make(map[string]int)
	var iterations []string
	runs := 1_000_000_000
	for i := range runs {
		s := string(programs)
		_, exists := seenAtIteration[s] //loop seemed to start at zero so don't need to worry about offsets
		if exists {
			return iterations[runs%i]
		}
		programs = dance(programs, d.instructions)
		seenAtIteration[s] = i
		iterations = append(iterations, s)
	}
	return string(programs)
}
