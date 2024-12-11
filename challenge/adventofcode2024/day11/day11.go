package day11

import (
	"adventofcode2024/util/ints"
	"math"
	"unicode"
)

type stone int
type stoneCounts map[stone]int

type Day11 struct {
	stones stoneCounts
}

func newDay11(data []string) *Day11 {
	nums := ints.ToInts(data[0], func(r rune) bool {
		return unicode.IsSpace(r)
	})
	stones := make(map[stone]int)
	for _, s := range nums {
		stones[stone(s)]++
	}

	return &Day11{stones: stones}
}

func (sc stoneCounts) blink() stoneCounts {
	next := make(map[stone]int)
	for s, count := range sc {
		switch {
		case s == 0:
			next[1] += count
		case s.digitCount()%2 == 0:
			l, r := s.split()
			next[l] += count
			next[r] += count
		default:
			next[s*2024] += count
		}
	}
	return next
}

func (sc stoneCounts) blinkN(n int) stoneCounts {
	current := sc
	for i := 0; i < n; i++ {
		current = current.blink()
	}
	return current
}

func (sc stoneCounts) count() int {
	total := 0
	for _, count := range sc {
		total += count
	}
	return total
}

func (s stone) digitCount() int {
	return int(math.Log10(float64(s))) + 1
}

func (s stone) split() (stone, stone) {
	rhs := stone(0)
	digits := s.digitCount() / 2
	for i := 0; i < digits; i++ {
		rDigit := s % 10
		rhs += stone(math.Pow(float64(10), float64(i))) * rDigit
		s /= 10
	}
	return s, rhs
}

func (d *Day11) part1() int {
	return d.stones.blinkN(25).count()
}

func (d *Day11) part2() int {
	return d.stones.blinkN(75).count()
}
