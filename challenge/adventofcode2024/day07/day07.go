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

func (d *Day07) part1() int {
	calibrationSum := 0
	for _, c := range d.calibrations {
		if isPossibleEquation(c, []func(int, int) int{add, multiply}) {
			calibrationSum += c.testValue
		}
	}

	return calibrationSum
}

func isPossibleEquation(c calibration, operations []func(int, int) int) bool {
	possibles := make([]map[int]bool, len(c.equation))
	possibles[0] = map[int]bool{c.equation[0]: true}
	for i := 1; i < len(c.equation); i++ {
		next := make(map[int]bool)
		for n := range possibles[i-1] {
			for _, o := range operations {
				res := o(n, c.equation[i])
				if res <= c.testValue {
					next[res] = true
				}
			}
		}
		if len(next) == 0 {
			return false
		}
		possibles[i] = next
	}
	return possibles[len(possibles)-1][c.testValue]
}

func (d *Day07) part2() int {
	calibrationSum := 0
	for _, c := range d.calibrations {
		if isPossibleEquation(c, []func(int, int) int{add, multiply, concat}) {
			calibrationSum += c.testValue
		}
	}

	return calibrationSum
}
