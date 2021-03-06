/*
binary search
*/

package main

import "fmt"

func binarySearchRecursion(a []int, target, low, high int) int {
	  if (low == high){
	    if a[low] == target{
	      return low
	      } else {
	        return -1
	      }
	  }
	  middle := int((low + high)/2)
	  if a[middle] == target{
	    return middle
	  }  else if target < a[middle]{
	    return binarySearchRecursion(a, target, low, middle-1)
	  } else {
	    return binarySearchRecursion(a, target, middle+1, high)
	  }
}

func initSearch(a []int) (int, int) {
	low := 0
	high := len(a) - 1
	if high < low {
		panic("Empty array.")
	}
	return low, high
}

func main() {

	sortedArr := []int{2, 5, 7, 11, 23, 32, 36}
	target := 23
	low, high := initSearch(sortedArr)
	targetIndex := binarySearchRecursion(sortedArr, target, low, high)
	fmt.Println(targetIndex, sortedArr[targetIndex])
}
