package day01

import (
	"strconv"
	"strings"
)

type Day01 struct {
	data []string
}

func newDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func captcha(digits []string, nextI int) int {
	sum := 0
	for i, n := range digits {
		next := i + nextI
		next = next % len(digits)
		if n == digits[next] {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}
	return sum
}

func (d *Day01) part1() int {
	digits := strings.Split(d.data[0], "")
	return captcha(digits, 1)
}

func (d *Day01) part2() int {
	digits := strings.Split(d.data[0], "")
	return captcha(digits, len(digits)/2)
}
