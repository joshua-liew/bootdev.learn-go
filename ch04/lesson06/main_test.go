package main

import (
	"fmt"
	"testing"
)

func TestGetBasicAuth(t *testing.T) {
	type testCase struct {
		auth     authenticationInfo
		expected string
	}

	runCases := []testCase{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}

	runCases = append(runCases, []testCase{
		{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"},
	}...)

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := test.auth.getBasicAuth()
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Fail
`, test.auth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Pass
`, test.auth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
