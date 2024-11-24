package day04

import (
	"maps"
	"math"
	"strconv"
	"strings"
)

type Day04 struct {
	data []string
}

func newDay04(data []string) *Day04 {
	return &Day04{data: data}
}

func (d *Day04) part1() int {
	total := 0
	for _, row := range d.data {
		count := countWinners(row)
		if count > 0 {
			total += int(math.Pow(2, float64(count-1)))
		}
	}
	return total
}

func countWinners(row string) int {
	game := strings.Split(strings.Split(row, ": ")[1], " | ")
	winning, _ := toInts(game[0])
	winningSet := createSet(winning)

	count := 0
	nums, _ := toInts(game[1])
	for _, num := range nums {
		if winningSet[num] {
			count++
		}
	}
	return count
}

type set map[int]bool

func createSet(nums []int) set {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	return m
}

func toInts(s string) ([]int, error) {
	var nums []int
	for _, part := range strings.Fields(s) {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func (d *Day04) part2() int {
	cardCounts := make(map[int]int)
	for rowIndex, row := range d.data {
		cardCounts[rowIndex] += 1
		count := countWinners(row)
		for i := 1; i <= count; i++ {
			cardCounts[rowIndex+i] += cardCounts[rowIndex]
		}
	}
	sum := 0
	for cardCount := range maps.Values(cardCounts) {
		sum += cardCount
	}
	return sum
}
