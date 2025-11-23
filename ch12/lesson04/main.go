package main

func addEmailsToQueue(emails []string) chan string {
	queueChannel := make(chan string, len(emails))
	func() {
		for _, email := range emails {
			queueChannel <- email
		}
	}()
	return queueChannel
}
