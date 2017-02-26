package main

import (
	"fmt"
	"math"
)

// Methods

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

// Structs
//func Area(c Circle) float64 {
//	return math.Pi * c.r * c.r
//}

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{0, 0, 0}
	fmt.Println("Enter your r: ")
	fmt.Scanf("%f", &c.r)
	fmt.Println("Your number:", c.r)
	//	fmt.Println("Area is:", Area(c))
	fmt.Println("Area is:", c.Area())

}
