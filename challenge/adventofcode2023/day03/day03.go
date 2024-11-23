package day03

import (
	"adventofcode2024/util/map_reduce"
	"regexp"
	"strconv"
	"unicode"
)

type Day03 struct {
	data []string
}

func newDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) part1() int {
	re := regexp.MustCompile("[0-9]+")
	var parts []int
	for row, text := range d.data {
		rowMatches := re.FindAllStringSubmatchIndex(text, -1)
		for _, match := range rowMatches {
			if isPart(match[0], match[1], row, d.data) {
				n, _ := strconv.Atoi(text[match[0]:match[1]])
				parts = append(parts, n)
			}
		}
	}
	return map_reduce.Sum(parts)
}

func isPart(start, end int, row int, data []string) bool {
	for offset := max(row-1, 0); offset <= min(row+1, len(data)-1); offset++ {
		for i := max(start-1, 0); i < min(end+1, len(data[offset])-1); i++ {
			r := rune(data[offset][i])
			if !unicode.IsDigit(r) && r != '.' {
				return true
			}
		}
	}
	return false
}

func (d *Day03) part2() int {
	var gears [][]int
	for row, text := range d.data {
		for col, c := range text {
			if c == '*' {
				gears = append(gears, findNumbers(row, col, d.data))
			}
		}
	}
	gears = map_reduce.Filter(gears, func(gear []int) bool {
		return len(gear) == 2
	})
	return map_reduce.Reduce(gears, func(sum int, gear []int) int {
		return sum + gear[0]*gear[1]
	}, 0)
}

func findNumbers(row, col int, data []string) []int {
	re := regexp.MustCompile("[0-9]+")
	var numbers []int
	for offset := max(row-1, 0); offset <= min(row+1, len(data)-1); offset++ {
		numberRanges := re.FindAllStringSubmatchIndex(data[offset], -1)
		for _, number := range numberRanges {
			if col >= number[0]-1 && col <= number[1] {
				n, _ := strconv.Atoi(data[offset][number[0]:number[1]])
				numbers = append(numbers, n)
			}
		}
	}
	return numbers
}
