package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is: ", i)
	// }

	names := []string{"f", "h", "a", "b", "c"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("value of x is: ", names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("[%v] = %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("[] = %v \n", value)
		value = "new string"
	}

	fmt.Println(names)

}
