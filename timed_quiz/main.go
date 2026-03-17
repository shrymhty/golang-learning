package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in format of question,answer")
	quizTimer := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()


	file, err := os.Open(*csvFileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			os.Exit(1)
		}
	}

	r := csv.NewReader(file)
	questions, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv file")
		os.Exit(1)
	}	

	problems := parseLines(questions)
	timer := time.NewTimer(time.Duration(*quizTimer) * time.Second)

	correct, incorrect := 0, 0
	
	for i, p := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, p.q)
		ansChan := make(chan string)

		go func() {
			var ans string
			fmt.Scanln(&ans)
			ansChan <- ans
		}()

		select {
		case <- timer.C:
			fmt.Println("Time is up!")
			fmt.Printf("Correct: %d\nIncorrect: %d\n", correct, incorrect)
			return
		case ans := <- ansChan:
			if strings.TrimSpace(ans) == p.a {
				correct++
			} else {
				incorrect++
			}
		}
	}

	fmt.Println("Quiz Completed")
	fmt.Printf("Correct: %d, Incorrect: %d\n", correct, incorrect)
}

func parseLines(lines [][]string) []Question {
	ret := make([]Question, len(lines))

	for i, line := range lines {
		ret[i] = Question{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}