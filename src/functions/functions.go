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

func main() {
	array := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(array))
}
