package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	// set up file scanner
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// find priority of item that appears in both compartments and add to overall sum
	sumOfPriorities := 0
	for scanner.Scan() {
		itemMap := make(map[rune]bool)
		firstCompartment := scanner.Text()[:len(scanner.Text())/2]
		secondCompartment := scanner.Text()[len(scanner.Text())/2:]

		// loop through first compartment and add each character to the map
		for _, item := range firstCompartment {
			itemMap[item] = true
		}

		// loop through second compartment until the common item is found
		for _, item := range secondCompartment {
			if itemMap[item] {
				sumOfPriorities += int(unicode.ToLower(item) - 96)
				if unicode.IsUpper(item) {
					sumOfPriorities += 26
				}
				break
			}
		}
	}

	// Print sum of priorities
	fmt.Printf("Sum of priorities: %d\n", sumOfPriorities)
}
