package main

import (
	"fmt"
)

func main() {
	name := "Thế giới"
	age := 23
	myname := "Tyler Locke"
	fmt.Println("Xin chào, " + name + "! \n")
	fmt.Println("New line!")
	fmt.Println("Hello VietNam")
	// fmt.Println("I'm ", age, "years old!")

	// Formatted string %_ = format specifier
	fmt.Printf("My name is %v. I'm %v years old \n", myname, age)
	fmt.Printf("My name is %q. I'm %v years old \n", myname, age)
	fmt.Printf("I'm %T years old\n", age)
	fmt.Printf("Your score %f years old\n", 25.25)
	fmt.Printf("Your score %0.1f years old\n", 25.25)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My name is %v. I'm %v years old \n", myname, age)
	fmt.Println("The string is: ", str)
}
