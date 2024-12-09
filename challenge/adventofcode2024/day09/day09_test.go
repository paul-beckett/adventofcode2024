package day09

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
			data: `2333133121414131402`,
			want: 1928},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay09(strings.Split(testCase.data, "\n")).part1()
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
			data: `2333133121414131402`,
			want: 2858},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay09(strings.Split(testCase.data, "\n")).part2()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
