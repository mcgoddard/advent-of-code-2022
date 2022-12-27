package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type void struct{}

var member void

const intMax int = int(^uint(0) >> 1)

func parseGraph(scanner *bufio.Scanner) ([][]int, int, int) {
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	rows := make([][]int, len(lines))
	startIndex := 0
	endIndex := 0
	for i, line := range lines {
		rows[i] = make([]int, len(line))
		for j, c := range line {
			if c == 'S' {
				startIndex = j + (i * len(lines[0]))
				rows[i][j] = 0
			} else if c == 'E' {
				endIndex = j + (i * len(lines[0]))
				rows[i][j] = 25
			} else {
				ascii := int(c)
				rows[i][j] = ascii - 97
			}
		}
	}
	verticies := make([][]int, len(lines)*len(lines[0]))
	for i, row := range rows {
		for j, vertexValue := range row {
			vertex := make([]int, len(lines)*len(lines[0]))
			if i > 0 {
				aboveValue := rows[i-1][j]
				delta := aboveValue - vertexValue
				if delta <= 1 {
					vertex[((i-1)*len(lines[0]))+j] = 1
				}
			}
			if i < len(lines)-1 {
				belowValue := rows[i+1][j]
				delta := belowValue - vertexValue
				if delta <= 1 {
					vertex[((i+1)*len(lines[0]))+j] = 1
				}
			}
			if j > 0 {
				leftValue := rows[i][j-1]
				delta := leftValue - vertexValue
				if delta <= 1 {
					vertex[(i*len(lines[0]))+(j-1)] = 1
				}
			}
			if j < len(lines[0])-1 {
				rightValue := rows[i][j+1]
				delta := rightValue - vertexValue
				if delta <= 1 {
					vertex[(i*len(lines[0]))+(j+1)] = 1
				}
			}
			verticies[(i*len(lines[0]))+j] = vertex
		}
	}
	return verticies, startIndex, endIndex
}

func getMinimum(queue map[int]void, distance []int) int {
	minDistance := intMax
	minIndex := -1
	for index, _ := range queue {
		if minDistance > distance[index] {
			minDistance = distance[index]
			minIndex = index
		}
	}
	return minIndex
}

func dijkstra(verticies [][]int, startIndex int) ([]int, []int) {
	distance := make([]int, len(verticies))
	previous := make([]int, len(verticies))
	visited := make(map[int]void)
	queue := make(map[int]void)
	for i, _ := range verticies {
		distance[i] = intMax
		queue[i] = member
	}
	distance[startIndex] = 0

	for len(queue) > 0 {
		minIndex := getMinimum(queue, distance)
		if minIndex == -1 {
			break
		}
		neighbours := verticies[minIndex]
		for neighbourIndex, neighbourWeight := range neighbours {
			if neighbourWeight == 0 {
				continue
			}
			if _, ok := visited[neighbourIndex]; ok {
				continue
			}
			tempDistance := distance[minIndex] + neighbourWeight
			if tempDistance < distance[neighbourIndex] {
				distance[neighbourIndex] = tempDistance
				previous[neighbourIndex] = minIndex
			}
		}
		delete(queue, minIndex)
		visited[minIndex] = member
	}

	return distance, previous
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	verticies, startIndex, endIndex := parseGraph(scanner)
	distance, _ := dijkstra(verticies, startIndex)

	fmt.Println("Result:", distance[endIndex])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
