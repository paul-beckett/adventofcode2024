package adventofcode2024

import (
	"adventofcode2024/challenge/adventofcode2024/day01"
	"adventofcode2024/challenge/adventofcode2024/day02"
	"adventofcode2024/challenge/adventofcode2024/day03"
	"adventofcode2024/challenge/adventofcode2024/day04"
	"adventofcode2024/challenge/adventofcode2024/day05"
	"adventofcode2024/challenge/adventofcode2024/day06"
	"adventofcode2024/challenge/adventofcode2024/day07"
	"adventofcode2024/challenge/adventofcode2024/day08"
	"adventofcode2024/challenge/adventofcode2024/day09"
	"adventofcode2024/challenge/adventofcode2024/day10"
	"adventofcode2024/challenge/adventofcode2024/day11"
	"adventofcode2024/challenge/adventofcode2024/day12"
	"adventofcode2024/challenge/adventofcode2024/day13"
	"adventofcode2024/challenge/adventofcode2024/day14"
	"adventofcode2024/challenge/adventofcode2024/day15"
	"adventofcode2024/challenge/adventofcode2024/day16"
	"adventofcode2024/challenge/adventofcode2024/day17"
	"adventofcode2024/challenge/adventofcode2024/day18"
	"fmt"
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	day01.NewCommand(),
	day02.NewCommand(),
	day03.NewCommand(),
	day04.NewCommand(),
	day05.NewCommand(),
	day06.NewCommand(),
	day07.NewCommand(),
	day08.NewCommand(),
	day09.NewCommand(),
	day10.NewCommand(),
	day11.NewCommand(),
	day12.NewCommand(),
	day13.NewCommand(),
	day14.NewCommand(),
	day15.NewCommand(),
	day16.NewCommand(),
	day17.NewCommand(),
	day18.NewCommand(),
}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2024",
		Run: func(cmd *cobra.Command, args []string) {
			for _, subCommand := range subCommands {
				fmt.Printf("running %s:\n", subCommand.Name())
				subCommand.Run(cmd, args)
			}
		},
	}
	for _, cmd := range subCommands {
		year.AddCommand(cmd)
	}
	return year
}
