package day12

import (
	"adventofcode2024/util/ints"
	"maps"
	"unicode"
)

type Day12 struct {
	data []string
}

type node struct {
	id         int
	neighbours map[*node]bool
}

type nodeSet map[*node]bool

func newDay12(data []string) *Day12 {
	return &Day12{data: data}
}

func (d *Day12) part1() int {
	nodesById := d.createGraph()
	start := nodesById[0]
	group := findGroup(start)
	return len(group)
}

func (d *Day12) part2() int {
	nodesById := d.createGraph()
	var groups []nodeSet
	visited := make(nodeSet)

	for _, n := range nodesById {
		if visited[n] {
			continue
		}
		group := findGroup(n)
		groups = append(groups, group)
		maps.Copy(visited, group)
	}
	return len(groups)
}

func (d *Day12) createGraph() map[int]*node {
	nodesById := make(map[int]*node)
	addOrCreate := func(id int) *node {
		n, ok := nodesById[id]
		if !ok {
			n = &node{id: id, neighbours: make(nodeSet)}
			nodesById[id] = n
		}
		return n
	}

	for _, line := range d.data {
		fields := ints.ToInts(line, func(r rune) bool { return !unicode.IsNumber(r) })
		parent := addOrCreate(fields[0])
		for _, childId := range fields[1:] {
			child := addOrCreate(childId)
			parent.neighbours[child] = true
			child.neighbours[parent] = true
		}
	}
	return nodesById
}

func findGroup(start *node) nodeSet {
	visited := nodeSet{start: true}
	queue := []*node{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for n := range current.neighbours {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	return visited
}
