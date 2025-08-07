package day17

import (
	"strconv"
)

type Day17 struct {
	data []string
}

func newDay17(data []string) *Day17 {
	return &Day17{data: data}
}

func (d *Day17) part1() int {
	steps, err := strconv.Atoi(d.data[0])
	if err != nil {
		panic(err)
	}
	current := 0
	spinLock := []int{0}
	for i := 1; i <= 2017; i++ {
		insertAfter := (steps + current) % len(spinLock)
		var temp []int
		temp = append(temp, spinLock[:insertAfter]...)
		temp = append(temp, i)
		temp = append(temp, spinLock[insertAfter:]...)
		spinLock = temp
		current = insertAfter + 1
	}
	return spinLock[current]
}

func (d *Day17) part2() int {
	steps, err := strconv.Atoi(d.data[0])
	if err != nil {
		panic(err)
	}
	current := 0
	oneSlot := 0
	for i := 1; i <= 50_000_000; i++ {
		current = ((current + steps) % i) + 1
		if current == 1 {
			oneSlot = i
		}
	}
	return oneSlot
}
