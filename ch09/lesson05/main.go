package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, messagedUser := range messagedUsers {
		if _, exists := validUsers[messagedUser]; exists {
			validUsers[messagedUser]++
		}
	}
}
