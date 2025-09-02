package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	timeLimit := flag.Int("limit", 4, "Specify max time limit for each question")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error in CSV file: %s\n", *csvFilename))
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	c := 0
	for i, v := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You Scored %d/%d", c, len(problems))
			return
		}
	
	}
	fmt.Printf("You Scored %d/%d", c, len(problems))
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, v := range lines {
		ret[i] = problem{v[0], strings.TrimSpace(v[1])}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(s string) {
	fmt.Print(s)
	os.Exit(1)
}
