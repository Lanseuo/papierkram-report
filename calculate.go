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

	roundedBalance := math.Round(balance*100) / 100
	expectedRoundedBalance := math.Round(expectedBalance*100) / 100
	return roundedBalance, expectedRoundedBalance, nil
}
