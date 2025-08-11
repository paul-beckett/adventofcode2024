package day24

import (
	"adventofcode2024/util/ints"
	"unicode"
)

type Day24 struct {
	components componentsByPorts
}

type port int

type component struct {
	start port
	end   port
}

func addToMultiMap(multimap componentsByPorts, p, q port, c component) {
	m, exists := multimap[p]
	if !exists {
		m = make(map[port]component)
		multimap[p] = m
	}
	m[q] = c
}

type componentsByPorts map[port]map[port]component

func newDay24(data []string) *Day24 {
	components := make(componentsByPorts)

	for _, line := range data {
		fields := ints.ToInts(line, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		c := component{
			start: port(fields[0]),
			end:   port(fields[1]),
		}
		addToMultiMap(components, c.start, c.end, c)
		addToMultiMap(components, c.end, c.start, c)
	}
	return &Day24{components: components}
}

func (d *Day24) part1() int {
	used := make(map[component]bool)
	return strongestBridge(d.components, used, 0, 0)
}

func strongestBridge(components componentsByPorts, used map[component]bool, current port, strength int) int {
	best := strength
	for end, c := range components[current] {
		if used[c] {
			continue
		}
		used[c] = true
		best = max(best, strongestBridge(components, used, end, strength+int(current)+int(end)))
		delete(used, c)
	}
	return best
}

func longestBridge(components componentsByPorts, used map[component]bool, current port, strength int) (int, int) {
	strongest := strength
	longest := len(used)
	for end, c := range components[current] {
		if used[c] {
			continue
		}
		used[c] = true
		l, s := longestBridge(components, used, end, strength+int(current)+int(end))
		if l > longest {
			longest = l
			strongest = s
		} else if l == longest {
			strongest = max(strongest, s)
		}
		delete(used, c)
	}
	return longest, strongest
}

func (d *Day24) part2() int {
	used := make(map[component]bool)
	_, strength := longestBridge(d.components, used, 0, 0)
	return strength
}
