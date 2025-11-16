package main

func maxMessages(thresh int) int {
	baseRate := 100
	addFee := 1
	for numberOfMessages := 0; ; numberOfMessages++ {
		thresh -= baseRate + (addFee * numberOfMessages)
		if thresh < 0 {
			return numberOfMessages
		}
	}
}
