package day01

import (
	"adventofcode2024/util/map_reduce"
	"strings"
	"unicode"
)

type Day01 struct {
	data []string
}

func newDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func (d *Day01) part1() any {
	calibrationValues := map_reduce.Map(d.data, calibrationValue)
	return map_reduce.Sum(calibrationValues)
}

func (d *Day01) part2() any {
	wordValues := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	calibrationValues := map_reduce.Map(d.data, func(a string) int {
		return calibrationValuePart2(a, wordValues)
	})
	return map_reduce.Sum(calibrationValues)
}

func calibrationValue(s string) int {
	v := 0
	chars := []rune(s)
	for _, c := range chars {
		if unicode.IsDigit(c) {
			v += int(c-'0') * 10
			break
		}
	}

	for i := len(chars) - 1; i >= 0; i-- {
		c := chars[i]
		if unicode.IsDigit(c) {
			v += int(c - '0')
			break
		}
	}
	return v
}

func calibrationValuePart2(s string, wordValues map[string]int) int {
	positions := make([]int, len(s))
	for k, v := range wordValues {
		if i := strings.Index(s, k); i != -1 {
			positions[i] = v
		}

		if i := strings.LastIndex(s, k); i != -1 {
			positions[i] = v
		}
	}

	found := map_reduce.Filter(positions, func(i int) bool {
		return i != 0
	})

	return found[0]*10 + found[len(found)-1]
}
