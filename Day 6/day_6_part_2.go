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

	// obtain line of input
	scanner.Scan()
	buffer := scanner.Text()

	// determine where marker is
	numCharsNeeded := 14
	for i := range buffer {
		initialChars := make(map[rune]bool)
		for j := 0; j < numCharsNeeded; j++ {
			char := rune(buffer[i+j])
			initialChars[char] = true
		}
		if len(initialChars) == numCharsNeeded {
			fmt.Println(i + numCharsNeeded)
			break
		}
	}
}
