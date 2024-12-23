package day23

import (
	"maps"
	"slices"
	"strings"
)

type Day23 struct {
	connections map[string]set
}

type set map[string]bool

func newDay23(data []string) *Day23 {
	connections := make(map[string]set)
	for _, row := range data {
		nodes := strings.Split(row, "-")
		lhs := nodes[0]
		rhs := nodes[1]
		if connections[lhs] == nil {
			connections[lhs] = make(set)
		}
		if connections[rhs] == nil {
			connections[rhs] = make(set)
		}
		connections[lhs][rhs] = true
		connections[rhs][lhs] = true
	}
	return &Day23{connections: connections}
}

func makeNetwork(connections map[string]set) []set {
	visited := make(set)
	var network []set

	for node, n := range connections {
		visited[node] = true
		var neighbours []string
		neighbours = slices.AppendSeq(neighbours, maps.Keys(n))
		for i := 0; i < len(neighbours)-1; i++ {
			l := neighbours[i]
			if visited[l] {
				continue
			}
			for j := i + 1; j < len(neighbours); j++ {
				r := neighbours[j]
				if visited[r] {
					continue
				}
				if connections[l][r] {
					network = append(network, set{
						node: true,
						l:    true,
						r:    true,
					})
				}
			}
		}
	}
	return network
}

func (s *set) intersect(other set) set {
	intersect := make(set)
	for v := range *s {
		if other[v] {
			intersect[v] = true
		}
	}
	return intersect
}

func (s *set) makeKey() string {
	var key []string
	key = slices.AppendSeq(key, maps.Keys(*s))
	slices.Sort(key)
	return strings.Join(key, ",")
}

func (d *Day23) biggestGroup(including set, neighbours set, cache map[string]set) set {
	key := including.makeKey()
	best, ok := cache[key]
	if ok {
		return best
	}

	biggest := including
	for neighbour := range neighbours {
		next := d.connections[neighbour]
		next = next.intersect(neighbours)
		child := make(set)
		maps.Copy(child, including)
		child[neighbour] = true
		found := d.biggestGroup(child, next, cache)
		if len(found) > len(biggest) {
			biggest = found
		}
	}
	cache[key] = biggest
	return biggest
}

func (d *Day23) part1() int {
	total := 0
	network := makeNetwork(d.connections)
	for _, nodes := range network {
		for node := range nodes {
			if strings.HasPrefix(node, "t") {
				total++
				break
			}
		}
	}

	return total
}

func (d *Day23) part2() string {
	var biggestGroup set
	cache := make(map[string]set)
	for node := range d.connections {
		found := d.biggestGroup(set{node: true}, d.connections[node], cache)
		if len(found) > len(biggestGroup) {
			biggestGroup = found
		}
	}

	var password []string
	password = slices.AppendSeq(password, maps.Keys(biggestGroup))
	slices.Sort(password)
	return strings.Join(password, ",")
}
