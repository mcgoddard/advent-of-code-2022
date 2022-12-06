package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func main() {
	equals := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	beats := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	shape_scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	loose_score := 0
	draw_score := 3
	win_score := 6

	// Read input file
	content, _ := ioutil.ReadFile("input.txt")
	content_string := string(content)
	// Split by line
	rounds := strings.Split(content_string, "\n")
	// Calculate round scores
	round_scores := map2(rounds, func(round string) int {
		shapes := strings.Split(round, " ")
		opponent := shapes[0]
		me := shapes[1]
		var outcome_score int
		if opponent == equals[me] {
			outcome_score = draw_score
		} else if opponent == beats[me] {
			outcome_score = win_score
		} else {
			outcome_score = loose_score
		}
		shape_score := shape_scores[me]
		return outcome_score + shape_score
	})
	// Sum scores
	total_score := addArray(round_scores...)
	// Print result
	fmt.Println(total_score)
}
