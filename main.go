package main

import (
	"adventofcode2024/cmd"
	"os"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
