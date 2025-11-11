package main

func getMonthlyPrice(tier string) int {
	penniesPerDollar := 100
	switch tier {
		case "basic":
			return 100 * penniesPerDollar
		case "premium":
			return 150 * penniesPerDollar
		case "enterprise":
			return 500 * penniesPerDollar
		default:
			return 0
	}
}
