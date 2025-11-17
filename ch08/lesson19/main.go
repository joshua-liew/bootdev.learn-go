package main

import (
	"strings"
)

type sms struct {
	id 		string
	content	string
	tags	[]string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	res := []sms{}
	for _, message := range messages {
		message.tags = tagger(message)
		if message.tags == nil {
			return nil
		}
		res = append(res, message)
	}
	return res
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
