package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func addArray(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func maxInArray(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func main() {
	start := time.Now()
	// Read input file
	content, _ := ioutil.ReadFile("input.txt")
	content_string := string(content)
	// Split on double newline
	groups := strings.Split(content_string, "\n\n")
	// For each group
	group_calories := map2(groups, func(group string) int {
		// Split on single newline
		calory_strings := strings.Split(group, "\n")
		// Convert each line to a number
		calory_numbers := map2(calory_strings, func(calories string) int {
			number, _ := strconv.Atoi(calories)
			return number
		})
		// Sum the numbers for that group
		return addArray(calory_numbers...)
	})
	// Find the max value
	max_value := maxInArray(group_calories)
	// Print the max value
	fmt.Println(max_value)
	elapsed := time.Since(start)
	fmt.Println("Took ", elapsed)
}
