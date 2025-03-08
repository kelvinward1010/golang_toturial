package main

import "fmt"

var points = []int{1, 2, 3, 4, 5}

func sayHello(n string) {
	fmt.Printf("Hello, %v", n)
}

func showScore() {
	fmt.Printf("You scored in many points: %0.2f", score)
}
