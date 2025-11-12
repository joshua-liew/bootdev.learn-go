package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}

	runCases := []testCase{
		{2, 89, 102, 26},
		{2, 98, 104, 12},
	}

	runCases = append(runCases, []testCase{
		{3, 50, 40, -30},
		{3, 60, 60, 0},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := monthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
		_ = getBillForMonth(0, 0)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

