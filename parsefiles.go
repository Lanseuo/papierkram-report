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
	State             string
	NetAmount         float64
	Vat               string
	GrossAmount       float64
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

		receipts = append(receipts, Receipt{
			ID:                line[0],
			Subject:           line[1],
			Supplier:          line[2],
			Date:              line[3],
			PaymentTerm:       line[4],
			PaymentDate:       line[5],
			State:             line[6],
			NetAmount:         parseAmount(line[7]),
			Vat:               line[8],
			GrossAmount:       parseAmount(line[9]),
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

func parseAmount(s string) float64 {
	sWithPoint := strings.Replace(s, ",", ".", -1)
	amount, err := strconv.ParseFloat(sWithPoint, 64)
	if err != nil {
		log.Fatal(err)
	}
	return amount
}
