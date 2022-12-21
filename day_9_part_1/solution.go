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

	headPosition := []int{0, 0}
	tailPosition := []int{0, 0}
	tailPositions := map[string]void{}
	tailPositions["0:0"] = member

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			lastHeadPosition := []int{
				headPosition[0],
				headPosition[1],
			}
			// Move the head
			if move.Direction == "D" {
				headPosition[1] -= 1
			} else if move.Direction == "U" {
				headPosition[1] += 1
			} else if move.Direction == "L" {
				headPosition[0] -= 1
			} else if move.Direction == "R" {
				headPosition[0] += 1
			}
			// Determine if head has moved too far
			xDiff := headPosition[0] - tailPosition[0]
			yDiff := headPosition[1] - tailPosition[1]
			if absolute(xDiff) > 1 || absolute(yDiff) > 1 {
				// Move the tail and record new position in set
				tailPosition = lastHeadPosition
				tailPositions[fmt.Sprintf("%d:%d", tailPosition[0], tailPosition[1])] = member
			}
		}
	}

	fmt.Println("Result:", len(tailPositions))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
