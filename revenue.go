package main

import (
	"log"
	"strconv"
	"time"
)

func calculateRevenue(receipts []Receipt, invoices []Invoice) ([]Month, error) {
	months := getEmptyMonths()
	months = addLabelsToMonths(months)

	for monthIndex, month := range months {
		earningsInMonth := 0.0
		spendingsInMonth := 0.0

		for _, receipt := range receipts {
			if paymentDateInMonth(receipt.PaymentDate, month.Date) {
				spendingsInMonth -= receipt.NetAmount
			}
		}

		for _, invoice := range invoices {
			if paymentDateInMonth(invoice.LastPaymentDate, month.Date) {
				earningsInMonth += invoice.NetAmount
			}
		}

		months[monthIndex].EearningsAmount = roundBalance(earningsInMonth)
		months[monthIndex].SpendingsAmount = roundBalance(spendingsInMonth)
	}

	return months, nil
}

func paymentDateInMonth(paymentDate string, monthDate time.Time) bool {
	if paymentDate == "" {
		return false
	}

	paymentMonth, monthErr := strconv.Atoi(paymentDate[3:5])
	paymentYear, yearErr := strconv.Atoi(paymentDate[6:])
	if monthErr != nil || yearErr != nil {
		log.Fatalln("Unable to parse date from receipt", monthErr, yearErr)
	}

	return paymentMonth == int(monthDate.Month()) && paymentYear == monthDate.Year()
}
