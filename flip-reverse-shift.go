/*
reverse bits
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(flipBits1(a[i]))
	}

}

func flipBits1(n int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		ans <<= 1
		ans |= n & 1
		n >>= 1
	}
	return ans
}
