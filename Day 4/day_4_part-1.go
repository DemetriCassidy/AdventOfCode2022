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

	// parse pairs of ranges and determine if one is fully contained by the other
	// if so, increment numFullyContainedRanges
	numFullyContainedRanges := 0

	for scanner.Scan() {
		var elfOneStart, elfOneEnd, elfTwoStart, elfTwoEnd int

		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elfOneStart, &elfOneEnd, &elfTwoStart, &elfTwoEnd)
		if err != nil {
			log.Fatal(err)
		}

		if elfTwoStart >= elfOneStart && elfTwoEnd <= elfOneEnd || elfOneStart >= elfTwoStart && elfOneEnd <= elfTwoEnd {
			numFullyContainedRanges += 1
		}
	}

	// print numFullyContainedRanges
	fmt.Printf("# of fully contained ranges: %d\n", numFullyContainedRanges)
}
