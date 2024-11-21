package graph

import (
	"github.com/fzipp/astar"
	"iter"
	"math"
)

type Vector2Graph struct {
	neighbours map[Vector2][]Vector2
}

func (v *Vector2Graph) Neighbours(n Vector2) iter.Seq[Vector2] {
	return func(yield func(Vector2) bool) {
		for _, neighbour := range v.neighbours[n] {
			if !yield(neighbour) {
				return
			}
		}
	}
}

func manhattanDistance(p, q Vector2) float64 {
	return math.Abs(float64(p.X-q.X)) + math.Abs(float64(p.Y-q.Y))
}

func (v *Vector2Graph) FindPath(start, goal Vector2) []Vector2 {
	return astar.FindPath(v, start, goal, manhattanDistance, manhattanDistance)
}
