package day09

type Day09 struct {
	data []string
}

func newDay09(data []string) *Day09 {
	return &Day09{data: data}
}

type rule struct {
	closedBy func(rune) bool
	canOpen  func(rune) bool
}

var rulesByOpen = map[rune]rule{
	'{': {
		closedBy: func(r rune) bool { return r == '}' },
		canOpen:  func(r rune) bool { return true },
	},
	'<': {
		closedBy: func(r rune) bool { return r == '>' },
		canOpen:  func(r rune) bool { return r == '!' },
	},
	'!': {
		closedBy: func(r rune) bool { return true },
		canOpen:  func(r rune) bool { return false },
	},
}

func (d *Day09) part1() int {
	score := 0
	var stack []rune

	for _, c := range d.data[0] {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		top := stack[len(stack)-1]
		_, isOpen := rulesByOpen[c]
		if isOpen && rulesByOpen[top].canOpen(c) {
			stack = append(stack, c)
		} else if rulesByOpen[top].closedBy(c) {
			if top == '{' {
				score += len(stack)
			}
			stack = stack[:len(stack)-1]
		}
	}

	return score
}

func (d *Day09) part2() int {
	score := 0
	var stack []rune

	for _, c := range d.data[0] {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		top := stack[len(stack)-1]
		_, isOpen := rulesByOpen[c]
		if isOpen && rulesByOpen[top].canOpen(c) {
			stack = append(stack, c)
		} else if rulesByOpen[top].closedBy(c) {
			stack = stack[:len(stack)-1]
		} else if top == '<' {
			score++
		}
	}

	return score
}
