package day10

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var testCases = []struct {
		name     string
		data     string
		listSize int
		want     int
	}{
		{
			name:     "example 1",
			data:     `3, 4, 1, 5`,
			listSize: 5,
			want:     12},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay10(strings.Split(testCase.data, "\n")).part1Size(testCase.listSize)
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
		want string
	}{
		{
			name: "The empty string",
			data: ``,
			want: `a2582a3a0e66e6e86e3812dcb672a272`},
		{
			name: "AoC 2017",
			data: `AoC 2017`,
			want: `33efeb34ea91902bb2f59c9920caa6cd`},
		{
			name: "1,2,3",
			data: `1,2,3`,
			want: `3efbe78a8d82f29979031a4aa0b16a9d`},
		{
			name: "1,2,4",
			data: `1,2,4`,
			want: `63960835bcdc130f0b66d7ff4f6a5a8e`},
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
