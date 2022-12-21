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
	Direction string
	Steps     int
}

type void struct{}

var member void

func absolute(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	moves := make([]Move, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[1])
		moves = append(moves, Move{
			Direction: parts[0],
			Steps:     steps,
		})
	}

	positions := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	tailPositions := map[string]void{}
	tailPositions["0:0"] = member

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			// Move the head
			if move.Direction == "D" {
				positions[0][1] -= 1
			} else if move.Direction == "U" {
				positions[0][1] += 1
			} else if move.Direction == "L" {
				positions[0][0] -= 1
			} else if move.Direction == "R" {
				positions[0][0] += 1
			}
			// Follow through for additional knots
			for s := 1; s < 10; s++ {
				diffX := positions[s-1][0] - positions[s][0]
				diffY := positions[s-1][1] - positions[s][1]
				if diffY == 0 {
					// No y differences
					if diffX > 1 {
						positions[s][0]++
					} else if diffX < -1 {
						positions[s][0]--
					}
				} else if diffX == 0 {
					// No x differences
					if diffY > 1 {
						positions[s][1]++
					} else if diffY < -1 {
						positions[s][1]--
					}
				} else if absolute(diffX) == 2 || absolute(diffY) == 2 {
					// y or x differences
					if diffX == 2 && diffY == 2 {
						positions[s][0]++
						positions[s][1]++
					} else if diffX == -2 && diffY == -2 {
						positions[s][0]--
						positions[s][1]--
					} else if diffX == 2 && diffY == -2 {
						positions[s][0]++
						positions[s][1]--
					} else if diffX == -2 && diffY == 2 {
						positions[s][0]--
						positions[s][1]++
					} else if (diffX == 1 && diffY == 2) || (diffX == 2 && diffY == 1) {
						positions[s][0]++
						positions[s][1]++
					} else if (diffX == -2 && diffY == -1) || (diffX == -1 && diffY == -2) {
						positions[s][0]--
						positions[s][1]--
					} else if (diffX == 1 && diffY == -2) || (diffX == 2 && diffY == -1) {
						positions[s][0]++
						positions[s][1]--
					} else if (diffX == -2 && diffY == 1) || (diffX == -1 && diffY == 2) {
						positions[s][0]--
						positions[s][1]++
					}
				}
			}
			// Add final knot position to set
			tailPositions[fmt.Sprintf("%d:%d", positions[9][0], positions[9][1])] = member
		}
	}

	fmt.Println("Result:", len(tailPositions))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
