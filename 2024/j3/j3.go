package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed j3_inputTest.txt
var inputTest string

//go:embed j3_inputDay.txt
var inputDay string

var regexs1 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var regexs2 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func solution1(txt string) int {
	var count int
	var sos = regexs1.FindAllStringSubmatch(txt, -1)
	for _, v := range sos {
		var i1, _ = strconv.Atoi(v[1])
		var i2, _ = strconv.Atoi(v[2])
		count += i1 * i2
	}
	return count
}

func solution2(input string) int {
	var count int
	var do bool = true
	var sos = regexs2.FindAllStringSubmatch(input, -1)
	for _, v := range sos {
		if v[0] == "do()" {
			do = true
		} else if v[0] == "don't()" {
			do = false
		}
		if do {
			var i1, _ = strconv.Atoi(v[1])
			var i2, _ = strconv.Atoi(v[2])
			count += i1 * i2
		}
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
