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

func string_to_set(input string) map[string]void {
	string_array := strings.Split(input, "")
	set := make(map[string]void)
	for _, s := range string_array {
		set[s] = member
	}
	return set
}

func intersect_sets(set_1 map[string]void, set_2 map[string]void) map[string]void {
	intersection := make(map[string]void)
	if len(set_1) > len(set_2) {
		set_1, set_2 = set_2, set_1 // better to iterate over a shorter set
	}
	for k, _ := range set_1 {
		if _, ok := set_2[k]; ok {
			intersection[k] = member
		}
	}
	return intersection
}

func read_line(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return scanner.Text(), true
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	still_reading := true
	scanner := bufio.NewScanner(file)
	rucksacks := make([]string, 3)
	sets := make([]map[string]void, 3)
	for still_reading {
		// Read three rucksacks
		for i := 0; i < 3; i++ {
			rucksacks[i], still_reading = read_line(scanner)
		}
		if !still_reading {
			break
		}
		// Convert to sets
		for i := 0; i < 3; i++ {
			sets[i] = string_to_set(rucksacks[i])
		}
		intersection := intersect_sets(sets[0], sets[1])
		intersection = intersect_sets(intersection, sets[2])
		// Add to score
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
	fmt.Println("Total score:", total)
}
