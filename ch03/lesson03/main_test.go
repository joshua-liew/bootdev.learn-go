package main

import (
	"fmt"
	"testing"
)

func Test(t *testing .T) {
	type testCase struct {
		tier		string
		expected	int
	}

	runCases := []testCase {
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
		{"invalid", 0},
		{"", 0},
	}
	testCases := runCases

	passCount := 0
	failCount := 0
	for _, test := range testCases {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed %d failed\n", passCount, failCount)
}
