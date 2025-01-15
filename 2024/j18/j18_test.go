package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	result := solution1(inputTest, 12, 6)
	expected := 22
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := solution2(inputTest, 12, 6)
	expected := Pos{6, 1}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := solution1(inputDay, 1024, 70)
	expected := 344
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := solution2(inputDay, 1024, 70)
	expected := Pos{46, 18}
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
