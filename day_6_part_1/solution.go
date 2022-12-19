package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type void struct{}

var member void

func areUnique(input string) bool {
	set := make(map[rune]void)
	for _, c := range input {
		set[c] = member
	}
	return len(input) == len(set)
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stream := scanner.Text()
	for i := 0; i < len(stream)-4; i++ {
		if areUnique(stream[i : i+4]) {
			fmt.Println("Result:", i+4)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
