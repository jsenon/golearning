package main

import (
	"fmt"
)

func main() {
	total := [5]float64{1, 2, 3, 4, 5}
	for i := 0; i < len(total); i++ {
		fmt.Printf("%f\n", total[i])
	}
}
