package day13

import (
	"adventofcode2024/util/ints"
	"unicode"
)

type Day13 struct {
	data []string
}

func newDay13(data []string) *Day13 {
	return &Day13{data: data}
}

func (d *Day13) part1() int {
	scanners := make(map[int]int)
	for _, line := range d.data {
		fields := ints.ToInts(line, func(r rune) bool { return !unicode.IsNumber(r) })
		scanners[fields[0]] = fields[1]
	}

	severity := 0
	for d, s := range scanners {
		if d%((s-1)*2) == 0 {
			severity += d * s
		}
	}
	return severity
}

func (d *Day13) part2() int {
	scanners := make(map[int]int)
	for _, line := range d.data {
		fields := ints.ToInts(line, func(r rune) bool { return !unicode.IsNumber(r) })
		scanners[fields[0]] = fields[1]
	}

	isHit := func(offset int) bool {
		for layer, s := range scanners {
			time := layer + offset
			if time%((s-1)*2) == 0 {
				return true
			}
		}
		return false
	}

	time := 0
	for {
		if !isHit(time) {
			return time
		}
		time++
	}
}
