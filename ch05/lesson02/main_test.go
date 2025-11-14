package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		emp      employee
		expected int
	}

	runCases := []testCase{
		{fullTime{name: "Bob", salary: 7300}, 7300},
		{contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}, 856304},
	}

	runCases = append(runCases, []testCase{
		{fullTime{name: "Alice", salary: 10000}, 10000},
		{contractor{name: "John", hourlyPay: 1000, hoursPerYear: 1000}, 1000000},
	}...)

	testCases := runCases
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		salary := test.emp.getSalary()
		if salary != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Fail
`, test.emp, test.expected, salary)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Pass
`, test.emp, test.expected, salary)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
