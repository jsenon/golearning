package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	MyPtr := new(int)
	fmt.Println(*MyPtr)
	fmt.Println(MyPtr)
	*MyPtr = 3
	fmt.Println(*MyPtr)
	one(MyPtr)
	fmt.Println(*MyPtr)
}
