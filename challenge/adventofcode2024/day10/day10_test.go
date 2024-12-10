package day10

import (
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
			data: `0123
1234
8765
9876`,
			want: 1,
		},
		{
			name: "example 2",
			data: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 36,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay10(strings.Split(testCase.data, "\n")).part1()
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
			data: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			want: 81},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay10(strings.Split(testCase.data, "\n")).part2()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
