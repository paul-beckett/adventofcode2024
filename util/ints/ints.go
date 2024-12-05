package ints

import (
	"strconv"
	"strings"
)

func ToInts(line string, f func(rune) bool) []int {
	fields := strings.FieldsFunc(line, f)
	var nums []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
