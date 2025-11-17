package main

type cost struct {
	day int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dailyCosts := []float64{}
	for i := 0; i < len(costs); i++ {
		if day == costs[i].day {
			dailyCosts = append(dailyCosts, costs[i].value)
		}
	}
	return dailyCosts
}
