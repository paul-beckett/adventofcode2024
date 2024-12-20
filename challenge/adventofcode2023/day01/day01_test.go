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
		{"example 1",
			`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, 142},
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
		{"example 1",
			`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, 281},
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
