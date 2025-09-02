package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
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
	problems:=parseLines(lines)
	for i,v:= range problems{
		fmt.Printf("Problem #%d: %s\n",i+1,v.question)
		var ans string
		ans: =fmt.Scanf("%d")
	}
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, v := range lines {
		ret[i] = problem{v[0], v[1]}
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
