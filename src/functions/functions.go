package main

import (
	"fmt"
)

func average(xs []float64) float64 {
	var total float64 = 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	//	array := []float64{1, 2, 3, 4, 5}
	//	fmt.Println(average(array))
	fmt.Println("Enter your Number: ")
	var input uint
	fmt.Scanf("%d", &input)
	fmt.Println("Your Result :", factorial(input))
}

// Variadic functions wrote func Myfunc(args ...int) int {}
// Function calling functon called functional programming
// Function defer is called after the function complete
