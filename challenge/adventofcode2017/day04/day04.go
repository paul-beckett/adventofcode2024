package day04

import (
	"sort"
	"strings"
)

type Day04 struct {
	data []string
}

func newDay04(data []string) *Day04 {
	return &Day04{data: data}
}

func (d *Day04) part1() int {
	noMapping := func(s string) string { return s }
	validCount := 0
	for _, line := range d.data {
		if validPassPhrase(line, noMapping) {
			validCount++
		}
	}
	return validCount
}

func validPassPhrase(phrase string, mappingF func(string) string) bool {
	words := strings.Fields(phrase)
	seen := make(map[string]bool)
	for _, word := range words {
		mapped := mappingF(word)
		if seen[mapped] {
			return false
		}
		seen[mapped] = true
	}
	return true
}

func (d *Day04) part2() int {
	sortedLetters := func(s string) string {
		letters := strings.Split(s, "")
		sort.Strings(letters)
		return strings.Join(letters, "")
	}
	validCount := 0
	for _, line := range d.data {
		if validPassPhrase(line, sortedLetters) {
			validCount++
		}
	}
	return validCount
}
