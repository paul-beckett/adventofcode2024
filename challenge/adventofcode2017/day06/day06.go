package day06

import (
	"adventofcode2024/util/ints"
	"cmp"
	"fmt"
	"slices"
	"unicode"
)

type Day06 struct {
	data []string
}

func newDay06(data []string) *Day06 {
	return &Day06{data: data}
}

func (d *Day06) part1() int {
	size, _ := findCycle(ints.ToInts(d.data[0], unicode.IsSpace))
	return size
}

func findCycle(dataBanks []int) (int, []int) {
	var indices []int
	for i := range dataBanks {
		indices = append(indices, i)
	}

	cycle := 0
	seen := map[string]bool{toKey(dataBanks): true}
	for {
		cycle++
		slices.SortFunc(indices, func(a, b int) int {
			return cmp.Or(cmp.Compare(dataBanks[b], dataBanks[a]), cmp.Compare(a, b))
		})
		index := indices[0]
		n := dataBanks[index]
		dataBanks[index] = 0
		for n > 0 {
			index++
			if index >= len(dataBanks) {
				index = 0
			}
			dataBanks[index]++
			n--
		}
		key := toKey(dataBanks)
		if seen[key] {
			return cycle, dataBanks
		}
		seen[key] = true
	}
}

func (d *Day06) part2() int {
	_, dataBanks := findCycle(ints.ToInts(d.data[0], unicode.IsSpace))
	size, _ := findCycle(dataBanks)
	return size
}

func toKey(nums []int) string {
	return fmt.Sprint(nums)
}
