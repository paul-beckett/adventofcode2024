package day23

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Day23 struct {
	data []string
}

func newDay23(data []string) *Day23 {
	return &Day23{data: data}
}

func (d *Day23) part1() int {
	m := machinePart1{
		registers: make(map[string]int),
	}
	return m.run(d.data)
}

type machinePart1 struct {
	pointer   int
	mulCount  int
	registers map[string]int
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

func (m *machinePart1) run(program []string) int {
	for {
		//fmt.Printf("%d : %v\n", m.pointer, m.registers)
		fields := strings.Fields(program[m.pointer])
		switch fields[0] {
		case "set":
			m.registers[fields[1]] = registerValue(m.registers, fields[2])
		case "sub":
			m.registers[fields[1]] -= registerValue(m.registers, fields[2])
		case "mul":
			m.registers[fields[1]] *= registerValue(m.registers, fields[2])
			m.mulCount++
		case "jnz":
			if registerValue(m.registers, fields[1]) != 0 {
				m.pointer += registerValue(m.registers, fields[2]) - 1
			}
		default:
			panic(fmt.Errorf("unknown instr at %d: %s", m.pointer, fields[0]))
		}
		m.pointer++
		if m.pointer < 0 || m.pointer >= len(program) {
			return m.mulCount
		}
	}
	return m.mulCount
}

func (d *Day23) part2() int {
	panic("not implemented")
}
