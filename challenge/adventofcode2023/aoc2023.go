package adventofcode2023

import (
	"adventofcode2024/challenge/adventofcode2023/day01"
	"adventofcode2024/challenge/adventofcode2023/day02"
	"github.com/spf13/cobra"
)

func addDays(root *cobra.Command) {
	day01.AddCommands(root)
	day02.AddCommands(root)
}

func AddYear(root *cobra.Command) {
	year := &cobra.Command{
		Use: "2023",
	}
	addDays(year)
	root.AddCommand(year)
}
