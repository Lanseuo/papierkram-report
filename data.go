package main

import "log"

type Data struct {
	receipts []Receipt
	invoices []Invoice
}

var data Data

func parseData() {
	receipts, err := ParseReceipts()
	if err != nil {
		log.Fatalln(err)
	}
	data.receipts = receipts

	invoices, err := ParseInvoices()
	if err != nil {
		log.Fatalln(err)
	}
	data.invoices = invoices
}
