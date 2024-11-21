package adventofcode2024

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subCommands = []*cobra.Command{}

func NewCommand() *cobra.Command {
	year := &cobra.Command{
		Use: "2024",
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
