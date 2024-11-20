package day01

import (
	"adventofcode2024/util"
	"adventofcode2024/util/file"
	"github.com/spf13/cobra"
)

func AddCommands(root *cobra.Command) {
	day := &cobra.Command{
		Use: "day01",
		Run: func(cmd *cobra.Command, args []string) {
			day01 := newDay01(file.ReadFile("./input/adventofcode2023/day01.txt"))
			util.PrintResultAndTime("part1", day01.part1)
			util.PrintResultAndTime("part2", day01.part2)
		},
	}

	root.AddCommand(day)
}
