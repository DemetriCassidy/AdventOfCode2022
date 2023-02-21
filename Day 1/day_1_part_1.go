package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// set up file scanner
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// search through lines to find maximum sum of calories
	maxCalories, currentCalories := 0, 0

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			currentCalories += snack
		}
	}

	fmt.Printf("Maximum calories carried: %d\n", maxCalories)
}
