package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type OperationType int32

const (
	add      OperationType = 0
	multiply OperationType = 1
	square   OperationType = 2
)

type Operation struct {
	Operation OperationType
	Value     int
}

type Monkey struct {
	Items     []int
	Operation Operation
	Test      int
	IfTrue    int
	IfFalse   int
}

func leastCommonMultiple(monkeys []Monkey) int {
	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.Test
	}
	return lcm
}

func parseMonkeys(scanner *bufio.Scanner) []Monkey {
	monkeys := make([]Monkey, 8)

	for i, _ := range monkeys {
		// don't care about name
		scanner.Scan()
		// parse out starting items
		scanner.Scan()
		startingItems := strings.Split(strings.TrimPrefix(scanner.Text(), "  Starting items: "), ", ")
		startingItemsInts := []int{}
		for _, i := range startingItems {
			iInt, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			startingItemsInts = append(startingItemsInts, iInt)
		}
		// parse out operation
		scanner.Scan()
		operationParts := strings.Split(scanner.Text(), " ")
		operationValue, valueErr := strconv.Atoi(operationParts[len(operationParts)-1])
		var operationType OperationType
		if valueErr != nil {
			operationType = square
		} else if operationParts[len(operationParts)-2] == "+" {
			operationType = add
		} else {
			operationType = multiply
		}
		operation := Operation{
			Operation: operationType,
			Value:     operationValue,
		}
		// parse out test
		scanner.Scan()
		testParts := strings.Split(scanner.Text(), " ")
		testValue, _ := strconv.Atoi(testParts[len(testParts)-1])
		// parse out true
		scanner.Scan()
		trueParts := strings.Split(scanner.Text(), " ")
		trueValue, _ := strconv.Atoi(trueParts[len(trueParts)-1])
		// parse out false
		scanner.Scan()
		falseParts := strings.Split(scanner.Text(), " ")
		falseValue, _ := strconv.Atoi(falseParts[len(falseParts)-1])
		// build monkey
		monkeys[i] = Monkey{
			Items:     startingItemsInts,
			Operation: operation,
			Test:      testValue,
			IfTrue:    trueValue,
			IfFalse:   falseValue,
		}
		// Jump an extra line
		scanner.Scan()
	}

	return monkeys
}

func playRound(monkeys *[]Monkey, inspectCounts []int, lcm int) {
	for i, monkey := range *monkeys {
		for _, worryLevel := range monkey.Items {
			// Apply worry
			if monkey.Operation.Operation == add {
				worryLevel += monkey.Operation.Value
			} else if monkey.Operation.Operation == multiply {
				worryLevel *= monkey.Operation.Value
			} else {
				worryLevel *= worryLevel
			}
			// Apply relief by modding the least common multiple
			worryLevel %= lcm
			// Test
			if worryLevel%monkey.Test == 0 {
				(*monkeys)[monkey.IfTrue].Items = append((*monkeys)[monkey.IfTrue].Items, worryLevel)
			} else {
				(*monkeys)[monkey.IfFalse].Items = append((*monkeys)[monkey.IfFalse].Items, worryLevel)
			}
			inspectCounts[i] += 1
		}
		// Clear this monkey's items
		(*monkeys)[i].Items = []int{}
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

	monkeys := parseMonkeys(scanner)
	inspectCounts := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	lcm := leastCommonMultiple(monkeys)

	for i := 0; i < 10000; i++ {
		playRound(&monkeys, inspectCounts, lcm)
	}

	sort.Ints(inspectCounts)
	highestTwo := inspectCounts[len(inspectCounts)-2:]

	fmt.Println("Result:", highestTwo[0]*highestTwo[1])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
