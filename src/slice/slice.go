package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7}
	slice3 := make([]int, 5)
	copy(slice3, slice2)
	fmt.Println(slice1, slice2, slice3)
}
