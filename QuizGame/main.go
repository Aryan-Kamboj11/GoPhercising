package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// csvFileName = flag.String("csv", "Problems.csv", "A csv file with 'question,answer' format. ")
	// flag.Parse()

	file, err := os.Open("Problems.csv")

	if err != nil {
		fmt.Printf("Failed to open the file:%s\n", file)
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	problems := parseLine(lines)
	fmt.Println(problems)

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
