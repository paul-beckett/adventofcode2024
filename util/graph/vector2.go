package graph

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x int, y int) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func (v Vector2) Add(o Vector2) *Vector2 {
	return NewVector2(v.X+o.X, v.Y+o.Y)
}

func (v Vector2) Minus(o Vector2) *Vector2 { return NewVector2(v.X-o.X, v.Y-o.Y) }

func (v Vector2) Scale(factor int) *Vector2 {
	return NewVector2(v.X*factor, v.Y*factor)
}

func (v Vector2) ManhattanDistance(o Vector2) int {
	dx := math.Abs(float64(o.X - v.X))
	dy := math.Abs(float64(o.Y - v.Y))
	return int(dx + dy)
}

func (v Vector2) String() string {
	return fmt.Sprintf("(x=%d, y=%d)", v.X, v.Y)
}
