package distance

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		list1, list2 []int
		expected     int
		expectErr    bool
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 9, false},
		{[]int{1}, []int{1}, 0, false},
		{[]int{1, 2}, []int{1}, 0, true},
		{[]int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9}, 11, false},
	}

	for _, test := range tests {
		result, err := SumOfDifference(test.list1, test.list2)
		if test.expectErr && err == nil {
			t.Errorf("Expected error, got none")
		} else if !test.expectErr && err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
