package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fist, second := getInitials("Kelvin Ward")
	fmt.Printf("Fist name: %v \nSecond name %v", fist, second)

	fist2, second2 := getInitials("Kelvin")
	fmt.Printf("\nFist name 2: %v \nSecond name 2: %v", fist2, second2)
}
