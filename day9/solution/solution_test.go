package solution

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1("input_test.txt")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := "1928"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("input_test.txt")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := "2858"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}