package main

import (
	"strings"
)

func removeProfanity(message *string) {
	targetWords := []string{"fubb", "shiz", "witch"}
	for _, word := range targetWords {
		if strings.Contains(*message, word) {
			*message = strings.ReplaceAll(*message, word, strings.Repeat("*", len(word)))
		}
	}
}
