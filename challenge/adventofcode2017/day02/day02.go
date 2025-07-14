package day02

import (
	"slices"
	"strconv"
	"strings"
)

type Day02 struct {
	data []string
}

func newDay02(data []string) *Day02 {
	return &Day02{data: data}
}

func (d *Day02) part1() int {
	checksum := 0
	for _, row := range d.data {
		var nums []int
		for _, s := range strings.Fields(row) {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		diff := slices.Max(nums) - slices.Min(nums)
		if diff < 0 {
			diff *= -1
		}
		checksum += diff
	}
	return checksum
}

func (d *Day02) part2() int {
	checksum := 0
	for _, row := range d.data {
		var nums []int
		for _, s := range strings.Fields(row) {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
		slices.Sort(nums)
		for i, n := range nums[:len(nums)-1] {
			for _, m := range nums[i+1:] {
				if m%n == 0 {
					checksum += m / n
				}
			}
		}
	}
	return checksum
}
