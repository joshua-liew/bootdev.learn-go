package main

func bulkSend(numMessages int) float64 {
	var totalCost float64
	baseRate := 1.0
	incCost := 0.01
	for i := 0; i < numMessages; i++ {
		totalCost += baseRate + (incCost * float64(i))
	}
	return totalCost
}

