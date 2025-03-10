package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (s circle) area() float64 {
	return math.Pi * s.radius * s.radius
}

func (s circle) circumf() float64 {
	return math.Pi * s.radius * 2
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.area())
}

func main() {
	shapes := []shape{
		square{length: 4},
		circle{radius: 5.5},
	}

	for _, v := range shapes {
		printShapeInfo(v)
	}
}
