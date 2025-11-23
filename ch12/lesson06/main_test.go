package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		n        int
		expected []int
	}

	runCases := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}

	runCases = append(runCases, []testCase{
		{0, []int{}},
		{1, []int{0}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
