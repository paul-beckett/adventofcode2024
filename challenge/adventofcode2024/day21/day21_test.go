package day21

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
			data: `029A
980A
179A
456A
379A`,
			want: 126384},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay21(strings.Split(testCase.data, "\n")).part1()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	//no example test for part2
	t.SkipNow()
}
