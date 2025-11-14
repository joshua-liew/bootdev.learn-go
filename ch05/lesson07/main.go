package main

import "fmt"

func (e email) cost() int {
	costPerChar := 5
	if e.isSubscribed {
		costPerChar = 2
	}
	return len(e.body) * costPerChar
}

func (e email) format() string {
	subStr := "Not Subscribed"
	if e.isSubscribed {
		subStr = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subStr)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
