package day17

import (
	"fmt"
	"math"
)

type opcode int

const (
	adv opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

var comboByOpcode = map[opcode]func(*computer, int) int{
	adv: combo,
	bxl: literal,
	bst: combo,
	jnz: literal,
	bxc: literal,
	out: combo,
	bdv: combo,
	cdv: combo,
}

var fByOpcode = map[opcode]func(*computer, int){
	adv: func(c *computer, i int) { c.a /= power(2, i) },
	bxl: func(c *computer, i int) { c.b = c.b ^ i },
	bst: func(c *computer, i int) { c.b = i % 8 },
	jnz: func(c *computer, i int) {
		if c.a == 0 {
			c.pointer++
		} else {
			c.pointer = i
		}
	},
	bxc: func(c *computer, _ int) {
		c.b = c.b ^ c.c
	},
	out: func(c *computer, i int) { c.stdOut = append(c.stdOut, i%8) },
	bdv: func(c *computer, i int) { c.b = c.a / power(2, i) },
	cdv: func(c *computer, i int) { c.c = c.a / power(2, i) },
}

type instruction struct {
	opcode  opcode
	operand int
}

type computer struct {
	a            int
	b            int
	c            int
	instructions []instruction
	pointer      int
	stdOut       []int
}

func (c *computer) run() []int {
	for c.hasNext() {
		c.nextInstruction()
	}
	return c.stdOut
}

func (c *computer) hasNext() bool {
	return c.pointer < len(c.instructions)
}

func (c *computer) nextInstruction() {
	if !c.hasNext() {
		return
	}
	instr := c.instructions[c.pointer]
	operandF := comboByOpcode[instr.opcode]
	fByOpcode[instr.opcode](c, operandF(c, instr.operand))
	if instr.opcode != jnz {
		c.pointer++
	}
}

func (c *computer) clearOutput() {
	c.stdOut = nil
}

func power(i int, j int) int {
	return int(math.Pow(float64(i), float64(j)))
}

func literal(_ *computer, op int) int {
	return op
}

func combo(c *computer, op int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic(fmt.Sprintf("%d is not a valid combo operand", op))
	}
}
