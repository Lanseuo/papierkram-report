package main

import (
	"math"
)

func calculateBalance(receipts []Receipt, invoices []Invoice) (float64, error) {
	var balance float64

	for _, receipt := range receipts {
		balance += receipt.NetAmount
	}

	for _, invoice := range invoices {
		balance += invoice.NetAmount
	}

	roundedBalance := math.Round(balance*100) / 100
	return roundedBalance, nil
}
