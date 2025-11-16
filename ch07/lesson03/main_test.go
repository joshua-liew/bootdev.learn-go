package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costMultiplier   float64
		maxCostInPennies int
		expected         int
	}

	runCases := []testCase{
		{1.1, 5, 4},
		{1.3, 10, 5},
		{1.35, 25, 7},
	}

	runCases = append(runCases, []testCase{
		{1.2, 1, 1},
		{1.2, 15, 7},
		{1.3, 20, 7},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMaxMessagesToSend(test.costMultiplier, test.maxCostInPennies)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costMultiplier, test.maxCostInPennies, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
