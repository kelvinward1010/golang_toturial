package main

import "fmt"

func main() {
	myBill := newBill("Kelvin's bill!", map[string]float64{"coffee": 3.99, "cake": 5.99})

	fmt.Println(myBill.formatBill())
}
