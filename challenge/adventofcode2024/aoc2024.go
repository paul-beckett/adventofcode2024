package adventofcode2024

import (
	"github.com/spf13/cobra"
)

func addDays(root *cobra.Command) {
	//TODO
}

func AddYear(root *cobra.Command) {
	year := &cobra.Command{
		Use: "2024",
	}
	addDays(year)
	root.AddCommand(year)
}
