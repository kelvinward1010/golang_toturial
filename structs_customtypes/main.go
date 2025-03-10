package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Created a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options (a - add item,  t - add tip, s - save bill, x: exit): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}

		b.addItems(name, p)

		fmt.Printf("You chose to add %v with price is %v:", name, p)
		promtOptions(b)
	case "t":
		tip, _ := getInput("Tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promtOptions(b)
		}

		b.updateTip(t)

		fmt.Println("You chose to add a tip:", tip)
		promtOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("You chose to save the bill:", b.name)
	case "x":
		fmt.Println("Exited")
	default:
		fmt.Println("Invalid option")
		promtOptions(b)
	}
}

func main() {
	myBill := newBill("Kelvin's bill!", map[string]float64{"coffee": 3.99, "cake": 5.99})

	myBill.updateTip(10)
	myBill.addItems("Iphone X", 555.5)

	fmt.Println(myBill.formatBill())

	bill1 := createBill()
	fmt.Println(bill1)
	promtOptions(myBill)
}
