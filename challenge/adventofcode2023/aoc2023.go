package adventofcode2023

import (
	"adventofcode2024/challenge/adventofcode2023/day01"
	"adventofcode2024/challenge/adventofcode2023/day02"
	"adventofcode2024/challenge/adventofcode2023/day03"
	"adventofcode2024/challenge/adventofcode2023/day04"
	"adventofcode2024/challenge/adventofcode2023/day07"
	"fmt"
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	day01.NewCommand(),
	day02.NewCommand(),
	day03.NewCommand(),
	day04.NewCommand(),
	day07.NewCommand(),
}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2023",
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
