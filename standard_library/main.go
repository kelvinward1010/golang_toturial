package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {
	// greeting := "Hello, i'm kelvin"

	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println((strings.Index(greeting, "l"))) // start index
	// fmt.Println(strings.Split(greeting, " "))

	// // The original value is unchanged
	// fmt.Printf("Original string value = %v", greeting)

	ages := []int{1, 2, 5, 6, 7, 8, 9, 3, 4}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 7)
	fmt.Println(index)

	names := []string{"f", "h", "a", "b", "c"}

	sort.Strings(names)
	fmt.Println(names)

	indexStr := sort.SearchStrings(names, "a")
	fmt.Println(indexStr)
}
