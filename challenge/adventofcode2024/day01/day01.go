package day01

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day01 struct {
	data []string
}

func newDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func (d *Day01) part1() int {
	var lhs []int
	var rhs []int

	for _, row := range d.data {
		fields := strings.Fields(row)
		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])
		lhs = append(lhs, l)
		rhs = append(rhs, r)
	}
	slices.Sort(lhs)
	slices.Sort(rhs)

	totalDiff := 0
	for i := 0; i < len(lhs); i++ {
		totalDiff += int(math.Abs(float64(lhs[i] - rhs[i])))
	}
	return totalDiff
}

func (d *Day01) part2() int {
	lhsCounts := make(map[string]int)
	rhsCounts := make(map[string]int)

	for _, row := range d.data {
		fields := strings.Fields(row)
		lhsCounts[fields[0]]++
		rhsCounts[fields[1]]++
	}
	total := 0
	for k, count := range lhsCounts {
		lhs, _ := strconv.Atoi(k)
		total += (lhs * rhsCounts[k]) * count
	}
	return total
}
