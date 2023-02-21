package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// search through lines to get sum of calories for each elf
	var calories []int
	currentCalories := 0

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())

		if err != nil {
			calories = append(calories, currentCalories)
			currentCalories = 0
		} else {
			currentCalories += snack
		}
	}

	// sort list of calories in descending order
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	// print sum of top 3 carried calorie amounts
	fmt.Printf("Sum of top 3 most calories carried: %d\n", calories[0]+calories[1]+calories[2])
}
