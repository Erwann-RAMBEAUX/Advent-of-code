package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	result := solution1(inputTest)
	expected := 41
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {

	result := solution2(inputTest)
	expected := 6
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := solution1(inputDay)
	expected := 5329
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {

	result := solution2(inputDay)
	expected := 2162
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
