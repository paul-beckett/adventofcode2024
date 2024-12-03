package day03

import (
	"regexp"
	"strconv"
	"strings"
)

type Day03 struct {
	data []string
}

func newDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) part1() int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0
	for _, row := range d.data {
		rowMatches := re.FindAllStringSubmatch(row, -1)
		for i := range rowMatches {
			a, _ := strconv.Atoi(rowMatches[i][1])
			b, _ := strconv.Atoi(rowMatches[i][2])
			total += a * b
		}
	}

	return total
}

func (d *Day03) part2() int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0

	s := strings.Join(d.data, "")
	rowMatches := re.FindAllStringSubmatchIndex(s, -1)
	/*
		in each match 0,1 is full string. 2,3 is first num and 4,5 is second num.
	*/
	for i := range rowMatches {
		lastEnabled := strings.LastIndex(s[:rowMatches[i][0]], "do()")
		lastDisabled := strings.LastIndex(s[:rowMatches[i][0]], "don't()")
		if lastDisabled > lastEnabled {
			continue
		}

		a, _ := strconv.Atoi(s[rowMatches[i][2]:rowMatches[i][3]])
		b, _ := strconv.Atoi(s[rowMatches[i][4]:rowMatches[i][5]])
		total += a * b
	}

	return total
}
