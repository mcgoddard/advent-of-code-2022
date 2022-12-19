package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_line(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return scanner.Text(), true
}

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	min_max_for_pair := [2][2]int{
		{0, 0},
		{0, 0},
	}

	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		for i := 0; i < 2; i++ {
			values := strings.Split(pair[i], "-")
			for j, v := range values {
				min_max_for_pair[i][j], _ = strconv.Atoi(v)
			}
		}
		if (min_max_for_pair[0][0] <= min_max_for_pair[1][0] && min_max_for_pair[0][1] >= min_max_for_pair[1][0]) ||
			(min_max_for_pair[1][0] <= min_max_for_pair[0][0] && min_max_for_pair[1][1] >= min_max_for_pair[0][0]) {
			total += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total score:", total)
}
