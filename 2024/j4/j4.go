package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

//go:embed j4_inputTest.txt
var inputTest string

//go:embed j4_inputDay.txt
var inputDay string

func find_xmas(i1, i2 int, sos [][]uint8) int {
	var count int
	var xmas []uint8 = []uint8{'M', 'A', 'S'}
	var haut, bas, gauche, droite, d1, d2, d3, d4 bool = true, true, true, true, true, true, true, true
	var lens1, lens2 int = len(sos) - 1, len(sos[0]) - 1
	for i := 1; i <= 3; i++ {
		if haut {
			if i1-i < 0 || sos[i1-i][i2] != xmas[i-1] {
				haut = false
			}
		}
		if bas {
			if i1+i > lens1 || sos[i1+i][i2] != xmas[i-1] {
				bas = false
			}
		}
		if gauche {
			if i2-i < 0 || sos[i1][i2-i] != xmas[i-1] {
				gauche = false
			}
		}
		if droite {
			if i2+i > lens2 || sos[i1][i2+i] != xmas[i-1] {
				droite = false
			}
		}
		if d1 {
			if i2+i > lens2 || i1+i > lens1 || sos[i1+i][i2+i] != xmas[i-1] {
				d1 = false
			}
		}
		if d2 {
			if i2-i < 0 || i1-i < 0 || sos[i1-i][i2-i] != xmas[i-1] {
				d2 = false
			}
		}
		if d3 {
			if i2-i < 0 || i1+i > lens1 || sos[i1+i][i2-i] != xmas[i-1] {
				d3 = false
			}
		}
		if d4 {
			if i2+i > lens2 || i1-i < 0 || sos[i1-i][i2+i] != xmas[i-1] {
				d4 = false
			}
		}

		if !haut && !bas && !gauche && !droite && !d1 && !d2 && !d3 && !d4 {
			return 0
		}
	}
	if haut {
		count++
	}
	if bas {
		count++
	}
	if gauche {
		count++
	}
	if droite {
		count++
	}
	if d1 {
		count++
	}
	if d2 {
		count++
	}
	if d3 {
		count++
	}
	if d4 {
		count++
	}
	return count
}
func solution1(input string) int {
	var sos [][]uint8 = utils.Sos_uint8(input)
	var count int
	for i1 := range sos {
		for i2, v := range sos[i1] {
			if v == 'X' {
				count += find_xmas(i1, i2, sos)
			}
		}
	}
	return count
}

func find_Xmas(i1, i2 int, sos [][]uint8) bool {
	if i1-1 >= 0 && i2-1 >= 0 && i1+1 <= len(sos)-1 && i2+1 <= len(sos[0])-1 {
		if ((sos[i1-1][i2-1] == 'M' && sos[i1+1][i2+1] == 'S') || (sos[i1-1][i2-1] == 'S' && sos[i1+1][i2+1] == 'M')) && ((sos[i1-1][i2+1] == 'M' && sos[i1+1][i2-1] == 'S') || (sos[i1-1][i2+1] == 'S' && sos[i1+1][i2-1] == 'M')) {
			return true
		}
	}
	return false
}

func solution2(input string) int {
	var sos [][]uint8 = utils.Sos_uint8(input)
	var count int
	for i1 := range sos {
		for i2, v := range sos[i1] {
			if v == 'A' {
				if find_Xmas(i1, i2, sos) {
					count++
				}
			}
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
