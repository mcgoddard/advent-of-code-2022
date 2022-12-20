package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func treeVisible(forest [][]int, y int, x int) bool {
	treeValue := forest[y][x]
	// Check above
	visibleAbove := true
	for i := y - 1; i >= 0; i-- {
		if treeValue <= forest[i][x] {
			visibleAbove = false
			break
		}
	}
	// Check below
	visibleBelow := true
	for i := y + 1; i < len(forest); i++ {
		if treeValue <= forest[i][x] {
			visibleBelow = false
			break
		}
	}
	// Check left
	visibleLeft := true
	for i := x - 1; i >= 0; i-- {
		if treeValue <= forest[y][i] {
			visibleLeft = false
			break
		}
	}
	// Check right
	visibleRight := true
	for i := x + 1; i < len(forest[y]); i++ {
		if treeValue <= forest[y][i] {
			visibleRight = false
			break
		}
	}
	return visibleAbove || visibleBelow || visibleLeft || visibleRight
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

	// We can count all the edge trees to start
	visibleTrees := (len(forest) * 2) + (len(forest[0]) * 2) - 4

	// Check each other tree
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			if treeVisible(forest, i, j) {
				visibleTrees += 1
			}
		}
	}

	fmt.Println("Result:", visibleTrees)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
