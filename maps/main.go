package main

import (
	"fmt"
)

func main() {

	menus := map[string]float64{
		"soup":   64.55,
		"pie":    5.55,
		"salad":  6.99,
		"coffee": 3.55,
	}

	fmt.Println(menus)
	fmt.Println(menus["pie"])
	fmt.Println(menus["coffee"])

	//looping maps
	for index, value := range menus {
		fmt.Printf("[%v] : %v\n", index, value)
	}

	//inits as ley type
	phonebook := map[int]string{
		12345:  "kelvin",
		678910: "Ward",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[12345])

	phonebook[12345] = "KELVIN"
	fmt.Println(phonebook[12345])

	phonebook[12345] = "Changed"
	fmt.Println(phonebook[12345])
}
