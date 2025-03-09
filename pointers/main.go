package main

import "fmt"

func updateName(x *string) {
	*x = "Kelvin"
}
func main() {

	name := "Tyler"
	//name = updateName(name)

	fmt.Println("The location of name is: ", &name)
	fmt.Println("The memory address: ", name)

	m := &name

	fmt.Println("The location of name is: ", m)
	fmt.Println("The memory address: ", *m)

	updateName(m)

	fmt.Println("The location of name is: ", m)
	fmt.Println("The memory address: ", *m)

	fmt.Println(name)
}
