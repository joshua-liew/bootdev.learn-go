package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	type testCase struct {
		x, y, expected float64
		expectedErr    string
	}

	runCases := []testCase{
		{10, 0, 0, "no dividing by 0"},
		{10, 2, 5, ""},
		{15, 30, 0.5, ""},
		{6, 3, 2, ""},
	}

	runCases = append(runCases, []testCase{
		{0, 10, 0, ""},
		{100, 0, 0, "no dividing by 0"},
		{-10, -2, 5, ""},
		{-10, 2, -5, ""},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		result, err := divide(test.x, test.y)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		if result != test.expected || errString != test.expectedErr {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.x, test.y, test.expected, test.expectedErr, result, errString)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
