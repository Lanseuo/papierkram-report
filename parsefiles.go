package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Receipt struct {
	ID                string
	Subject           string
	Supplier          string
	Date              string
	PaymentTerm       string
	PaymentDate       string
	State             string
	NetAmount         string
	Vat               string
	GrossAmount       string
	EquityRatio       string
	PlaceOfOrigin     string
	Category          string
	ProjectCustomer   string
	ProjectCustomerID string
	Extra             string
	Note              string
}

func ParseReceipts() ([]Receipt, error) {
	csvFile, err := os.Open("/tmp/papierkram-report/CSV/Belege.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	var receipts []Receipt

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		receipts = append(receipts, Receipt{
			ID:                line[0],
			Subject:           line[1],
			Supplier:          line[2],
			Date:              line[3],
			PaymentTerm:       line[4],
			PaymentDate:       line[5],
			State:             line[6],
			NetAmount:         line[7],
			Vat:               line[8],
			GrossAmount:       line[9],
			EquityRatio:       line[10],
			PlaceOfOrigin:     line[11],
			Category:          line[12],
			ProjectCustomer:   line[13],
			ProjectCustomerID: line[14],
			Extra:             line[15],
			Note:              line[16],
		})
	}

	return receipts, nil
}
