package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

//go:embed j20_inputTest.txt
var inputTest string

//go:embed j20_inputDay.txt
var inputDay string

type Pos struct {
	x, y int
}

func findLetter(letter uint8, grid [][]uint8) Pos {
	for x, line := range grid {
		for y, square := range line {
			if square == letter {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}

func inGrid(i, j, size int) bool {
	return i >= 0 && i < size && j >= 0 && j < size
}

func originalPath(start, end Pos, grid [][]uint8) []Pos {
	var directions = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var size = len(grid)

	path := []Pos{start}
	for path[len(path)-1] != end {
		x, y := path[len(path)-1].x, path[len(path)-1].y
		for _, dir := range directions {
			xi, yj := x+dir.x, y+dir.y
			if !inGrid(xi, yj, size) || grid[xi][yj] == '#' {
				continue
			}
			if len(path) > 1 && (xi == path[len(path)-2].x && yj == path[len(path)-2].y) {
				continue
			}
			path = append(path, Pos{xi, yj})
			break
		}
	}
	return path
}

func manhattanDistance(point, goal Pos) int {
	return abs(point.x-goal.x) + abs(point.y-goal.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solution(input string, picoCheat int, picoSave int) int {
	var result int
	var grid = utils.Sos_uint8(input)
	var start = findLetter('S', grid)
	var end = findLetter('E', grid)
	var path = originalPath(start, end, grid)
	var picoSeconds = len(path)
	var squares = make(map[Pos]int)

	for i, pos := range path {
		squares[pos] = picoSeconds - i
	}

	var savedPico = make(map[[4]int]int)
	var size = len(grid)
	for i, pos := range path {
		for x := pos.x - picoCheat; x <= pos.x+picoCheat; x++ {
			for y := pos.y - picoCheat; y <= pos.y+picoCheat; y++ {
				var distance = manhattanDistance(Pos{x, y}, pos)
				if !inGrid(x, y, size) || distance > picoCheat || grid[x][y] == '#' {
					continue
				}
				var squareDistance = squares[Pos{x, y}]
				savedPico[[4]int{pos.x, pos.y, x, y}] = picoSeconds - (i + squareDistance + distance)
			}
		}
	}

	for _, v := range savedPico {
		if v >= picoSave {
			result++
		}
	}
	return result
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution(inputTest, 2, 2))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution(inputDay, 20, 100))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
