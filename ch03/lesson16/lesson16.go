package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return len(s) * 2
	}, intro)
	printCostReport(func(s string) int {
		return len(s) * 3
	}, body)
	printCostReport(func(s string) int {
		return len(s) * 4
	}, outro)
}

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
