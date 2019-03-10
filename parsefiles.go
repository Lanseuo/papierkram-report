package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Receipt struct {
	ID                string
	Subject           string
	Supplier          string
	Date              string
	PaymentTerm       string
	PaymentDate       string
	Status            string
	NetAmount         float64
	VatAmount         float64
	GrossAmount       float64
	EquityRatioAmount string
	PlaceOfOrigin     string
	Categories        string
	Projects          string
	ProjectCustomer   string
	ProjectCustomerID string
	Extra             string
	Note              string
}

type Invoice struct {
	ID              string
	Subject         string
	Customer        string
	Project         string
	CustomerID      string
	Date            string
	PaymentTerm     string
	Status          string
	Type            string
	NetAmount       float64
	VatAmount       float64
	GrossAmount     float64
	DiscountAmount  float64
	PaidAmount      float64
	LastPaymentDate string
	Note            string
}

func ParseReceipts() ([]Receipt, error) {
	items, err := parseCSV("/tmp/papierkram-report/CSV/Belege.csv")
	if err != nil {
		return nil, err
	}

	var receipts []Receipt
	for _, receipt := range items {
		receipts = append(receipts, Receipt{
			ID:                receipt[0],
			Subject:           receipt[1],
			Supplier:          receipt[2],
			Date:              receipt[3],
			PaymentTerm:       receipt[4],
			PaymentDate:       receipt[5],
			Status:            receipt[6],
			NetAmount:         parseAmount(receipt[7]),
			VatAmount:         parseAmount(receipt[8]),
			GrossAmount:       parseAmount(receipt[9]),
			EquityRatioAmount: receipt[10],
			PlaceOfOrigin:     receipt[11],
			Categories:        receipt[12],
			Projects:          receipt[13],
			ProjectCustomer:   receipt[14],
			ProjectCustomerID: receipt[15],
			Extra:             receipt[16],
			Note:              receipt[17],
		})
	}

	return receipts, nil
}

func ParseInvoices() ([]Invoice, error) {
	items, err := parseCSV("/tmp/papierkram-report/CSV/Rechnungen.csv")
	if err != nil {
		return nil, err
	}

	var invoices []Invoice
	for _, receipt := range items {
		invoices = append(invoices, Invoice{
			ID:              receipt[0],
			Subject:         receipt[1],
			Customer:        receipt[2],
			Project:         receipt[3],
			CustomerID:      receipt[4],
			Date:            receipt[5],
			PaymentTerm:     receipt[6],
			Status:          receipt[7],
			Type:            receipt[8],
			NetAmount:       parseAmount(receipt[9]),
			VatAmount:       parseAmount(receipt[10]),
			GrossAmount:     parseAmount(receipt[11]),
			DiscountAmount:  parseAmount(receipt[12]),
			PaidAmount:      parseAmount(receipt[13]),
			LastPaymentDate: receipt[14],
			Note:            receipt[15],
		})
	}

	return invoices, nil
}

func parseCSV(filepath string) ([][]string, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	var lines [][]string

	for index := 0; true; index++ {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Remove header
		if index == 0 {
			continue
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func parseAmount(s string) float64 {
	sWithPoint := strings.Replace(s, ",", ".", -1)
	amount, err := strconv.ParseFloat(sWithPoint, 64)
	if err != nil {
		log.Fatal(err)
	}
	return amount
}
