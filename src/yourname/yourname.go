package main

import "fmt"

func main() {
	fmt.Println("Enter your Name: ")
	var input string
	fmt.Scanf("%s", &input)

	yourName := input
	fmt.Println("Your Name is:", yourName)
}
