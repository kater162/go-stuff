package main

import (
	"fmt"
)

func main() {
    // takes 2 numbers from user and returns their sum
	fmt.Println("Enter two numbers: ")
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// takes 2 numbers from user and returns their difference
	fmt.Println("Enter two numbers: ")
	var a1, b1 int
	fmt.Scan(&a1, &b1)
	diff := a1 - b1
	fmt.Printf("%d - %d = %d\n", a1, b1, diff)
}
