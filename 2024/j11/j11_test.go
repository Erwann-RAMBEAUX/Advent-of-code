package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	result := solution1(inputTest)
	expected := 55312
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := solution2(inputTest)
	expected := 65601038650482
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := solution1(inputDay)
	expected := 220999
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := solution2(inputDay)
	expected := 261936432123724
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
