package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func String_in_slice(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}

func Uint8_in_slice(s uint8, list []uint8) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}

func Int_in_slice(n int, list []int) bool {
	for _, v := range list {
		if n == v {
			return true
		}
	}
	return false
}

// Create a sos of string from input
func Sos_string(input string) [][]string {
	// Function to create the grid from the text
	var lines = strings.Split(input, "\n")
	var sos [][]string
	for _, line := range lines {
		var x = strings.Split(line, "")
		sos = append(sos, x)
	}
	return sos
}

// Create a sos of int from the input
func Sos_int(input string) [][]int {
	var lines = strings.Split(input, "\n")
	var sos [][]int
	for _, line := range lines {
		var x = strings.Split(line, " ")
		var slice []int
		for _, v := range x {
			if v[len(v)-1] == ':' {
				v = v[:len(v)-1]
			}
			str_to_int, err := strconv.Atoi(v)
			if err == nil {
				slice = append(slice, str_to_int)
			} else {
				fmt.Println("This element, is not an int! " + v)
				os.Exit(0)
			}
		}
		sos = append(sos, slice)
	}
	return sos
}

func Sos_uint8(input string) [][]uint8 {
	var sos [][]uint8
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var slice = []uint8{}
		for _, caractere := range line {
			slice = append(slice, uint8(caractere))
		}
		sos = append(sos, slice)
	}
	return sos
}
