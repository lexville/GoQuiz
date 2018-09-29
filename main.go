package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file that is a has a question coma answer format")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to open the file: ", *csvFileName)
	}
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Uanable to read the csv file")
	}
	problems := parseLines(lines)
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}
	fmt.Printf("Congratulations you've got %d correct \n", correct)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}
