package solution

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct{ name, inputPath, expected string }{
		{"Test1", "test_inputs/part1/1.txt", "12"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := Part1(tc.inputPath, 11, 7, 100)
			if err != nil {
				t.Fatalf("%s", err.Error())
			}

			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %q but got %q", tc.name, tc.expected, result)
			}
		})
	}
}
