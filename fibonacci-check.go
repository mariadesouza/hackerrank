/*
is a fibonacci
*/

package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(nr uint64) bool {
	s := math.Sqrt(float64(nr))
	return uint64(s)*uint64(s) == nr
}
func isFibonacci(nr uint64) bool {
	return isPerfectSquare(5*nr*nr+4) || isPerfectSquare(5*nr*nr-4)
}

func main() {

	var n uint64
	fmt.Scanf("%v", &n)

	if (n > 0 && n < 4) || isFibonacci(n) {
		fmt.Println("Is Fibonacci")
	} else {
		fmt.Println("Is Not a Fibonacci")
	}
}
