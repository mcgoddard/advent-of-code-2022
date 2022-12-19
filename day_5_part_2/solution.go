package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	count int
	from  int
	to    int
}

func parse_input(scanner *bufio.Scanner) ([]Move, [][]rune) {
	var line string
	var instructions []Move
	var stacks [][]rune
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\n") {
			continue
		} else if strings.HasPrefix(line, "[") {
			number_stacks := (len(line) + 1) / 4
			if len(stacks) == 0 {
				for i := 0; i < number_stacks; i++ {
					stack := []rune{}
					stacks = append(stacks, stack)
				}
			}
			for i := 0; i < number_stacks; i++ {
				character := line[1+(i*4)]
				if character != ' ' {
					stacks[i] = append(stacks[i], rune(character))
				}
			}
		} else if strings.HasPrefix(line, "move") {
			instruction_parts := strings.Split(line, " ")
			count, _ := strconv.Atoi(instruction_parts[1])
			from, _ := strconv.Atoi(instruction_parts[3])
			to, _ := strconv.Atoi(instruction_parts[5])
			instructions = append(instructions, Move{
				count: count,
				from:  from - 1,
				to:    to - 1,
			})
		}
	}
	return instructions, stacks
}

func make_move(instruction Move, stacks [][]rune) {
	characters_copy := make([]rune, instruction.count)
	copy(characters_copy, stacks[instruction.from][0:instruction.count])
	stacks[instruction.from] = stacks[instruction.from][instruction.count:]
	new_to := append(characters_copy, stacks[instruction.to]...)
	stacks[instruction.to] = new_to
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions, stacks := parse_input(scanner)

	for _, instruction := range instructions {
		make_move(instruction, stacks)
	}

	top_of_stacks := make([]rune, len(stacks))
	for i, stack := range stacks {
		top_of_stacks[i] = stack[0]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Final result:", string(top_of_stacks))
}
