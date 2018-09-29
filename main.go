package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", `a csv file that is 
	a has a question coma answer format`)
	timeLimit := flag.Int("limit", 30, `the time limit for the quiz
	is 30 sec`)
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.answer {
				correct++
				fmt.Println("Correct")
			} else {
				fmt.Println("Incorrect")
			}
		}

	}
	fmt.Printf("You've scored %d out of %d \n", correct, len(problems))
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
