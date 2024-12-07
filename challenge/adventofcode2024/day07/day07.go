package day07

import (
	"adventofcode2024/util/ints"
	"math"
	"unicode"
)

type Day07 struct {
	calibrations []calibration
}

type calibration struct {
	testValue int
	equation  []int
}

func newCalibration(t int, e []int) *calibration {
	return &calibration{
		testValue: t,
		equation:  e,
	}
}

type operator func(int, int) int

func newDay07(data []string) *Day07 {
	var calibrations []calibration
	for _, row := range data {
		nums := ints.ToInts(row, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		calibrations = append(calibrations, *newCalibration(nums[0], nums[1:]))
	}
	return &Day07{calibrations: calibrations}
}

var (
	add = func(a int, b int) int {
		return a + b
	}
	multiply = func(a int, b int) int {
		return a * b
	}
	concat = func(a int, b int) int {
		bLen := int(math.Log10(float64(b))) + 1
		return a*int(math.Pow10(bLen)) + b
	}
)

func (d *Day07) calibrationSum(operators []operator) int {
	calibrationSum := 0
	for _, c := range d.calibrations {
		if c.isPossible(c.equation[0], 1, operators) {
			calibrationSum += c.testValue
		}
	}

	return calibrationSum
}

func (c *calibration) isPossible(total int, i int, operators []operator) bool {
	if i >= len(c.equation) {
		return total == c.testValue
	}

	for _, o := range operators {
		possible := c.isPossible(o(total, c.equation[i]), i+1, operators)
		if possible {
			return true
		}
	}
	return false
}

func (d *Day07) part1() int {
	return d.calibrationSum([]operator{add, multiply})
}

func (d *Day07) part2() int {
	return d.calibrationSum([]operator{add, multiply, concat})
}
