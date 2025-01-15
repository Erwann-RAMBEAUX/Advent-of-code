package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

//go:embed j12_inputTest.txt
var inputTest string

//go:embed j12_inputDay.txt
var inputDay string

type Pos struct {
	x, y int
}

func pos_in_slice(pos Pos, list []Pos) bool {
	for _, v := range list {
		if pos.x == v.x && pos.y == v.y {
			return true
		}
	}
	return false
}

func findArea(sos [][]uint8, area map[uint8][][]Pos) {
	for i, line := range sos {
		for j, letter := range line {
			var val, ok = area[letter]
			var position = Pos{i, j}
			if !ok {
				var slice = around(sos, letter, i, j, []Pos{position})
				area[letter] = [][]Pos{slice}
			} else {
				var find = false
				for _, slice := range val {
					if pos_in_slice(position, slice) {
						find = true
						break
					}
				}
				if !find {
					var slice = around(sos, letter, i, j, []Pos{{i, j}})
					area[letter] = append(area[letter], slice)
				}
			}
		}
	}
}

func around(sos [][]uint8, letter uint8, x, y int, slice []Pos) []Pos {
	if x-1 >= 0 && sos[x-1][y] == letter && !pos_in_slice(Pos{x - 1, y}, slice) {
		slice = append(slice, Pos{x - 1, y})
		slice = around(sos, letter, x-1, y, slice)
	}
	if x+1 < len(sos) && sos[x+1][y] == letter && !pos_in_slice(Pos{x + 1, y}, slice) {
		slice = append(slice, Pos{x + 1, y})
		slice = around(sos, letter, x+1, y, slice)
	}
	if y-1 >= 0 && sos[x][y-1] == letter && !pos_in_slice(Pos{x, y - 1}, slice) {
		slice = append(slice, Pos{x, y - 1})
		slice = around(sos, letter, x, y-1, slice)
	}
	if y+1 < len(sos[0]) && sos[x][y+1] == letter && !pos_in_slice(Pos{x, y + 1}, slice) {
		slice = append(slice, Pos{x, y + 1})
		slice = around(sos, letter, x, y+1, slice)
	}
	return slice
}

func calculatePrice(slice []Pos) int {
	var sum int
	for _, position := range slice {
		var nb_perimeter = 4
		if pos_in_slice(Pos{position.x - 1, position.y}, slice) {
			nb_perimeter--
		}
		if pos_in_slice(Pos{position.x + 1, position.y}, slice) {
			nb_perimeter--
		}
		if pos_in_slice(Pos{position.x, position.y - 1}, slice) {
			nb_perimeter--
		}
		if pos_in_slice(Pos{position.x, position.y + 1}, slice) {
			nb_perimeter--
		}
		sum += nb_perimeter
	}
	return sum * len(slice)
}

func solution1(input string) int {
	var sos = utils.Sos_uint8(input)
	var area = make(map[uint8][][]Pos)
	findArea(sos, area)
	var sum int
	for _, sos_pos := range area {
		for _, slice := range sos_pos {
			sum += calculatePrice(slice)
		}
	}
	return sum
}

func calculateSides(slice []Pos) int {
	var nb int
	var top, bottom, left, right []Pos
	for _, pos := range slice {
		if !pos_in_slice(Pos{pos.x - 1, pos.y}, slice) {
			top = append(top, pos)
		}
		if !pos_in_slice(Pos{pos.x + 1, pos.y}, slice) {
			bottom = append(bottom, pos)
		}
		if !pos_in_slice(Pos{pos.x, pos.y - 1}, slice) {
			left = append(left, pos)
		}
		if !pos_in_slice(Pos{pos.x, pos.y + 1}, slice) {
			right = append(right, pos)
		}
	}
	for _, pos := range top {
		if !pos_in_slice(Pos{pos.x, pos.y - 1}, top) {
			nb++
		}
	}
	for _, pos := range bottom {
		if !pos_in_slice(Pos{pos.x, pos.y - 1}, bottom) {
			nb++
		}
	}
	for _, pos := range left {
		if !pos_in_slice(Pos{pos.x - 1, pos.y}, left) {
			nb++
		}
	}
	for _, pos := range right {
		if !pos_in_slice(Pos{pos.x - 1, pos.y}, right) {
			nb++
		}
	}
	return nb
}

func solution2(input string) int {
	var sos = utils.Sos_uint8(input)
	var area = make(map[uint8][][]Pos)
	findArea(sos, area)
	var sum int
	for _, sos_pos := range area {
		for _, slice := range sos_pos {
			var side = calculateSides(slice)
			sum += side * len(slice)
		}
	}
	return sum
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputDay))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
