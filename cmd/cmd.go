package cmd

import (
	"adventofcode2024/challenge/adventofcode2023"
	"adventofcode2024/challenge/adventofcode2024"
	"fmt"
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{
	adventofcode2023.NewCommand(),
	adventofcode2024.NewCommand(),
}

func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:     "aoc",
		Example: "go run main.go aoc2024 day01",
		Run: func(cmd *cobra.Command, args []string) {
			for _, subCommand := range subCommands {
				fmt.Printf("running %s:\n", subCommand.Name())
				subCommand.Run(cmd, args)
			}
		},
	}
	for _, subCommand := range subCommands {
		root.AddCommand(subCommand)
	}

	return root
}
