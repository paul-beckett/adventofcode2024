package day25

import (
	"adventofcode2024/util/map_reduce"
	"strconv"
	"strings"
)

type Day25 struct {
	beginState   string
	statesByName map[string]state
	stepCount    int
}

func newDay25(data []string) *Day25 {
	parts := map_reduce.ChunkBy(data, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})
	beginState := string(strings.Fields(parts[0][0])[3][0])
	stepCount, _ := strconv.Atoi(strings.Fields(parts[0][1])[5])

	statesByName := make(map[string]state)

	for _, part := range parts[1:] {
		name := string(part[0][len(part[0])-2])
		zero := createStep(part[2:5])
		one := createStep(part[6:])
		statesByName[name] = state{zeroStep: zero, oneStep: one}
	}

	return &Day25{beginState: beginState,
		statesByName: statesByName,
		stepCount:    stepCount}
}

func createStep(part []string) step {
	write, _ := strconv.Atoi(string(part[0][len(part[0])-2]))
	moveDir := strings.Fields(part[1])[6]
	move := -1
	if moveDir == "right." {
		move = 1
	}
	nextState := string(part[2][len(part[2])-2])
	return step{write: write, nextState: nextState, move: move}
}

type state struct {
	zeroStep step
	oneStep  step
}

type step struct {
	write     int
	move      int
	nextState string
}

func (d *Day25) part1() int {
	tape := make(map[int]bool)
	nextState := d.beginState
	pos := 0
	for range d.stepCount {
		current := d.statesByName[nextState]
		var currentStep step
		if tape[pos] {
			currentStep = current.oneStep
		} else {
			currentStep = current.zeroStep
		}
		if currentStep.write == 1 {
			tape[pos] = true
		} else {
			delete(tape, pos)
		}
		pos += currentStep.move
		nextState = currentStep.nextState
	}

	return len(tape)
}

func (d *Day25) part2() string {
	return "there is no part2!"
}
