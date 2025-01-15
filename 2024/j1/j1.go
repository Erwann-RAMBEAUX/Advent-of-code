package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:embed j1_inputTest.txt
var inputTest string

//go:embed j1_inputDay.txt
var inputDay string

func split_text_to_2_slice_ordered(txt string) ([]int, []int) {
	var s1, s2 []int
	var lines []string = strings.Split(txt, "\n")
	for _, l := range lines {
		var line []string = strings.Split(l, " ")
		str_to_int, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Error, is not an int!")
			fmt.Println(line[0])
			os.Exit(0)
		}
		s1 = add_in_order(s1, str_to_int)
		str_to_int, err = strconv.Atoi(line[len(line)-1])
		if err != nil {
			fmt.Println("Error, is not an int!")
			os.Exit(0)
		}
		s2 = add_in_order(s2, str_to_int)
	}

	return s1, s2
}

func add_in_order(s []int, number int) []int {
	if len(s) == 0 {
		return append(s, number)
	}
	for i, v := range s {
		if v >= number {
			var s_sort []int
			s_sort = append(s_sort, s[:i]...)
			s_sort = append(s_sort, number)
			return append(s_sort, s[i:]...)
		}
	}
	return append(s, number)
}

func solution1(input string) int {
	var s1, s2 []int = split_text_to_2_slice_ordered(input)
	var count int
	for i := range s1 {
		if s1[i] > s2[i] {
			count += s1[i] - s2[i]
		} else {
			count += s2[i] - s1[i]
		}
	}
	return count
}

func solution2(input string) int {
	var s1, s2 []int = split_text_to_2_slice_ordered(input)
	var count int
	var m = make(map[int]int)
	for _, v := range s1 {
		var presence int
		presence, ok := m[v]
		if !ok {
			for _, v2 := range s2 {
				if v == v2 {
					presence += 1
				}
			}
			m[v] = presence
		}
		count += v * presence
	}
	return count
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputDay))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
