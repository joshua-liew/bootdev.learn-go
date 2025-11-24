package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		input    interface{}
		expected interface{}
	}
	runCases := []testCase{
		{[]int{}, 0},
		{[]bool{true, false, true, true, false}, false},
	}

	runCases = append(runCases, []testCase{
		{[]int{1, 2, 3, 4}, 4},
		{[]string{"a", "b", "c", "d"}, "d"},
	}...)

	testCases := runCases

	passed, failed := 0, 0

	for _, test := range testCases {
		switch v := test.input.(type) {
		case []int:
			if output := getLast(v); output != test.expected {
				t.Errorf(`
---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`
---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []string:
			if output := getLast(v); output != test.expected {
				t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []bool:
			if output := getLast(v); output != test.expected {
				t.Errorf(`---------------------------------
Test Failed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
Test Passed:
  input:    %v
  expected: %v
  actual:   %v
`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passed, failed)
}
