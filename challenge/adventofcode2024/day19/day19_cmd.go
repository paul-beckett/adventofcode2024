package day19

import (
	"adventofcode2024/util"
	"adventofcode2024/util/file"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "day19",
		Run: func(cmd *cobra.Command, args []string) {
			day := newDay19(file.ReadFile("./input/adventofcode2024/day19.txt"))
			util.PrintResultAndTime("part1", day.part1)
			util.PrintResultAndTime("part2", day.part2)
		},
	}
}