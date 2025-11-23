package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		numDBs int
	}

	runCases := []testCase{
		{1},
		{3},
		{4},
	}

	runCases = append(runCases, []testCase{
		{0},
		{13},
	}...)

	testCases := runCases

	passed, failed := 0, 0

	for _, test := range testCases {
		fmt.Printf(`---------------------------------`)
		fmt.Printf("\nTesting %v Databases...\n\n", test.numDBs)
		dbChan, count := getDBsChannel(test.numDBs)
		waitForDBs(test.numDBs, dbChan)
		for *count != test.numDBs {
			fmt.Println("...")
		}
		if len(dbChan) == 0 && *count == test.numDBs {
			passed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
PASS
`,
				test.numDBs, len(dbChan), *count)
		} else {
			failed++
			fmt.Printf(`
expected length: 0, count: %v
actual length:   %v, count: %v
FAIL
`,
				test.numDBs, len(dbChan), *count)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passed, failed)
}
