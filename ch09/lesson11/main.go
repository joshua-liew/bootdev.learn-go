package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			distinctWords[strings.ToLower(word)] = struct{}{}
		}
	}
	return len(distinctWords)
}
