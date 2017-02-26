package main

import (
	"fmt"
	m "packaging/math"
)

func Max(xs []float64) float64 {
	var ma float64
	for i, x := range xs {
		if i == 0 || x > ma {
			ma = x
		}
	}
	return ma
}

func Min(xs []float64) float64 {
	var mi float64
	for i, x := range xs {
		if i == 0 || x < mi {
			mi = x
		}
	}
	return mi
}

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	max := Max(xs)
	min := Min(xs)
	fmt.Println("Average:", avg)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
