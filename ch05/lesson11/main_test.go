package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}

	runCases := []testCase{
		{
			email{true, "hello there", "kit@boga.com"},
			"kit@boga.com",
			0.11,
		},
		{
			sms{false, "This is Microsoft. Your computer has a virus. Please send gift cards to fix it.", "+155555509832"},
			"+155555509832",
			7.9,
		},
	}

	runCases = append(runCases, []testCase{
		{invalid{}, "", 0},
		{
			email{false, "This meeting could have been an email", "jane@doe.com"},
			"jane@doe.com",
			1.85,
		},
		{
			sms{false, "Please sir/madam", "+155555504444"},
			"+155555504444",
			1.6,
		},
	}...)

	testCases := runCases
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
