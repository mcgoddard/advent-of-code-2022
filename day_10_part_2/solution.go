package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func writeToBuffer(screenBuffer *[][]string, clockCount int, registerValue int) {
	lineNumber := clockCount / 40
	yPosition := clockCount % 40
	if yPosition == registerValue-1 || yPosition == registerValue+1 || yPosition == registerValue {
		(*screenBuffer)[lineNumber][yPosition] = "#"
	} else {
		(*screenBuffer)[lineNumber][yPosition] = "."
	}
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	screenBuffer := [][]string{
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
	}
	clockCount := 0
	x := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "noop") {
			writeToBuffer(&screenBuffer, clockCount, x)
			clockCount += 1
		} else if strings.HasPrefix(line, "addx") {
			writeToBuffer(&screenBuffer, clockCount, x)
			clockCount += 1
			writeToBuffer(&screenBuffer, clockCount, x)
			clockCount += 1
			parts := strings.Split(line, " ")
			valueChange, _ := strconv.Atoi(parts[1])
			x += valueChange
		}
	}

	fmt.Println("Result:")
	for _, line := range screenBuffer {
		fmt.Println(strings.Join(line, ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
