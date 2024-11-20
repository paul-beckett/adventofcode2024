package day02

import (
	"adventofcode2024/util/map_reduce"
	"strconv"
	"strings"
)

type Day02 struct {
	data []string
}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	blue  int
	green int
}

func newDay02(data []string) *Day02 {
	return &Day02{data: data}
}

func newGame(row string) *Game {
	rowParts := strings.Split(row, ": ")
	gameId, _ := strconv.Atoi(strings.TrimPrefix(rowParts[0], "Game "))
	roundsAsString := strings.Split(rowParts[1], "; ")
	var rounds []Round
	for _, round := range roundsAsString {
		rounds = append(rounds, *newRound(round))
	}
	return &Game{
		id:     gameId,
		rounds: rounds,
	}
}

func newRound(round string) *Round {
	cubesAsString := strings.Split(round, ", ")
	cubesByColour := make(map[string]int)
	for _, cube := range cubesAsString {
		numberAndColour := strings.Split(cube, " ")
		cubeCount, _ := strconv.Atoi(numberAndColour[0])
		cubesByColour[numberAndColour[1]] = cubeCount
	}
	return &Round{
		red:   cubesByColour["red"],
		blue:  cubesByColour["blue"],
		green: cubesByColour["green"],
	}
}

func (d *Day02) parseGames() []Game {
	var games []Game
	for _, row := range d.data {
		games = append(games, *newGame(row))
	}
	return games
}

func (d *Day02) part1() int {
	roundPossible := func(r Round) bool { return r.red <= 12 && r.green <= 13 && r.blue <= 14 }
	gameFilter := func(g Game) bool { return map_reduce.All(g.rounds, roundPossible) }
	possibleGames := map_reduce.Filter(d.parseGames(), gameFilter)
	idSum := func(acc int, g Game) int { return acc + g.id }

	return map_reduce.Reduce(possibleGames, idSum, 0)
}

func (d *Day02) part2() int {
	minPossible := func(acc Round, r Round) Round {
		acc.red = max(acc.red, r.red)
		acc.green = max(acc.green, r.green)
		acc.blue = max(acc.blue, r.blue)
		return acc
	}
	power := func(r Round) int { return r.red * r.green * r.blue }

	powers := map_reduce.Map(d.parseGames(), func(g Game) int {
		requiredCubes := map_reduce.Reduce(g.rounds, minPossible, Round{})
		return power(requiredCubes)
	})
	return map_reduce.Sum(powers)
}
