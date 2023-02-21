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
	scoreMap := map[string]int{"A X": 4, "B X": 1, "C X": 7, "A Y": 8, "B Y": 5, "C Y": 2, "A Z": 3, "B Z": 9, "C Z": 6}

	for scanner.Scan() {
		totalScore += scoreMap[scanner.Text()]
	}

	// Print final score
	fmt.Printf("Total score: %d\n", totalScore)
}
