package day02

import (
	"adventofcode2024/util"
	"adventofcode2024/util/file"
	"github.com/spf13/cobra"
)

func AddCommands(root *cobra.Command) {
	day := &cobra.Command{
		Use: "day02",
		Run: func(cmd *cobra.Command, args []string) {
			day := newDay02(file.ReadFile("./input/adventofcode2023/day02.txt"))
			util.PrintResultAndTime("part1", day.part1)
			util.PrintResultAndTime("part2", day.part2)
		},
	}

	root.AddCommand(day)
}
