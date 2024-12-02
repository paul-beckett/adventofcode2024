package day02

import (
	"adventofcode2024/util/map_reduce"
	"strconv"
	"strings"
)

type Day02 struct {
	data []string
}

func newDay02(data []string) *Day02 {
	return &Day02{data: data}
}

func findUnsafe(reports []int) int {
	decreasing := reports[0] > reports[1]
	safeDecrease := func(a, b int) bool {
		return a > b && (a-b) <= 3
	}
	safeIncrease := func(a, b int) bool { return safeDecrease(b, a) }
	safe := safeDecrease
	if !decreasing {
		safe = safeIncrease
	}

	for i := 1; i < len(reports); i++ {
		if !safe(reports[i-1], reports[i]) {
			return i
		}
	}

	return -1
}

func hasMaxOneUnsafe(reports []int) bool {
	unsafeIndex := findUnsafe(reports)
	return unsafeIndex == -1 ||
		findUnsafe(remove(reports, unsafeIndex-1)) == -1 ||
		findUnsafe(remove(reports, unsafeIndex)) == -1 ||
		findUnsafe(remove(reports, 0)) == -1
}

func remove(slice []int, index int) []int {
	var s []int
	for i, n := range slice {
		if i == index {
			continue
		}
		s = append(s, n)
	}
	return s
}

func toInts(s string) []int {
	var nums []int
	for _, field := range strings.Fields(s) {
		n, _ := strconv.Atoi(field)
		nums = append(nums, n)
	}
	return nums
}

func (d *Day02) part1() int {
	reports := map_reduce.Map(d.data, toInts)
	return len(map_reduce.Filter(reports, func(reports []int) bool {
		return findUnsafe(reports) == -1
	}))
}

func (d *Day02) part2() int {
	reports := map_reduce.Map(d.data, toInts)
	return len(map_reduce.Filter(reports, hasMaxOneUnsafe))
}
