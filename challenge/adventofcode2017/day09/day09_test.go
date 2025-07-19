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
			name: "{}",
			data: `{}`,
			want: 1},
		{
			name: "{{{}}}",
			data: `{{{}}}`,
			want: 6},
		{
			name: "{{},{}}",
			data: `{{},{}}`,
			want: 5},
		{
			name: "{{{},{},{{}}}}",
			data: `{{{},{},{{}}}}`,
			want: 16},
		{
			name: "{<a>,<a>,<a>,<a>}",
			data: `{<a>,<a>,<a>,<a>}`,
			want: 1},
		{
			name: "{{<ab>},{<ab>},{<ab>},{<ab>}}",
			data: `{{<ab>},{<ab>},{<ab>},{<ab>}}`,
			want: 9},
		{
			name: "{{<!!>},{<!!>},{<!!>},{<!!>}}",
			data: `{{<!!>},{<!!>},{<!!>},{<!!>}}`,
			want: 9},
		{
			name: "{{<a!>},{<a!>},{<a!>},{<ab>}}",
			data: `{{<a!>},{<a!>},{<a!>},{<ab>}}`,
			want: 3},
		{
			name: "{!}}",
			data: `{!}}`,
			want: 1,
		},
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
			name: "<>",
			data: `<>`,
			want: 0},
		{
			name: "<random characters>",
			data: `<random characters>`,
			want: 17},
		{
			name: "<<<<>",
			data: `<<<<>`,
			want: 3},
		{
			name: "<{!>}>",
			data: `<{!>}>`,
			want: 2},
		{
			name: "<!!>",
			data: `<!!>`,
			want: 0},
		{
			name: "<!!!>>",
			data: `<!!!>>`,
			want: 0},
		{
			name: `<{o"i!a,<{i<a>`,
			data: `<{o"i!a,<{i<a>`,
			want: 10},
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
