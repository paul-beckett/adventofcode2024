package graph

import "fmt"

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) String() string {
	return fmt.Sprintf("(x=%d, y=%d)", v.X, v.Y)
}
