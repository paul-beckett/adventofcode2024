package main

import (
	"adventofcode2024/util/map_reduce"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type data struct {
	Year    string
	Day     string
	DayType string
	Days    []string
}

func main() {
	yearFlag := flag.String("year", "2024", "The year to generate day for")
	dayFlag := flag.String("day", "", "The days challenge")
	flag.Parse()
	year := *yearFlag
	day := strings.ToLower(*dayFlag)
	if day == "" {
		log.Fatal("day is required")
	}

	if year == "" {
		log.Fatal("year is required")
	}

	codePath := fmt.Sprintf("./challenge/adventofcode%s/%s", year, day)
	yearPath := fmt.Sprintf("./challenge/adventofcode%s/", year)

	_, err := os.Stat(codePath)
	if err == nil {
		log.Fatalf("code path already exists! %v", codePath)
	}

	log.Printf("attempting to create code files in: %s", filepath.Dir(codePath))
	if err := os.MkdirAll(codePath, 0770); err != nil {
		log.Fatalf("failed to create code dir %v", err)
	}

	entries, _ := os.ReadDir(yearPath)
	challenges := map_reduce.Filter(entries, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})
	days := map_reduce.Map(challenges, func(entry os.DirEntry) string {
		return entry.Name()
	})

	d := &data{
		Year:    year,
		Day:     day,
		DayType: strings.Title(day),
		Days:    days,
	}

	if err := createFromTemplate("day_cmd.tmpl", codePath, fmt.Sprintf("%s_cmd.go", day), d); err != nil {
		log.Fatalf("error in create cmd from template %v", err)
	}

	if err := createFromTemplate("day.tmpl", codePath, fmt.Sprintf("%s.go", day), d); err != nil {
		log.Fatalf("error in create day from template %v", err)
	}

	if err := createFromTemplate("day_test.tmpl", codePath, fmt.Sprintf("%s_test.go", day), d); err != nil {
		log.Fatalf("error in create day_test from template %v", err)
	}

	if err := createFromTemplate("year_cmd.tmpl", yearPath, fmt.Sprintf("aoc%s.go", year), d); err != nil {
		log.Fatalf("error in create year cmd from template %v", err)
	}
}

func createFromTemplate(templateName string, path string, filename string, d *data) error {
	templateFile := filepath.Join("./gen/", templateName)
	outputFile := filepath.Join(path, filename)
	log.Printf("creating %s from template %s", outputFile, templateFile)
	dayTemplate := template.Must(template.New(templateName).ParseFiles(templateFile))

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	err = dayTemplate.Execute(w, d)
	if err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	return nil
}
