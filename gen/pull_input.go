package main

import (
	"adventofcode2024/util/map_reduce"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	challengeBaseDir := "./challenge"
	yearDirs, _ := os.ReadDir(challengeBaseDir)
	years := map_reduce.Filter(yearDirs, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})
	for _, year := range years {
		yearNum := strings.TrimPrefix(year.Name(), "adventofcode")

		dayDirs, _ := os.ReadDir(filepath.Join(challengeBaseDir, year.Name()))
		days := map_reduce.Filter(dayDirs, func(entry os.DirEntry) bool { return entry.IsDir() })
		for _, day := range days {
			PullInput(yearNum, day.Name())
		}
	}
}

func PullInput(year string, day string) {
	inputDir := fmt.Sprintf("./input/adventofcode%s/", year)
	inputFile := fmt.Sprintf("%s.txt", day)
	path := filepath.Join(inputDir, inputFile)
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("skipping [%s]\n", path)
		return
	}
	dayNum, _ := strconv.Atoi(strings.TrimPrefix(day, "day"))
	data := getInputData(year, dayNum)
	if err := os.MkdirAll(inputDir, 0770); err != nil {
		log.Fatalf("failed to create code dir %v", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		log.Fatalf("error writing file [%s]: %v", path, err)
	}
}

func getInputData(year string, day int) []byte {
	token, err := os.ReadFile("./input/token")
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%d/input", year, day)
	fmt.Printf("downloading data from [%s]\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: strings.TrimSpace(string(token)),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request: %v", err)
	}
	defer mustClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("unexpected status code %s\n%s", resp.Status, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}
	return body
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		log.Fatalf("error closing file: %v", err)
	}
}
