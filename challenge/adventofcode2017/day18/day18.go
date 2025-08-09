package day18

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Day18 struct {
	data []string
}

func newDay18(data []string) *Day18 {
	return &Day18{data: data}
}

func (d *Day18) part1() int {
	m := machinePart1{
		registers: make(map[string]int),
	}
	return m.run(d.data)
}

func registerValue(r map[string]int, s string) int {
	if unicode.IsLetter(rune(s[0])) {
		return r[s]
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type machinePart1 struct {
	pointer   int
	lastSound int
	registers map[string]int
}

func (m *machinePart1) run(program []string) int {
	for {
		fields := strings.Fields(program[m.pointer])
		switch fields[0] {
		case "snd":
			m.lastSound = registerValue(m.registers, fields[1])
		case "set":
			m.registers[fields[1]] = registerValue(m.registers, fields[2])
		case "add":
			m.registers[fields[1]] += registerValue(m.registers, fields[2])
		case "mul":
			m.registers[fields[1]] *= registerValue(m.registers, fields[2])
		case "mod":
			m.registers[fields[1]] %= registerValue(m.registers, fields[2])
		case "rcv":
			if m.registers[fields[1]] != 0 {
				return m.lastSound
			}
		case "jgz":
			if m.registers[fields[1]] > 0 {
				m.pointer += registerValue(m.registers, fields[2]) - 1
			}
		}
		m.pointer++
		if m.pointer >= len(program) {
			return m.lastSound
		}
	}
}

type machinePart2 struct {
	programId int
	pointer   int
	sendCount int
	registers map[string]int
	rcv       []int
}

func (m *machinePart2) run(program []string, rcv <-chan int, snd chan<- int) {

	for m.pointer < len(program) {
		fields := strings.Fields(program[m.pointer])
		switch fields[0] {
		case "snd":
			snd <- registerValue(m.registers, fields[1])
			m.sendCount++
		case "set":
			m.registers[fields[1]] = registerValue(m.registers, fields[2])
		case "add":
			m.registers[fields[1]] += registerValue(m.registers, fields[2])
		case "mul":
			m.registers[fields[1]] *= registerValue(m.registers, fields[2])
		case "mod":
			m.registers[fields[1]] %= registerValue(m.registers, fields[2])
		case "rcv":
			select {
			case v := <-rcv:
				m.registers[fields[1]] = v
			case <-time.After(500 * time.Millisecond):
				fmt.Printf("%d timed out\n", m.programId)
				return
			}
		case "jgz":
			if registerValue(m.registers, fields[1]) > 0 {
				m.pointer += registerValue(m.registers, fields[2]) - 1
			}
		}
		m.pointer++
	}
}

func (d *Day18) part2() int {
	m0 := machinePart2{
		programId: 0,
		registers: map[string]int{"p": 0},
	}

	m1 := machinePart2{
		programId: 1,
		registers: map[string]int{"p": 1},
	}

	m0Send := make(chan int, 1000)
	m1Send := make(chan int, 1000)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		m0.run(d.data, m1Send, m0Send)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		m1.run(d.data, m0Send, m1Send)
	}()

	wg.Wait()
	close(m0Send)
	close(m1Send)
	return m1.sendCount
}
