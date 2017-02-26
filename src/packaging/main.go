package main

import (
	"fmt"
	m "packaging/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	max := m.Max(xs)
	min := m.Min(xs)
	fmt.Println("Average:", avg)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
