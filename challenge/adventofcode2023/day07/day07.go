package day07

import (
	"cmp"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day07 struct {
	data []string
}

func newDay07(data []string) *Day07 {
	return &Day07{data: data}
}

type HandType int

const (
	fiveOfAKind HandType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

type Hand struct {
	handType  HandType
	handValue int
	bid       int
}

func newHand(cards string, bid int, eval func(string) HandType, cardValues string) *Hand {
	return &Hand{
		handType:  eval(cards),
		handValue: handValuer(cards, cardValues),
		bid:       bid,
	}
}

func handTypePart1(cards string) HandType {
	cardCounts := make(map[rune]int)
	for _, card := range cards {
		cardCounts[card]++
	}
	sortedCounts := slices.SortedFunc(maps.Values(cardCounts), func(a int, b int) int {
		return cmp.Compare(b, a)
	})
	switch {
	case sortedCounts[0] == 5:
		return fiveOfAKind
	case sortedCounts[0] == 4:
		return fourOfAKind
	case sortedCounts[0] == 3 && sortedCounts[1] == 2:
		return fullHouse
	case sortedCounts[0] == 3:
		return threeOfAKind
	case sortedCounts[0] == 2 && sortedCounts[1] == 2:
		return twoPair
	case sortedCounts[0] == 2:
		return onePair
	default:
		return highCard
	}
}

func handTypePart2(cards string) HandType {
	cardCounts := make(map[rune]int)
	for _, card := range cards {
		cardCounts[card]++
	}
	jokers := cardCounts['J']
	delete(cardCounts, 'J')
	sortedCounts := slices.SortedFunc(maps.Values(cardCounts), func(a int, b int) int {
		return cmp.Compare(b, a)
	})
	switch {
	case jokers == 5:
		return fiveOfAKind
	case sortedCounts[0]+jokers == 5:
		return fiveOfAKind
	case sortedCounts[0]+jokers == 4:
		return fourOfAKind
	case sortedCounts[0]+jokers == 3 && sortedCounts[1] == 2:
		return fullHouse
	case sortedCounts[0] == 3 && sortedCounts[1]+jokers == 2:
		return fullHouse
	case sortedCounts[0]+jokers == 3:
		return threeOfAKind
	case sortedCounts[0]+jokers == 2 && sortedCounts[1] == 2:
		return twoPair
	case sortedCounts[0] == 2 && sortedCounts[1]+jokers == 2:
		return twoPair
	case sortedCounts[0]+jokers == 2:
		return onePair
	default:
		return highCard
	}
}

func handValuer(cards string, cardValues string) int {
	total := 0
	for i, card := range cards {
		v := strings.IndexRune(cardValues, card)
		total += int(math.Pow(float64(len(cardValues)), float64(len(cards)-i))) * v
	}
	return total
}

func (d *Day07) calculateWinnings(eval func(string) HandType, cardValues string) int {
	var hands []Hand
	for _, row := range d.data {
		fields := strings.Fields(row)
		bid, _ := strconv.Atoi(fields[1])
		hand := newHand(fields[0], bid, eval, cardValues)
		hands = append(hands, *hand)
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		typeSort := cmp.Compare(b.handType, a.handType)
		if typeSort != 0 {
			return typeSort
		}
		return cmp.Compare(a.handValue, b.handValue)
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}

func (d *Day07) part1() int {
	return d.calculateWinnings(handTypePart1, "23456789TJQKA")
}

func (d *Day07) part2() int {
	return d.calculateWinnings(handTypePart2, "J23456789TQKA")
}
