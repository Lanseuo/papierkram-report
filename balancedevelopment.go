package main

import (
	"sort"
	"time"
)

type Month struct {
	Date    string  `json:"date"`
	Label   string  `json:"label"`
	Balance float64 `json:"balance"`
}

func calculateBalanceDevelopment(receipts []Receipt, invoices []Invoice) ([]Month, error) {
	months := getEmptyMonths()
	return months, nil
}

func getEmptyMonths() []Month {
	now := time.Now()
	currentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	months := []Month{}

	for index := 0; index < 12; index++ {
		months = append(months, Month{
			Date: currentMonth.Format("02.01.2006"),
		})

		currentMonth = currentMonth.AddDate(0, -1, 0)
	}

	// Reverse order
	sort.SliceStable(months, func(i, j int) bool {
		return true
	})

	return months
}
