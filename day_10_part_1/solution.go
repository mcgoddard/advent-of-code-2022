package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func addSignals(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i] * (20 + (i * 40))
	}
	return result
}

func storeIfInteresting(interestingSignals *[]int, clockCount int, registerValue int) {
	if (clockCount-20)%40 == 0 {
		index := (clockCount - 20) / 40
		(*interestingSignals)[index] = registerValue
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

	interestingSignals := make([]int, 6)
	clockCount := 0
	x := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "noop") {
			clockCount += 1
			storeIfInteresting(&interestingSignals, clockCount, x)
		} else if strings.HasPrefix(line, "addx") {
			clockCount += 1
			storeIfInteresting(&interestingSignals, clockCount, x)
			clockCount += 1
			storeIfInteresting(&interestingSignals, clockCount, x)
			parts := strings.Split(line, " ")
			valueChange, _ := strconv.Atoi(parts[1])
			x += valueChange
		}
	}

	fmt.Println("Result:", addSignals(interestingSignals))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
