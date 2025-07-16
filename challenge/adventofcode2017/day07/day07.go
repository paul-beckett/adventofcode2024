package day07

import (
	"strconv"
	"strings"
	"unicode"
)

type Day07 struct {
	towersByName map[string]tower
}

func newDay07(data []string) *Day07 {
	towersByName := make(map[string]tower)
	for _, line := range data {
		parts := strings.Split(line, " -> ")
		nameAndWeight := strings.FieldsFunc(parts[0], func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) })
		name := nameAndWeight[0]
		weight, _ := strconv.Atoi(nameAndWeight[1])
		var children []string
		if len(parts) > 1 {
			children = strings.Split(parts[1], ", ")
		}

		towersByName[name] = tower{name: name, weight: weight, children: children}
	}
	return &Day07{
		towersByName: towersByName,
	}
}

type tower struct {
	name     string
	weight   int
	children []string
}

func bottomProgram(towersByName map[string]tower) tower {
	names := make(map[string]bool)
	for name := range towersByName {
		names[name] = true
	}

	for _, tower := range towersByName {
		for _, child := range tower.children {
			delete(names, child)
		}
	}

	for name := range names {
		return towersByName[name]
	}
	return tower{}
}

func (d *Day07) part1() string {
	return bottomProgram(d.towersByName).name
}

func (d *Day07) part2() int {
	panic("not implemented")
}
