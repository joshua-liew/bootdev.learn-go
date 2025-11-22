package main

type Analytics struct {
	MessagesTotal	int
	MessagesFailed	int
	MessagesSucceeded int
}

type Message struct {
	Recipient	string
	Success		bool
}

func analyzeMessage(aPtr *Analytics, m Message) {
	if m.Success {
		aPtr.MessagesSucceeded++
	} else {
		aPtr.MessagesFailed++
	}
	aPtr.MessagesTotal++
}
