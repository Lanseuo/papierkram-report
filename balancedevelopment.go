package main

import (
	"log"
	"strconv"
	"time"
)

func calculateBalanceDevelopment(receipts []Receipt, invoices []Invoice) ([]Month, error) {
	months := getEmptyMonths()
	months = addLabelsToMonths(months)

	for monthIndex, month := range months {
		balanceAfterMonth := 0.0

		for _, receipt := range receipts {
			if paymentDateWasBeforeEndOfMonth(receipt.PaymentDate, month.Date) {
				balanceAfterMonth += receipt.NetAmount
			}
		}

		for _, invoice := range invoices {
			if paymentDateWasBeforeEndOfMonth(invoice.LastPaymentDate, month.Date) {
				balanceAfterMonth += invoice.NetAmount
			}
		}

		months[monthIndex].Balance = roundBalance(balanceAfterMonth)
	}

	return months, nil
}

func paymentDateWasBeforeEndOfMonth(paymentDate string, monthDate time.Time) bool {
	if paymentDate == "" {
		return false
	}

	paymentMonth, monthErr := strconv.Atoi(paymentDate[3:5])
	paymentYear, yearErr := strconv.Atoi(paymentDate[6:])
	if monthErr != nil || yearErr != nil {
		log.Fatalln("Unable to parse date from receipt", monthErr, yearErr)
	}

	isInYearsBefore := paymentYear < monthDate.Year()
	isInSameYearButBefore := paymentYear == monthDate.Year() && paymentMonth <= int(monthDate.Month())
	return isInYearsBefore || isInSameYearButBefore
}
