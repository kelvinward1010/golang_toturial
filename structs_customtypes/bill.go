package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string, items ...map[string]float64) bill {
	b := bill{
		name: name,
		items: func() map[string]float64 {
			if len(items) > 0 {
				return items[0]
			}
			return make(map[string]float64)
		}(),
		tip: 0,
	}

	return b
}

// format the bill
func (b bill) formatBill() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%10v: ...%v$ \n", k+": ", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("The %-25v ...$%v \n", "tip: ", b.tip)

	//total
	fs += fmt.Sprintf("The %-25v ...$%0.2f", "total: ", total)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) saveBill() {
	data := []byte(b.formatBill())

	os.MkdirAll("bills", os.ModePerm)

	fileName := strings.ReplaceAll(b.name, " ", "_")
	fileName = strings.ReplaceAll(fileName, "'", "_")

	err := os.WriteFile("bills/"+fileName+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file!")
}
