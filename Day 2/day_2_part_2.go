package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// set up file scanner
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// calculate total score
	// A/X == Rock, B/Y == Paper, C/Z == Scissors
	totalScore := 0
	scoreMap := map[string]int{"A X": 3, "B X": 1, "C X": 2, "A Y": 4, "B Y": 5, "C Y": 6, "A Z": 8, "B Z": 9, "C Z": 7}

	for scanner.Scan() {
		totalScore += scoreMap[scanner.Text()]
	}

	// Print final score
	fmt.Printf("Total score: %d\n", totalScore)
}
