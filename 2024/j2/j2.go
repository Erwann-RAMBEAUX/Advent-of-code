package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

//go:embed j2_inputTest.txt
var inputTest string

//go:embed j2_inputDay.txt
var inputDay string

func solution1(input string) int {
	var sos [][]int = utils.Sos_int(input)
	var nb_safe int
	for _, line := range sos {
		if is_safe(line) {
			nb_safe++
		}
	}
	return nb_safe
}

func is_safe(line []int) bool {
	var increase bool
	for i := 0; i < len(line)-1; i++ {
		var ecart = line[i] - line[i+1]
		if ecart > 3 || ecart < -3 || ecart == 0 {
			return false
		}
		if i == 0 {
			increase = ecart < 0
		} else {
			if (ecart > 0 && increase) || (ecart < 0 && !increase) {
				return false
			}
		}
	}
	return true
}

func solution2(input string) int {
	var sos [][]int = utils.Sos_int(input)
	var nb_safe int
	for _, line := range sos {
		if is_safe(line) {
			nb_safe++
		} else {
			for i := range line {
				var l []int
				l = append(l, line[:i]...)
				l = append(l, line[i+1:]...)
				if is_safe(l) {
					nb_safe++
					break
				}
			}
		}
	}
	return nb_safe
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputDay))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
