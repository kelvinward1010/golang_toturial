package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string, items map[string]float64) bill {
	b := bill{
		name:  name,
		items: items,
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) formatBill() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%10v: ...$%10v \n", k+": ", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("The %-25v ...$%0.2f", "total: ", total)
	return fs
}
