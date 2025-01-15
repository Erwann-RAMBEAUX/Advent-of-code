package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed j5_inputTest.txt
var inputTest string

//go:embed j5_inputDay.txt
var inputDay string

func treatment_input(txt string) ([][]int, [][]int) {
	var rules [][]int
	var orders [][]int
	var lines []string = strings.Split(txt, "\n")
	var second_part bool
	for _, line := range lines {
		if line == "" {
			second_part = true
		} else {
			if second_part {
				var order []int
				var elements = strings.Split(line, ",")
				for _, element := range elements {
					var number, _ = strconv.Atoi(string(element))
					order = append(order, number)
				}
				orders = append(orders, order)
			} else {
				var rule []int
				var elements = strings.Split(line, "|")
				for _, element := range elements {
					var number, _ = strconv.Atoi(string(element))
					rule = append(rule, number)
				}
				rules = append(rules, rule)
			}
		}
	}
	return rules, orders
}

func rules_verify(element int, rules [][]int, next []int) bool {
	var count_nb_true int
	for _, rule := range rules {
		if element == rule[0] && utils.Int_in_slice(rule[1], next) {
			count_nb_true++
		}
	}
	return count_nb_true == len(next)
}

func solution1(input string) int {
	var rules, orders = treatment_input(input)
	var result int
	for _, order := range orders {
		var order_respect bool = true
		for i, element := range order {
			if !rules_verify(element, rules, order[i+1:]) {
				order_respect = false
				break
			}
		}
		if order_respect {
			var middle int = len(order) / 2
			result += order[middle]
		}

	}
	return result
}

func ordering(order []int, rules [][]int) []int {
	var order_ordered []int
	for _, element := range order {
		if len(order_ordered) == 0 {
			order_ordered = append(order_ordered, element)
		} else {
			var position int
			for _, v := range order_ordered {
				for _, rule := range rules {
					if rule[0] == element && rule[1] == v {
						break
					} else if rule[0] == v && rule[1] == element {
						position++
					}
				}
			}
			var ordered []int = make([]int, 0)
			var p1 []int = order_ordered[:position]
			ordered = append(ordered, p1...)
			ordered = append(ordered, element)
			if position+1 <= len(order_ordered) {
				var p2 []int = order_ordered[position:]
				ordered = append(ordered, p2...)
			}
			order_ordered = ordered
		}
	}
	return order_ordered
}

func solution2(input string) int {
	var rules, orders = treatment_input(input)
	var result int
	for _, order := range orders {
		var order_respect bool = true
		for i, element := range order {
			if !rules_verify(element, rules, order[i+1:]) {
				order_respect = false
				break
			}
		}
		if !order_respect {
			order = ordering(order, rules)
			var middle int = len(order) / 2
			result += order[middle]
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
