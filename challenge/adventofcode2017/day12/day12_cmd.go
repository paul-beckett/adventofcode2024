package day12

import (
	"adventofcode2024/util"
	"adventofcode2024/util/file"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day12",
		Run: func(cmd *cobra.Command, args []string) {
			day := newDay12(file.ReadFile("./input/adventofcode2017/day12.txt"))
			util.PrintResultAndTime("part1", day.part1)
			util.PrintResultAndTime("part2", day.part2)
		},
	}
}
