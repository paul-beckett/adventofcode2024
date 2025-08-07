package day18

import (
	"strconv"
	"strings"
	"unicode"
)

type Day18 struct {
	data []string
}

func newDay18(data []string) *Day18 {
	return &Day18{data: data}
}

func (d *Day18) part1() int {
	registers := make(map[string]int)
	value := func(s string) int {
		if unicode.IsLetter(rune(s[0])) {
			return registers[s]
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}
	lastSound := -1
	pointer := 0
	for {
		if pointer >= len(d.data) {
			break
		}
		fields := strings.Fields(d.data[pointer])

		instr := fields[0]
		arg1 := fields[1]
		arg2 := ""
		if len(fields) >= 3 {
			arg2 = fields[2]
		}

		switch instr {
		case "snd":
			lastSound = registers[arg1]
		case "set":
			registers[arg1] = value(arg2)
		case "add":
			registers[arg1] += value(arg2)
		case "mul":
			registers[arg1] *= value(arg2)
		case "mod":
			registers[arg1] %= value(arg2)
		case "rcv":
			if registers[arg1] != 0 {
				return lastSound
			}
		case "jgz":
			if registers[arg1] != 0 {
				pointer += value(arg2)
				continue
			}
		default:
			panic("invalid instruction: " + fields[0])
		}
		pointer++
	}
	return -1
}

func (d *Day18) part2() int {
	panic("not implemented")
}
