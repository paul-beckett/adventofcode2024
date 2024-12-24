package day24

import (
	"cmp"
	"maps"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Day24 struct {
	values map[string]bool
	gates  []gate
}

type gateType int

const (
	and gateType = iota
	or
	xor
)

var gateTypes = map[string]gateType{
	"AND": and,
	"OR":  or,
	"XOR": xor,
}

type gate struct {
	lhs string
	rhs string
	gt  gateType
	out string
}

func newDay24(data []string) *Day24 {
	values := make(map[string]bool)

	i := 0
	for i < len(data) {
		if strings.TrimSpace(data[i]) == "" {
			break
		}
		parts := strings.Split(data[i], ": ")
		values[parts[0]] = parts[1] == "1"
		i++
	}
	i++

	var gates []gate
	for i < len(data) {
		fields := strings.FieldsFunc(data[i], func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		gates = append(gates, gate{
			lhs: fields[0],
			gt:  gateTypes[fields[1]],
			rhs: fields[2],
			out: fields[3],
		})
		i++
	}
	return &Day24{values: values, gates: gates}
}

func resolve(g gate, gatesByOutput map[string]gate, values map[string]bool) bool {
	a, aIsResolved := values[g.lhs]
	if !aIsResolved {
		a = resolve(gatesByOutput[g.lhs], gatesByOutput, values)
	}
	b, bIsResolved := values[g.rhs]
	if !bIsResolved {
		b = resolve(gatesByOutput[g.rhs], gatesByOutput, values)
	}
	r := false
	switch g.gt {
	case or:
		r = a || b
	case and:
		r = a && b
	case xor:
		r = (a || b) && !(a && b)
	}
	//values[g.out] = r
	return r
}

// outputs are most significant first
func sum(gatesByOutput map[string]gate, inputs map[string]bool, outputs []gate) string {
	s := ""
	for _, g := range outputs {
		if resolve(g, gatesByOutput, inputs) {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func (d *Day24) part1() int {
	gatesByOutput := make(map[string]gate)
	for _, g := range d.gates {
		gatesByOutput[g.out] = g
	}

	inputs := make(map[string]bool)
	maps.Copy(inputs, d.values)

	var outputs []gate
	for _, g := range d.gates {
		if rune(g.out[0]) == 'z' {
			outputs = append(outputs, g)
		}
	}
	slices.SortFunc(outputs, func(a, b gate) int {
		return cmp.Compare(b.out, a.out)
	})

	n, _ := strconv.ParseInt(sum(gatesByOutput, inputs, outputs), 2, 64)
	return int(n)
}

func (d *Day24) part2() int {
	panic("not implemented")
}
