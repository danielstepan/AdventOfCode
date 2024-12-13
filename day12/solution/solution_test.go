package solution

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct{ name, inputPath, expected string }{
		{"Test1", "test_inputs/part1/1.txt", "140"},
		{"Test2", "test_inputs/part1/2.txt", "772"},
		{"Test3", "test_inputs/part1/3.txt", "1930"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := Part1(tc.inputPath)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %q but got %q", tc.name, tc.expected, result)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct{ name, inputPath, expected string }{
		{"Test1", "test_inputs/part1/1.txt", "80"},
		{"Test2", "test_inputs/part2/1.txt", "236"},
		{"Test3", "test_inputs/part2/2.txt", "368"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := Part2(tc.inputPath)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %q but got %q", tc.name, tc.expected, result)
			}
		})
	}
}
