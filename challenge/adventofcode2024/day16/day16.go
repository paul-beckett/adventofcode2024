package day16

import (
	"adventofcode2024/util/collection"
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"adventofcode2024/util/map_reduce"
	"container/heap"
	"iter"
	"maps"
	"math"
	"slices"
)

type Day16 struct {
	mazeGraph mazeGraph
	start     graph.Vector2
	end       graph.Vector2
}

func newDay16(data []string) *Day16 {
	var tiles []graph.Vector2
	var start graph.Vector2
	var end graph.Vector2
	for y, line := range data {
		for x, c := range line {
			if c == '#' {
				continue
			}
			pos := *graph.NewVector2(x, y)
			tiles = append(tiles, pos)
			if c == 'S' {
				start = pos
			} else if c == 'E' {
				end = pos
			}
		}
	}
	return &Day16{mazeGraph: *newMazeGraph(tiles), start: start, end: end}
}

type mazeNode struct {
	position  graph.Vector2
	direction direction.Direction
}

func newMazeNode(position graph.Vector2, direction direction.Direction) *mazeNode {
	return &mazeNode{position: position, direction: direction}
}

type mazeGraph struct {
	neighbours map[mazeNode][]mazeNode
}

func newMazeGraph(tiles []graph.Vector2) *mazeGraph {
	mazeNodes := make(map[mazeNode]bool)
	for _, t := range tiles {
		for _, d := range direction.Cardinals {
			mazeNodes[*newMazeNode(t, d)] = true
		}
	}
	neighbours := make(map[mazeNode][]mazeNode)
	for m := range mazeNodes {
		var n []mazeNode
		n = append(n, *newMazeNode(*m.position.Add(m.direction.Delta()), m.direction))
		n = append(n, *newMazeNode(m.position, m.direction.Clockwise()))
		n = append(n, *newMazeNode(m.position, m.direction.AntiClockwise()))
		n = map_reduce.Filter(n, func(node mazeNode) bool {
			return mazeNodes[node]
		})
		neighbours[m] = n
	}
	return &mazeGraph{
		neighbours: neighbours,
	}
}

func (m *mazeGraph) Neighbours(n mazeNode) iter.Seq[mazeNode] {
	return func(yield func(node mazeNode) bool) {
		for _, neighbour := range m.neighbours[n] {
			if !yield(neighbour) {
				return
			}
		}
	}
}

func moveCost(p, q mazeNode) float64 {
	if p.direction != q.direction {
		return 1000
	} else {
		return 1
	}
}

type pathRecord struct {
	from mazeNode
	to   mazeNode
	cost float64
}

func newPathRecord(from mazeNode, to mazeNode, cost float64) *pathRecord {
	return &pathRecord{from: from, to: to, cost: cost}
}

func (m *mazeGraph) expandFrom(start mazeNode, costF func(mazeNode, mazeNode) float64) map[mazeNode]pathRecord {
	visited := make(map[mazeNode]bool)
	pq := &collection.PriorityQueue[pathRecord]{}
	heap.Init(pq)
	var from mazeNode
	heap.Push(pq, &collection.Item[pathRecord]{
		Value: *newPathRecord(from, start, 0),
	})

	pathRecords := make(map[mazeNode]pathRecord)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*collection.Item[pathRecord]).Value
		if visited[current.to] {
			continue
		}
		visited[current.to] = true

		pathRecords[current.to] = current
		for _, neighbour := range m.neighbours[current.to] {
			moveCost := costF(current.to, neighbour)
			pr := *newPathRecord(current.to, neighbour, current.cost+moveCost)
			heap.Push(pq, &collection.Item[pathRecord]{
				Value:    pr,
				Priority: -pr.cost,
			})
		}
	}
	return pathRecords
}

func (d *Day16) part1() int {
	pathRecords := d.mazeGraph.expandFrom(*newMazeNode(d.start, direction.Right), moveCost)
	bestCost := math.MaxFloat64
	for _, dir := range direction.Cardinals {
		end := *newMazeNode(d.end, dir)
		pr, ok := pathRecords[end]
		if ok {
			bestCost = min(bestCost, pr.cost)
		}
	}

	return int(bestCost)
}

func (m *mazeGraph) nodesOnBestPaths(pathRecords map[mazeNode]pathRecord, end mazeNode) []mazeNode {
	neighboursByEnd := make(map[mazeNode][]mazeNode)
	for node, neighbours := range m.neighbours {
		for _, n := range neighbours {
			neighboursByEnd[n] = append(neighboursByEnd[n], node)
		}
	}
	visited := make(map[mazeNode]bool)

	var queue []pathRecord
	queue = append(queue, pathRecords[end])
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current.to] {
			continue
		}
		visited[current.to] = true
		if &current.from == nil {
			continue
		}
		for _, neighbour := range neighboursByEnd[current.to] {
			if pathRecords[neighbour].cost+moveCost(neighbour, current.to) == current.cost {
				queue = append(queue, pathRecords[neighbour])
			}
		}
	}
	return slices.Collect(maps.Keys(visited))
}

func (d *Day16) part2() int {
	mg := d.mazeGraph
	pathRecords := mg.expandFrom(*newMazeNode(d.start, direction.Right), moveCost)
	var bestEnds []pathRecord
	bestCost := math.MaxFloat64
	for _, dir := range direction.Cardinals {
		end := *newMazeNode(d.end, dir)
		pr, ok := pathRecords[end]
		if ok {
			if pr.cost <= bestCost {
				if pr.cost < bestCost {
					bestEnds = nil
					bestCost = pr.cost
				}
				bestEnds = append(bestEnds, pr)
			}
		}
	}

	visited := make(map[graph.Vector2]bool)
	for _, end := range bestEnds {
		for _, node := range mg.nodesOnBestPaths(pathRecords, end.to) {
			visited[node.position] = true
		}
	}

	return len(visited)
}
