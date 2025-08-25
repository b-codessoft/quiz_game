package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	listOfProblems := make([]problem, len(lines))

	for i, line := range lines {
		listOfProblems[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}

	return listOfProblems
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s\n", *csvFileName)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to parse the provided CSV file: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	correct := 0

	problems := parseLines(lines)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		scanner.Scan()
		userAnswer := strings.TrimSpace(scanner.Text())

		if userAnswer == problem.a {
			correct++
		}
	}

	fmt.Printf("\nYou got %d out of %d correct.\n", correct, len(problems))

}
