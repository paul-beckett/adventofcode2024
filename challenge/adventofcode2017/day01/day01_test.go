package day01

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
			data: `1122`,
			want: 3},
		{
			name: "example 2",
			data: `1111`,
			want: 4},
		{
			name: "example 3",
			data: `1234`,
			want: 0},
		{
			name: "example 1",
			data: `91212129`,
			want: 9},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay01(strings.Split(testCase.data, "\n")).part1()
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
			data: `1212`,
			want: 6},
		{
			name: "example 2",
			data: `1221`,
			want: 0},
		{
			name: "example 3",
			data: `123425`,
			want: 4},
		{
			name: "example 4",
			data: `123123`,
			want: 12},
		{
			name: "example 5",
			data: `12131415`,
			want: 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := newDay01(strings.Split(testCase.data, "\n")).part2()
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
