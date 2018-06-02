/*
NUmber swap using xor
*/

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	a, b = swap(a, b)
	fmt.Println(a, b)
}

func swap(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	a = a ^ b
	return a, b
}
