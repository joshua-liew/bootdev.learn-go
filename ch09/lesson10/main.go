package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		key := []rune(name)[0]
		if _, ok := nameCounts[key]; !ok {
			nameCounts[key] = make(map[string]int)
		}
		nameCounts[key][name]++
	}
	return nameCounts
}
