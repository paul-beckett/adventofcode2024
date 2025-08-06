package day15

import (
	"adventofcode2024/util/map_reduce"
	"strconv"
	"strings"
)

type Day15 struct {
	a int
	b int
}

func newDay15(data []string) *Day15 {
	nums := map_reduce.Map(data, func(s string) int {
		n, err := strconv.Atoi(strings.Fields(s)[4])
		if err != nil {
			panic(err)
		}
		return n
	})

	return &Day15{a: nums[0], b: nums[1]}
}

func nextNumber(seed, multiplier int) int {
	return (seed * multiplier) % 2147483647
}
func (d *Day15) part1() int {
	a := d.a
	b := d.b
	remainderOf := 65536

	count := 0
	for range 40_000_000 {
		a = nextNumber(a, 16807)
		b = nextNumber(b, 48271)

		if a%remainderOf == b%remainderOf {
			count++
		}
		//fmt.Printf("%.32b\n%.32b\n\n", a, b)
	}

	return count
}

func validNextNumber(seed, multiplier, multiple int) int {
	n := seed
	for {
		n = nextNumber(n, multiplier)
		if n%multiple == 0 {
			return n
		}
	}
}

func (d *Day15) part2() int {
	a := d.a
	b := d.b
	remainderOf := 65536

	count := 0
	for range 5_000_000 {
		a = validNextNumber(a, 16807, 4)
		b = validNextNumber(b, 48271, 8)

		if a%remainderOf == b%remainderOf {
			count++
		}
		//fmt.Printf("%.32b\n%.32b\n\n", a, b)
	}

	return count
}
