package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed j10_inputTest.txt
var inputTest string

//go:embed j10_inputDay.txt
var inputDay string

type Void struct{}

func sos_int(txt string) [][]int {
	var lines = strings.Split(txt, "\n")
	var sos [][]int
	for _, line := range lines {
		var slice = []int{}
		for _, element := range line {
			var nb, _ = strconv.Atoi(string(element))
			slice = append(slice, nb)
		}
		sos = append(sos, slice)
	}
	return sos
}

type Pos struct {
	x, y int
}

func find_starts(sos [][]int) []Pos {
	var starts = []Pos{}
	for i, line := range sos {
		for j, element := range line {
			if element == 0 {
				starts = append(starts, Pos{i, j})
			}
		}
	}
	return starts
}

func trailheads(start Pos, sos [][]int, find map[Pos]Void) int {
	var nb int
	var actual_number = sos[start.x][start.y]
	if actual_number == 9 {
		if _, ok := find[start]; !ok {
			find[start] = Void{}
			return 1
		}
		return 0
	}
	var directions = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions {
		var newPos = Pos{start.x + direction.x, start.y + direction.y}
		if newPos.x >= 0 && newPos.x <= len(sos)-1 && newPos.y >= 0 && newPos.y <= len(sos)-1 && sos[newPos.x][newPos.y] == actual_number+1 {
			nb += trailheads(newPos, sos, find)
		}
	}
	return nb
}

func trailheads2(start Pos, sos [][]int, find map[Pos]Void) int {
	var nb int
	var actual_number = sos[start.x][start.y]
	if actual_number == 9 {
		return 1
	}
	var directions = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions {
		var newPos = Pos{start.x + direction.x, start.y + direction.y}
		if newPos.x >= 0 && newPos.x <= len(sos)-1 && newPos.y >= 0 && newPos.y <= len(sos)-1 && sos[newPos.x][newPos.y] == actual_number+1 {
			nb += trailheads2(newPos, sos, find)
		}
	}
	return nb
}

func solution1(input string) int {
	var result int
	var sos = sos_int(input)
	var starts = find_starts(sos)
	for _, start := range starts {
		var nb = trailheads(start, sos, make(map[Pos]Void))
		result += nb
	}
	return result
}

func solution2(input string) int {
	var result int
	var sos = sos_int(input)
	var starts = find_starts(sos)
	for _, start := range starts {
		var nb = trailheads2(start, sos, make(map[Pos]Void))
		result += nb
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
