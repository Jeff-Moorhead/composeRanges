package composeRanges

import "testing"

type testCase struct {
	Input    []int
	Expected []string
}

func TestSolution(t *testing.T) {

	tests := []testCase{
		{
			Input:    []int{},
			Expected: []string{},
		},
		{
			Input:    []int{5},
			Expected: []string{"5"},
		},
		{
			Input:    []int{5, 6, 7, 8},
			Expected: []string{"5->8"},
		},
		{
			Input:    []int{-1, 0, 1, 2, 6, 8, 9, 10},
			Expected: []string{"-1->2", "6", "8->10"},
		},
		{
			Input:    []int{-1, 0, 1, 2, 6, 7, 9},
			Expected: []string{"-1->2", "6->7", "9"},
		},
	}

	for _, test := range tests {
		got := Solution(test.Input)
		for j, result := range got {
			if result != test.Expected[j] {
				t.Logf("Test: %#v", test)
				t.Logf("Got: %v", got)
				t.Logf("Expected: %#v", test.Expected)
				t.Fail()
			}
		}
	}
}
