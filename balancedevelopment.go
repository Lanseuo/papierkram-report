package main

import (
	"log"
	"sort"
	"strconv"
	"time"
)

type Month struct {
	Date       time.Time `json:"-"`
	DateString string    `json:"date"`
	Label      string    `json:"label"`
	Balance    float64   `json:"balance"`
}

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

func getEmptyMonths() []Month {
	now := time.Now()
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	months := []Month{}

	for index := 0; index < 12; index++ {
		months = append(months, Month{
			Date:       currentMonth,
			DateString: currentMonth.Format("02.01.2006"),
		})

		currentMonth = currentMonth.AddDate(0, -1, 0)
	}

	// Reverse order
	sort.SliceStable(months, func(i, j int) bool {
		return true
	})

	return months
}

// addLabelsToMonths adds a label consisting of the month of the year of each month
// e. g. Mar 19
func addLabelsToMonths(months []Month) []Month {
	for index, month := range months {
		monthName := month.Date.Month().String()[:3]
		year := strconv.Itoa(month.Date.Year())[2:]
		months[index].Label = monthName + " " + year
	}

	return months
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
