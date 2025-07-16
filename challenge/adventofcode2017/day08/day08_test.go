package day08

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
			data: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`,
			want: 1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay08(strings.Split(testCase.data, "\n")).part1()
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
			data: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`,
			want: 10},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay08(strings.Split(testCase.data, "\n")).part2()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
