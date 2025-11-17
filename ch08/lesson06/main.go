package main

func getMessageCosts(messages []string) []float64 {
	slc := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		slc[i] = float64(len(messages[i])) * 0.01
	}
	return slc
}
