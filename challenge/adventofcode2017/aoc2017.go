package adventofcode2017

import (
	"adventofcode2024/challenge/adventofcode2017/day01"
	"adventofcode2024/challenge/adventofcode2017/day02"
	"adventofcode2024/challenge/adventofcode2017/day04"
	"adventofcode2024/challenge/adventofcode2017/day05"
	"adventofcode2024/challenge/adventofcode2017/day06"
	"adventofcode2024/challenge/adventofcode2017/day07"
	"adventofcode2024/challenge/adventofcode2017/day08"
	"adventofcode2024/challenge/adventofcode2017/day09"
	"adventofcode2024/challenge/adventofcode2017/day10"
	"adventofcode2024/challenge/adventofcode2017/day11"
	"adventofcode2024/challenge/adventofcode2017/day12"
	"adventofcode2024/challenge/adventofcode2017/day13"
	"fmt"
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	day01.NewCommand(),
	day02.NewCommand(),
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
}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2017",
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
