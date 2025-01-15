package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed j19_inputTest.txt
var inputTest string

//go:embed j19_inputDay.txt
var inputDay string

type Void struct{}

func parse(input string) ([]string, map[string]Void) {
	var patterns = []string{}
	var list = map[string]Void{}
	var lines = strings.Split(input, "\n")
	for i, line := range lines {
		if i == 0 {
			var colors = strings.Split(line, ", ")
			for _, color := range colors {
				patterns = append(patterns, color)
			}
		} else {
			if line == "" {
				continue
			}
			list[line] = Void{}
		}
	}
	return patterns, list
}

func triPatterns(patterns []string) []string {
	var new = []string{}
	for i, pattern := range patterns {
		if i == 0 {
			new = append(new, pattern)
		} else {
			for j, newPattern := range new {
				if len(pattern) >= len(newPattern) {
					new = append(new[:j], append([]string{pattern}, new[j:]...)...)
					break
				}
			}
		}
	}
	return new
}

func isPossible(design string, patterns []string) bool {
	if len(design) == 0 {
		return true
	}
	for _, pattern := range patterns {
		if len(design) >= len(pattern) && design[:len(pattern)] == pattern {
			if isPossible(design[len(pattern):], patterns) {
				return true
			}
		}
	}
	return false
}

func solution1(input string) int {
	var patterns, list = parse(input)
	var result int
	for design := range list {
		if isPossible(design, patterns) {
			result++
		}
	}
	return result
}

func findOption(design string, patterns []string, memo map[int]int, position int) int {
	if position == len(design) {
		return 1
	}

	if val, ok := memo[position]; ok {
		return val
	}

	var result int
	for _, pattern := range patterns {
		if len(design[position:]) >= len(pattern) && design[position:position+len(pattern)] == pattern {
			result += findOption(design, patterns, memo, position+len(pattern))
		}
	}

	memo[position] = result
	return result
}

func solution2(input string) int {
	var patterns, list = parse(input)
	var result int
	for design := range list {
		var memo = map[int]int{}
		result += findOption(design, patterns, memo, 0)
	}
	return result
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputDay))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
