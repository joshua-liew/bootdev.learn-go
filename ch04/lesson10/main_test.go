package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name           string
		membershipType string
	}

	runCases := []testCase{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
	}

	runCases = append(runCases, []testCase{
		{"Renarin", "standard"},
		{"Lift", "premium"},
		{"Dalinar", "standard"},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		user := newUser(test.name, test.membershipType)

		msgCharLimit := 100
		if test.membershipType == "premium" {
			msgCharLimit = 1000
		}

		if user.Name != test.name {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (name):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.name, user.Name)
		} else if user.Type != test.membershipType {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (membership type):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.membershipType, user.Type)
		} else if user.MessageCharLimit != msgCharLimit {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (message character limit):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, msgCharLimit, user.MessageCharLimit)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v, %v, %v
Actual:     %v, %v, %v
`, test.name, test.membershipType, test.name, test.membershipType, msgCharLimit, user.Name, user.Type, user.MessageCharLimit)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
