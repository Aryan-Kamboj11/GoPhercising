package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	// csvFileName = flag.String("csv", "Problems.csv", "A csv file with 'question,answer' format. ")
	// flag.Parse()
	fmt.Println("Attention ! Time limit for the quiz is 30 Seconds for each question")
	file, err := os.Open("Problems.csv")
	var timeLimit int = 30
	if err != nil {
		fmt.Printf("Failed to open the file:%s\n", "Problems.csv")
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	problems := parseLine(lines)
	// fmt.Println(problems)

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	// <-timer.C

	counter := 0
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d:%s= \n ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Time limit exceeded!")
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				counter++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect! Try again from start")
				break problemLoop
			}
			break problemLoop
		}

	}
	fmt.Printf("You score %d out of %d", counter, len(problems))

}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}

	}

	return ret
}

type problem struct {
	q string
	a string
}
