package day05

import (
	"adventofcode2024/util/map_reduce"
	"strconv"
)

type Day05 struct {
	data []string
}

func newDay05(data []string) *Day05 {
	return &Day05{data: data}
}

func (d *Day05) part1() int {
	return d.findExit(func(i int) int {
		return i + 1
	})
}

func (d *Day05) findExit(modifier func(int) int) int {
	instructions := map_reduce.Map(d.data, func(s string) int {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return n
	})
	count := 0
	i := 0
	for {
		if i < 0 || i >= len(d.data) {
			return count
		}
		count++
		next := i + instructions[i]
		instructions[i] = modifier(instructions[i])
		i = next
	}
}

func (d *Day05) part2() int {
	return d.findExit(func(i int) int {
		if i >= 3 {
			return i - 1
		} else {
			return i + 1
		}
	})
}
