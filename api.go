package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	balance, expectedBalance, err := calculateBalance(data.receipts, data.invoices)
	if err != nil {
		log.Fatalln(err)
	}

	responseData := struct {
		Balance         float64 `json:"balance"`
		ExpectedBalance float64 `json:"expectedBalance"`
	}{
		Balance:         balance,
		ExpectedBalance: expectedBalance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func balanceDevelopmentHandler(w http.ResponseWriter, r *http.Request) {
	months, err := calculateBalanceDevelopment(data.receipts, data.invoices)
	if err != nil {
		log.Fatalln(err)
	}

	responseData := struct {
		Months []Month `json:"months"`
	}{
		Months: months,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func revenueHandler(w http.ResponseWriter, r *http.Request) {
	months, err := calculateRevenue(data.receipts, data.invoices)
	if err != nil {
		log.Fatalln(err)
	}

	responseData := struct {
		Months []Month `json:"months"`
	}{
		Months: months,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
