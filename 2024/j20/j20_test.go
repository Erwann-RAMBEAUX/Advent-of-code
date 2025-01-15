package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	result := solution(inputTest, 2, 2)
	expected := 44
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := solution(inputTest, 20, 50)
	expected := 285
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := solution(inputDay, 2, 100)
	expected := 1454
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := solution(inputDay, 20, 100)
	expected := 997879
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
