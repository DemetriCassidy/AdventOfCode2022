package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type stack struct {
	crates []rune
}

func (s *stack) getCrates(numToMove int) []rune {
	crates := []rune{}
	for i := 0; i < numToMove; i++ {
		crates = append([]rune{s.pop()}, crates...)
	}
	return crates
}

func (s *stack) addCrates(crates []rune) {
	s.crates = append(s.crates, crates...)
}

func (s *stack) pop() rune {
	crate := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return crate
}

func (s *stack) addToBottom(crate rune) {
	s.crates = append([]rune{crate}, s.crates...)
}

func main() {
	// set up file scanner
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// initialize stacks
	stacks := make([]stack, 9)

	scanner.Scan()
	for scanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		// crate positions = 1, 5, 9, 13, 17, 21, 25, 29, 33
		for i := 0; i < 9; i++ {
			crate := rune(scanner.Text()[(i*4)+1])
			if crate != ' ' {
				stacks[i].addToBottom(crate)
			}
		}
		scanner.Scan()
	}

	// read empty line
	scanner.Scan()

	// process instructions
	for scanner.Scan() {
		var numToMove, from, to int
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &numToMove, &from, &to)
		if err != nil {
			log.Fatal(err)
		}

		cratesToMove := stacks[from-1].getCrates(numToMove)
		stacks[to-1].addCrates(cratesToMove)
	}

	// get crates at the top of each stack
	answer := ""
	for _, stack := range stacks {
		answer = answer + string(stack.pop())
	}
	fmt.Println(answer)
}
