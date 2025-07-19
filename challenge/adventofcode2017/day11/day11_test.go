package day11

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
			name: "ne,ne,ne",
			data: `ne,ne,ne`,
			want: 3},
		{
			name: "ne,ne,sw,sw",
			data: `ne,ne,sw,sw`,
			want: 0},
		{
			name: "ne,ne,s,s",
			data: `ne,ne,s,s`,
			want: 2},
		{
			name: "se,sw,se,sw,sw",
			data: `se,sw,se,sw,sw`,
			want: 3},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay11(strings.Split(testCase.data, "\n")).part1()
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
