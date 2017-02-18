package main

import (
	"fmt"
)

/*func main() {
	array := [5]float64{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%f\n", array[i])
	}
}
*/

func main() {
	array := [5]float64{1, 2, 3, 4, 5}
	var total float64 = 0
	for _, value := range array {
		total += value
	}
	fmt.Println(total / float64(len(array)))
	fmt.Printf("%f", total/float64(len(array)))
}
