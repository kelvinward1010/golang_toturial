package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}

	name := [6]string{"K", "e", "l", "v", "i", "n"}
	name[1] = "o"

	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:4]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Tyler")
	fmt.Println(rangeOne)
}
