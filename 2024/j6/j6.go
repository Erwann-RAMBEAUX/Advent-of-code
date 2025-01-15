package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

//go:embed j6_inputTest.txt
var inputTest string

//go:embed j6_inputDay.txt
var inputDay string

type Pos struct {
	x, y int
}

type Guard struct {
	Pos
	direction string
}

func guard_position(sos [][]uint8) Guard {
	for i, l := range sos {
		for j, c := range l {
			if c == '^' {
				return Guard{Pos{i, j}, "haut"}
			} else if c == 'v' {
				return Guard{Pos{i, j}, "bas"}
			} else if c == '>' {
				return Guard{Pos{i, j}, "droite"}
			} else if c == '<' {
				return Guard{Pos{i, j}, "gauche"}
			}
		}
	}
	return Guard{}
}

func movement(sos [][]uint8, guard Guard) map[Pos]struct{} {
	var all_pos = make(map[Pos]struct{})
	for {
		all_pos[guard.Pos] = struct{}{}
		if guard.direction == "haut" {
			if guard.x-1 < 0 {
				break
			} else if sos[guard.x-1][guard.y] == '#' {
				guard.direction = "droite"
				guard.y += 1
			} else {
				guard.x -= 1
			}
		} else if guard.direction == "bas" {
			if guard.x+1 > len(sos)-1 {
				break
			} else if sos[guard.x+1][guard.y] == '#' {
				guard.direction = "gauche"
				guard.y -= 1
			} else {
				guard.x += 1
			}
		} else if guard.direction == "droite" {
			if guard.y+1 > len(sos[0])-1 {
				break
			} else if sos[guard.x][guard.y+1] == '#' {
				guard.direction = "bas"
				guard.x += 1
			} else {
				guard.y += 1
			}
		} else if guard.direction == "gauche" {
			if guard.y-1 < 0 {
				break
			} else if sos[guard.x][guard.y-1] == '#' {
				guard.direction = "haut"
				guard.x -= 1
			} else {
				guard.y -= 1
			}
		}
	}
	return all_pos
}

func solution1(input string) int {
	var sos = utils.Sos_uint8(input)
	var guard = guard_position(sos)
	var all_pos = movement(sos, guard)
	return len(all_pos)
}

func is_loop(sos [][]uint8, guard Guard) bool {
	visited := make(map[Pos]int)
	for {
		if visited[guard.Pos] > 5 {
			return true
		}
		visited[guard.Pos] += 1

		if guard.direction == "haut" {
			if guard.x-1 < 0 {
				break
			} else if sos[guard.x-1][guard.y] == '#' {
				guard.direction = "droite"
			} else {
				guard.x--
			}
		} else if guard.direction == "bas" {
			if guard.x+1 >= len(sos) {
				break
			} else if sos[guard.x+1][guard.y] == '#' {
				guard.direction = "gauche"
			} else {
				guard.x++
			}
		} else if guard.direction == "gauche" {
			if guard.y-1 < 0 {
				break
			} else if sos[guard.x][guard.y-1] == '#' {
				guard.direction = "haut"
			} else {
				guard.y--
			}
		} else if guard.direction == "droite" {
			if guard.y+1 >= len(sos[0]) {
				break
			} else if sos[guard.x][guard.y+1] == '#' {
				guard.direction = "bas"
			} else {
				guard.y++
			}
		}
	}
	return false
}

func solution2(input string) int {
	var sos = utils.Sos_uint8(input)
	var count int
	var all_pos = movement(sos, guard_position(sos))
	for position := range all_pos {
		var guard = guard_position(sos)
		if sos[position.x][position.y] == '#' || (position.x == guard.x && position.y == guard.y) {
			continue
		}
		sos[position.x][position.y] = '#'
		var ok = is_loop(sos, guard)
		if ok {
			count++
		}
		sos[position.x][position.y] = '.'
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
