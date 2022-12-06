package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	start := time.Now()
	equals := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	beats := map[string]string{
		"C": "Rock",
		"A": "Paper",
		"B": "Scissors",
	}
	looses := map[string]string{
		"A": "Scissors",
		"B": "Rock",
		"C": "Paper",
	}
	shape_scores := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}
	loose_score := 0
	draw_score := 3
	win_score := 6
	win_result := "Z"
	draw_result := "Y"

	// Read input file
	content, _ := ioutil.ReadFile("input.txt")
	content_string := string(content)
	// Split by line
	rounds := strings.Split(content_string, "\n")
	// Calculate round scores
	round_scores := map2(rounds, func(round string) int {
		decoded := strings.Split(round, " ")
		opponent := decoded[0]
		result := decoded[1]
		var outcome_score int
		var me_shape string
		if result == draw_result {
			outcome_score = draw_score
			me_shape = equals[opponent]
		} else if result == win_result {
			outcome_score = win_score
			me_shape = beats[opponent]
		} else {
			outcome_score = loose_score
			me_shape = looses[opponent]
		}
		shape_score := shape_scores[me_shape]
		return outcome_score + shape_score
	})
	// Sum scores
	total_score := addArray(round_scores...)
	// Print result
	fmt.Println(total_score)
	elapsed := time.Since(start)
	fmt.Println("Took ", elapsed)
}
