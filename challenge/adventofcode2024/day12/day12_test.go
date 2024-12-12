package day12

import (
	"adventofcode2024/util/graph"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var testCases = []struct {
		name string
		data string
		want int
	}{
		{
			name: "example 1",
			data: `AAAA
BBCD
BBCC
EEEC`,
			want: 140,
		},
		{
			name: "example 2",
			data: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 772,
		},
		{
			name: "larger example",
			data: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1930,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay12(strings.Split(testCase.data, "\n")).part1()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var testCases = []struct {
		name string
		data string
		want int
	}{
		{
			name: "example 1",
			data: `AAAA
BBCD
BBCC
EEEC`,
			want: 80,
		},
		{
			name: "example 2",
			data: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 436,
		},
		{
			name: "example 3",
			data: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236,
		},
		{
			name: "example 4",
			data: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			want: 368,
		},
		{
			name: "larger example",
			data: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			want: 1206,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay12(strings.Split(testCase.data, "\n")).part2()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestRegionSides(t *testing.T) {
	var testCases = []struct {
		name string
		r    region
		want int
	}{
		{
			name: "single cell",
			r:    region{*graph.NewVector2(0, 0): true},
			want: 4,
		},
		{
			name: "straight line",
			r:    region{*graph.NewVector2(0, 0): true, *graph.NewVector2(1, 0): true},
			want: 4,
		},
		{
			name: "box",
			r: region{
				*graph.NewVector2(0, 0): true, *graph.NewVector2(1, 0): true,
				*graph.NewVector2(0, 1): true, *graph.NewVector2(1, 1): true,
			},
			want: 4,
		},
		{
			name: "E",
			r: region{
				*graph.NewVector2(0, 0): true, *graph.NewVector2(1, 0): true, *graph.NewVector2(2, 0): true,
				*graph.NewVector2(0, 1): true,
				*graph.NewVector2(0, 2): true, *graph.NewVector2(1, 2): true,
				*graph.NewVector2(0, 3): true,
				*graph.NewVector2(0, 4): true, *graph.NewVector2(1, 4): true, *graph.NewVector2(2, 4): true,
			},
			want: 12,
		},
		{
			name: "box with hole",
			r: region{
				*graph.NewVector2(0, 0): true, *graph.NewVector2(1, 0): true, *graph.NewVector2(2, 0): true,
				*graph.NewVector2(0, 1): true, *graph.NewVector2(2, 1): true,
				*graph.NewVector2(0, 2): true, *graph.NewVector2(1, 2): true, *graph.NewVector2(2, 2): true,
			},
			want: 8,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.r.sides()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
