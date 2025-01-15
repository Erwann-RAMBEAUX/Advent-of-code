package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed j7_inputTest.txt
var inputTest string

//go:embed j7_inputDay.txt
var inputDay string

func generation_combinaison(operators []uint8, length int) [][]uint8 {
	if length == 0 {
		return [][]uint8{{}}
	}
	var combinaisons [][]uint8
	var sub_combinaisons [][]uint8 = generation_combinaison(operators, length-1)
	for _, operator := range operators {
		for _, sub := range sub_combinaisons {
			combinaisons = append(combinaisons, append([]uint8{operator}, sub...))
		}
	}
	return combinaisons
}

func try_all_possibility(values []int, objective int, operators []uint8) bool {
	var combinaisons [][]uint8 = generation_combinaison(operators, len(values)-1)
	for _, combinaison := range combinaisons {
		var total int = values[0]
		for i, operator := range combinaison {
			switch operator {
			case '+':
				total += values[i+1]
			case '*':
				total *= values[i+1]
			case '|':
				{
					var i1 string = strconv.Itoa(total)
					var i2 string = strconv.Itoa(values[i+1])
					total, _ = strconv.Atoi(i1 + i2)
				}
			}
		}
		if total == objective {
			return true
		}
	}
	return false
}

func solution1(input string) int {
	var sos = utils.Sos_int(input)
	var result int
	for _, line := range sos {
		var objective int = line[0]
		if try_all_possibility(line[1:], objective, []uint8{'+', '*'}) {
			result += objective
		}
	}
	return result
}
func solution2(input string) int {
	var sos = utils.Sos_int(input)
	var result int
	for _, line := range sos {
		var objective int = line[0]
		if try_all_possibility(line[1:], objective, []uint8{'+', '*', '|'}) {
			result += objective
		}
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
