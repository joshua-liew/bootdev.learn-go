package main

import (
	"errors"
)

type customer struct {
	id		int
	balance	float64
}

type transactionType string

const (
	transactionDeposit		transactionType = "deposit"
	transactionWithdrawal	transactionType = "withdrawal"
)

type transaction struct {
	customerID		int
	amount			float64
	transactionType	transactionType
}

func updateBalance(cp *customer, t transaction) error {
	if cp == nil {
		return errors.New("customer pointer is nil")
	}

	switch t.transactionType {
	case transactionDeposit:
		cp.balance += t.amount
	case transactionWithdrawal: {
		newBalance := cp.balance - t.amount
		if newBalance < 0 {
			return errors.New("insufficient funds")
		}
		cp.balance = newBalance
	}
	default:
		return errors.New("unknown transaction type")
	}

	return nil
}
