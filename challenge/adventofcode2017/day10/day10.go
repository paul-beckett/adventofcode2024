package day10

import (
	"adventofcode2024/util/ints"
	"fmt"
	"slices"
	"strings"
	"unicode"
)

type Day10 struct {
	data []string
}

func newDay10(data []string) *Day10 {
	return &Day10{data: data}
}

func (d *Day10) part1() int {
	return d.part1Size(256)
}

type sparseHash struct {
	nums    []int
	current int
	skip    int
}

func (s *sparseHash) hash(lengths []int) {
	for _, l := range lengths {

		var newOrder []int
		if s.current+l < len(s.nums) {
			before := s.nums[:s.current]
			reverse := s.nums[s.current : s.current+l]
			after := s.nums[s.current+l:]

			slices.Reverse(reverse)
			newOrder = append(newOrder, before...)
			newOrder = append(newOrder, reverse...)
			newOrder = append(newOrder, after...)
		} else {
			startCount := (s.current + l) % len(s.nums)
			numsAtStart := s.nums[:startCount]
			reverse := append(s.nums[s.current:], numsAtStart...)
			ignored := s.nums[startCount:s.current]

			slices.Reverse(reverse)
			newOrder = append(newOrder, reverse[len(reverse)-startCount:]...)
			newOrder = append(newOrder, ignored...)
			newOrder = append(newOrder, reverse[:len(reverse)-startCount]...)
		}
		s.nums = newOrder
		s.current += l + s.skip
		s.current = s.current % len(s.nums)
		s.skip++
	}
}

func (d *Day10) part1Size(size int) int {
	var list []int
	for i := range size {
		list = append(list, i)
	}

	s := sparseHash{nums: list}
	s.hash(ints.ToInts(d.data[0], func(r rune) bool { return !unicode.IsDigit(r) }))

	return s.nums[0] * s.nums[1]
}

func (d *Day10) part2() string {
	var list []int
	for i := range 256 {
		list = append(list, i)
	}

	s := sparseHash{nums: list}
	var lengths []int
	for _, c := range d.data[0] {
		lengths = append(lengths, int(c))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	for range 64 {
		s.hash(lengths)
	}

	var outputNumbers []int
	for i := range 16 {
		start := i * 16
		xor := 0
		for _, n := range s.nums[start : start+16] {
			xor ^= n
		}
		outputNumbers = append(outputNumbers, xor)
	}

	var hex []string
	for _, n := range outputNumbers {
		hex = append(hex, fmt.Sprintf("%02x", n))
	}
	return strings.Join(hex, "")
}
