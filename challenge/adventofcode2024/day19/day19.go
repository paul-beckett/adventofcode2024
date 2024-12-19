package day19

import (
	"strings"
)

type Day19 struct {
	towels  map[string]bool
	designs []string
}

func newDay19(data []string) *Day19 {
	towels := make(map[string]bool)
	for _, t := range strings.Split(data[0], ", ") {
		towels[t] = true
	}

	var designs []string
	for i := 2; i < len(data); i++ {
		designs = append(designs, data[i])
	}
	return &Day19{
		towels:  towels,
		designs: designs,
	}
}

func (d *Day19) countCombinations(s string, cache map[string]int) int {
	v, ok := cache[s]
	if ok {
		return v
	}
	count := 0
	if d.towels[s] {
		count++
	}
	for i := 1; i < len(s); i++ {
		if d.towels[s[:i]] {
			count += d.countCombinations(s[i:], cache)
		}
	}
	cache[s] = count
	return count
}

func (d *Day19) part1() int {
	total := 0
	cache := make(map[string]int)
	for i := 0; i < len(d.designs); i++ {
		if d.countCombinations(d.designs[i], cache) > 0 {
			total++
		}
	}
	return total
}

func (d *Day19) part2() int {
	total := 0
	cache := make(map[string]int)
	for i := 0; i < len(d.designs); i++ {
		total += d.countCombinations(d.designs[i], cache)
	}
	return total
}
