package day_template

import (
	"adventofcode2024/util"
	"adventofcode2024/util/file"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "dayTemplate",
		Run: func(cmd *cobra.Command, args []string) {
			day := newDayTemplate(file.ReadFile("./input/day_template.txt"))
			util.PrintResultAndTime("part1", day.part1)
			util.PrintResultAndTime("part2", day.part2)
		},
	}
}
