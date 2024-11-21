package graph

import (
	"fmt"
	"slices"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	v1 := Vector2{1, 5}
	v2 := Vector2{-1, 25}

	var testCases = []struct {
		start, end Vector2
		want       float64
	}{
		{v1, v2, 22.0},
		{v2, v1, 22.0},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("from: %v, to: %v", testCase.start, testCase.end)
		t.Run(testName, func(t *testing.T) {
			got := manhattanDistance(testCase.start, testCase.end)
			if got != testCase.want {
				t.Errorf("got manhattanDistance(v1,v2) = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestVector2Graph_Neighbours(t *testing.T) {
	v1 := Vector2{0, 0}
	v2 := Vector2{0, 1}
	v3 := Vector2{0, 2}
	v4 := Vector2{100, 100}

	neighbours := map[Vector2][]Vector2{
		v1: {v2},
		v2: {v1, v3},
		v3: {v2},
	}

	graph := Vector2Graph{neighbours: neighbours}
	var testCases = []struct {
		from Vector2
		want []Vector2
	}{
		{v1, []Vector2{v2}},
		{v2, []Vector2{v1, v3}},
		{v4, nil},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("from: %v", testCase.from)
		t.Run(testName, func(t *testing.T) {
			iter := graph.Neighbours(testCase.from)
			got := slices.Collect(iter)
			if !slices.Equal(got, testCase.want) {
				t.Errorf("got graph.Neighbours(%v) = %v, want %v", testCase.from, got, testCase.want)
			}
		})
	}
}

func TestVector2Graph_FindPath(t *testing.T) {
	v1 := Vector2{0, 0}
	v2 := Vector2{0, 1}
	v3 := Vector2{0, 2}
	v4 := Vector2{0, 3}
	v5 := Vector2{100, 100}

	neighbours := map[Vector2][]Vector2{
		v1: {v2},
		v2: {v1, v3},
		v3: {v2, v4},
		v4: {v3},
	}

	graph := Vector2Graph{neighbours: neighbours}
	var testCases = []struct {
		start, end Vector2
		want       []Vector2
	}{
		{v1, v4, []Vector2{v1, v2, v3, v4}},
		{v4, v1, []Vector2{v4, v3, v2, v1}},
		{v2, v3, []Vector2{v2, v3}},
		{v1, v5, nil},
		{v5, v1, nil},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("from: %v, to: %v", testCase.start, testCase.end)
		t.Run(testName, func(t *testing.T) {
			got := graph.FindPath(testCase.start, testCase.end)
			if !slices.Equal(got, testCase.want) {
				t.Errorf("FindPath(%v, %v) = %v, want %v", testCase.start, testCase.end, got, testCase.want)
			}
		})
	}
}
