package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func treeScore(forest [][]int, y int, x int) int {
	treeValue := forest[y][x]
	// Check above
	upScore := 0
	for i := y - 1; i >= 0; i-- {
		upScore += 1
		if treeValue <= forest[i][x] {
			break
		}
	}
	// Check below
	downScore := 0
	for i := y + 1; i < len(forest); i++ {
		downScore += 1
		if treeValue <= forest[i][x] {
			break
		}
	}
	// Check left
	leftScore := 0
	for i := x - 1; i >= 0; i-- {
		leftScore += 1
		if treeValue <= forest[y][i] {
			break
		}
	}
	// Check right
	rightScore := 0
	for i := x + 1; i < len(forest[y]); i++ {
		rightScore += 1
		if treeValue <= forest[y][i] {
			break
		}
	}
	return upScore * downScore * leftScore * rightScore
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	forest := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		temp := make([]int, len(line))
		for i, s := range line {
			temp[i], _ = strconv.Atoi(string(s))
		}
		forest = append(forest, [][]int{temp}...)
	}

	treeScores := make([]int, len(forest)*len(forest[0]))

	// Check each tree
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			treeScores[j+(len(forest[i])*i)] = treeScore(forest, i, j)
		}
	}

	sort.Ints(treeScores)

	fmt.Println("Result:", treeScores[len(treeScores)-1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
