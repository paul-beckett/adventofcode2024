package day05

import (
	"adventofcode2024/util/ints"
	"adventofcode2024/util/map_reduce"
	"slices"
	"strings"
)

type Day05 struct {
	data []string
}

func newDay05(data []string) *Day05 {
	return &Day05{data: data}
}

type rule []int

func parseRule(line string) rule {
	return ints.ToInts(line, func(r rune) bool {
		return r == '|'
	})
}

type pageUpdate []int

func parsePageUpdate(line string) pageUpdate {
	return ints.ToInts(line, func(r rune) bool { return r == ',' })
}

type safetyManual struct {
	rules       []rule
	pageUpdates []pageUpdate
}

func parseSafetyManual(data []string) *safetyManual {
	manual := map_reduce.ChunkBy(data, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})
	var rules []rule
	for _, s := range manual[0] {
		rules = append(rules, parseRule(s))
	}

	var pageUpdates []pageUpdate
	for _, s := range manual[1] {
		pageUpdates = append(pageUpdates, parsePageUpdate(s))
	}
	return &safetyManual{
		rules:       rules,
		pageUpdates: pageUpdates,
	}
}

func (d *Day05) part1() int {
	safetyManual := parseSafetyManual(d.data)
	validUpdates := map_reduce.Filter(safetyManual.pageUpdates, func(p pageUpdate) bool {
		return validPageUpdate(p, safetyManual.rules)
	})
	middlePages := map_reduce.Map(validUpdates, func(p pageUpdate) int { return p[(len(p))/2] })
	return map_reduce.Sum(middlePages)
}

func validPageUpdate(p pageUpdate, rules []rule) bool {
	positions := make(map[int]int)
	for i, page := range p {
		positions[page] = i
	}

	return map_reduce.All(rules, func(r rule) bool {
		lhs, ok := positions[r[0]]
		if !ok {
			return true
		}
		rhs, ok := positions[r[1]]
		if !ok {
			return true
		}
		return lhs < rhs
	})
}

func (d *Day05) part2() int {
	safetyManual := parseSafetyManual(d.data)
	invalidUpdates := map_reduce.Filter(safetyManual.pageUpdates, func(p pageUpdate) bool {
		return !validPageUpdate(p, safetyManual.rules)
	})

	ruleMatcher := make(map[int]map[int]rule)
	for _, r := range safetyManual.rules {
		lhs, ok := ruleMatcher[r[0]]
		if !ok {
			lhs = make(map[int]rule)
			ruleMatcher[r[0]] = lhs
		}
		lhs[r[1]] = r
	}
	match := func(a, b int) rule {
		if ruleMatcher[a] != nil && ruleMatcher[a][b] != nil {
			return ruleMatcher[a][b]
		} else if ruleMatcher[b] != nil && ruleMatcher[b][a] != nil {
			return ruleMatcher[b][a]
		} else {
			return nil
		}
	}

	sortByRules := func(a, b int) int {
		rule := match(a, b)
		if rule == nil {
			return 0
		} else if rule[0] == a {
			return -1
		} else {
			return 1
		}
	}

	invalidUpdates = map_reduce.Map(invalidUpdates, func(p pageUpdate) pageUpdate {
		slices.SortFunc(p, sortByRules)
		return p
	})
	middlePages := map_reduce.Map(invalidUpdates, func(p pageUpdate) int { return p[(len(p))/2] })
	return map_reduce.Sum(middlePages)
}
