package day03

import "regexp"

type Day03 struct {
	data []string
}

func newDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) part1() int {
	re := regexp.MustCompile("[0-9]+")
	var rowMatches [][][]int
	for _, text := range d.data {
		rowMatches = append(rowMatches, re.FindAllStringSubmatchIndex(text, -1))
	}
	return -1
}

func (d *Day03) part2() int {
	panic("not implemented")
}
