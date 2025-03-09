package main

import (
	"bufio"
	"fmt"
	"os"
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

	opt, _ := getInput("Choose options (a - add item, s - save bill, t - add tip, x: exit): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Printf("You chose to add %v with price is %v:", name, price)
	case "s":
		fmt.Println("You chose to save the bill:", b.name)
	case "t":
		tip, _ := getInput("Tip amount: ", reader)
		fmt.Println("You chose to add a tip:", tip)
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
