package main

import (
	"math"
)

func calculateBalance(receipts []Receipt, invoices []Invoice) (balance, expectedBalance float64, err error) {
	for _, receipt := range receipts {
		expectedBalance += receipt.NetAmount

		if receipt.PaymentDate != "" {
			balance += receipt.NetAmount
		}
	}

	for _, invoice := range invoices {
		expectedBalance += invoice.NetAmount

		if invoice.PaidAmount == invoice.GrossAmount {
			balance += invoice.NetAmount
		}
	}

	roundedBalance := roundBalance(balance)
	expectedRoundedBalance := roundBalance(expectedBalance)
	return roundedBalance, expectedRoundedBalance, nil
}

func roundBalance(balance float64) float64 {
	return math.Round(balance*100) / 100
}
