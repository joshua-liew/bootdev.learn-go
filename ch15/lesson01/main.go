package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	statusErr := em.recipient.updateStatus(em.status)
	if statusErr != nil {
		return fmt.Errorf("error updating user status: %w", statusErr)
	}
	trackErr := a.track(em.status)
	if trackErr != nil {
		return fmt.Errorf("error tracking user bounce: %w", trackErr)
	}
	return nil
}

