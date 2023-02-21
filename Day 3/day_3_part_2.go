package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getMapOfItems(rucksack string) map[rune]bool {
	itemMap := make(map[rune]bool)
	for _, item := range rucksack {
		itemMap[item] = true
	}
	return itemMap
}

func main() {
	// set up file scanner
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// find priority of common item and add to overall sum
	sumOfPriorities := 0
	for scanner.Scan() {
		// get item maps for all three elves in a group
		firstItemMap := getMapOfItems(scanner.Text())
		scanner.Scan()
		secondItemMap := getMapOfItems(scanner.Text())
		scanner.Scan()
		thirdItemMap := getMapOfItems(scanner.Text())

		// find item that appears in all three item maps
		for item := range firstItemMap {
			if secondItemMap[item] && thirdItemMap[item] {
				sumOfPriorities += int(unicode.ToLower(item) - 96)
				if unicode.IsUpper(item) {
					sumOfPriorities += 26
				}
				break
			}
		}

	}

	// print sum of priorities
	fmt.Printf("Sum of priorities: %d\n", sumOfPriorities)
}
