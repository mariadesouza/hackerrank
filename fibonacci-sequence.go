/*
Fibonacci Sequence

*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fibonacci(n)
}

func fibonacci(n int) {

	//var a[]int
	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			a[i] = i
		} else {
			a[i] = a[i-1] + a[i-2]
		}
	}

	fmt.Println(a)

}
