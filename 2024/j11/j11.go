package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed j11_inputTest.txt
var inputTest string

//go:embed j11_inputDay.txt
var inputDay string

func transform(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	stoneStr := strconv.Itoa(stone)
	if len(stoneStr)%2 == 0 {
		mid := len(stoneStr) / 2
		n1, _ := strconv.Atoi(stoneStr[:mid])
		n2, _ := strconv.Atoi(stoneStr[mid:])
		return []int{n1, n2}
	}
	return []int{stone * 2024}
}

func solution1(input string) int {
	var stones = utils.Sos_int(input)[0]
	for range 25 {
		for i, stone := range stones {
			var transformedStone = transform(stone)
			stones[i] = transformedStone[0]
			if len(transformedStone) == 2 {
				stones = append(stones, transformedStone[1])
			}
		}
	}
	return len(stones)
}

func solution2(input string) int {
	var stones_slice = utils.Sos_int(input)[0]
	var stones = make(map[int]int)
	for _, stone := range stones_slice {
		stones[stone]++
	}
	for range 75 {
		var newStones = make(map[int]int)

		for stone, count := range stones {
			var transformedStone = transform(stone)

			for _, newStone := range transformedStone {
				newStones[newStone] += count
			}
		}
		stones = newStones
	}
	var nb int
	for _, count := range stones {
		nb += count
	}
	return nb
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputDay))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
