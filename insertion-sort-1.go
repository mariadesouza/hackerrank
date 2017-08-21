package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	insertSort(arr)

}

func insertSort(a []int) {
	n := len(a)
	i := n - 1
	x := a[n-1]

	for ; i > 0 && a[i-1] > x; i-- {
		a[i] = a[i-1]
		printArr(a)
	}
	a[i] = x
	printArr(a)
	//fmt.Println(a)
}

func printArr(arr []uint64, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}
