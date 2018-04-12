/*
We define a modified Fibonacci sequence using the following definition:

Given terms  and  where , term  is computed using the following relation:

For example, if  and ,

,
,
,
and so on.
Given three integers, , , and , compute and print term  of a modified Fibonacci sequence.

Note: The value of  may far exceed the range of a -bit integer. Many submission languages have libraries that can handle such large results but, for those that don't (e.g., C++), you will need to compensate for the size of the result.

Input Format

A single line of three space-separated integers describing the respective values of , , and .

Constraints

 may far exceed the range of a -bit integer.
Output Format

Print a single integer denoting the value of term  in the modified Fibonacci sequence where the first two terms are  and .

Sample Input

0 1 5
Sample Output

5
Explanation

The first two terms of the sequence are  and , which gives us a modified Fibonacci sequence of . Because , we print term .
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var a, b, target int64
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &target)
	t1 := big.NewInt(a)
	t2 := big.NewInt(b)
	if target > 0 {
		fmt.Println(modifiedFibonacci(t1, t2, 3, target))
	}

}

func modifiedFibonacci(t1, t2 *big.Int, n, target int64) *big.Int {

	if target == 1 {
		return t1
	}
	if target == 2 {
		return t2
	}
	//fmt.Printf("%d + sq(%d) : %d\n", t1, t2, target)
	next := new(big.Int)
	sq := new(big.Int)
	sq.Mul(t2, t2)
	next.Add(t1, sq)
	if n == target {
		return next
	} else {
		return modifiedFibonacci(t2, next, n+1, target)
	}

}
