package day01

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	data := strings.Split(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, "\n")

	got := newDay01(data).part1()
	want := 142

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := strings.Split(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, "\n")

	got := newDay01(data).part2()
	want := 281

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
