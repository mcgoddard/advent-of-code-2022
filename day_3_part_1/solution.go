package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type void struct{}

var member void

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		// Split backpack in half
		first_compartment := strings.Split(rucksack[0:len(rucksack)/2], "")
		second_compartment := strings.Split(rucksack[len(rucksack)/2:], "")
		// Turn each compartment into a set
		first_set := make(map[string]void)
		for _, s := range first_compartment {
			first_set[s] = member
		}
		second_set := make(map[string]void)
		for _, s := range second_compartment {
			second_set[s] = member
		}
		// Get set intersection
		intersection := make(map[string]void)
		if len(first_set) > len(second_set) {
			first_set, second_set = second_set, first_set // better to iterate over a shorter set
		}
		for k, _ := range first_set {
			if _, ok := second_set[k]; ok {
				intersection[k] = member
			}
		}
		// Calculate score for item in intersection
		for c, _ := range intersection {
			ascii := int([]rune(c)[0])
			// Add to running total
			if ascii > 96 {
				total += ascii - 96
			} else {
				total += ascii - 38
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total score: ", total)
}
