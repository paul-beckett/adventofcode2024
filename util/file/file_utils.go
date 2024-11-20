package file

import (
	"bufio"
	"os"
)

func ReadFile(filename string) []string {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	content := make([]string, 0)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return content
}
