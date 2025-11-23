package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numBatches int
		expected   int
	}

	runCases := []testCase{
		{3, 114},
		{4, 198},
	}

	runCases = append(runCases, []testCase{
		{0, 0},
		{1, 15},
		{6, 435},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		numSentCh := make(chan int)
		go sendReports(test.numBatches, numSentCh)
		output := countReports(numSentCh)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
