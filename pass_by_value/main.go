package main

import "fmt"

func updateName(_ string) string {
	return "Kelvin"
}

func updateMenus(y map[string]float64) {
	y["coffee"] = 1.99
}

func main() {

	name := "Tyler"

	name = updateName(name)

	fmt.Println(name)

	menus := map[string]float64{
		"soup":   64.55,
		"pie":    5.55,
		"salad":  6.99,
		"coffee": 3.55,
	}

	updateMenus(menus)
	fmt.Println(menus["coffee"])
}
