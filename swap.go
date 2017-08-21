package main

import (
	"fmt"
)

func main() {

	var num1 int
	fmt.Scan(&num1)

	var num2 int
	fmt.Scan(&num2)

	fmt.Printf("Before: %d %d\n", num1, num2)

	num1, num2 = num2, num1

	fmt.Printf("After: %d %d\n", num1, num2)
}
