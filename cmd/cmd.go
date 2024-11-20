package cmd

import (
	"adventofcode2024/challenge/adventofcode2023"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:     "aoc",
		Example: "go run main.go aoc day01",
	}

	adventofcode2023.AddYear(root)

	return root
}
