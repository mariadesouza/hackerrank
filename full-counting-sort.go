/*
In this challenge you need to print the data that accompanies each integer in a list. In addition, if two strings have the same integers, you need to print the strings in their original order. Hence, your sorting algorithm should be stable, i.e. the original order should be maintained for equal elements.

Insertion Sort and the simple version of Quicksort were stable, but the faster in-place version of Quicksort was not (since it scrambled around elements while sorting).

In cases where you care about the original order, it is important to use a stable sorting algorithm. In this challenge, you will use counting sort to sort a list while keeping the order of the strings (with the accompanying integer) preserved.

Challenge
In the previous challenge, you created a "helper array" that contains information about the starting position of each element in a sorted array. Can you use this array to help you create a sorted array of the original list?

Hint: You can go through the original array to access the strings. You can then use your helper array to help determine where to place those strings in the sorted array. Be careful about being one off.

Details and a Twist
You will be given a list that contains both integers and strings. Can you print the strings in order of their accompanying integers? If the integers for two strings are equal, ensure that they are print in the order they appeared in the original list.

The Twist - Your clients just called with an update. They don't want you to print the first half of the original array. Instead, they want you to print a dash for any element from the first half.

Input Format
- , the size of the list .
-  lines follow, each containing an integer  and a string .

Output Format
Print the strings in their correct order.

Constraints

 is even
 length(s)

The characters in every string in lowercase.

Sample Input

20
0 ab
6 cd
0 ef
6 gh
4 ij
0 ab
6 cd
0 ef
6 gh
0 ij
4 that
3 be
0 to
1 be
5 question
1 or
2 not
4 is
2 to
4 the
Sample Output

- - - - - to be or not to be - that is the question - - - -
Explanation
Below is the list in the correct order. The strings of each number were printed above for the second half of the array. Elements from the first half of the original array were replaced with dashes.

0 ab
0 ef
0 ab
0 ef
0 ij
0 to
1 be
1 or
2 not
2 to
3 be
4 ij
4 that
4 is
4 the
5 question
6 cd
6 gh
6 cd
6 gh
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	Index int
	Word  string
}

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	var half int
	half = n / 2
	m := make(map[int][]string)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		a := strings.Fields(line)
		//fmt.Println(a)
		if len(a) == 2 {
			ind, _ := strconv.Atoi(a[0])
			if count < half {
				m[ind] = append(m[ind], "-")
			} else {
				m[ind] = append(m[ind], a[1])
			}
			count++
			if count == n {
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		for _, value := range m[i] {
			fmt.Printf("%s ", value)
		}
	}
	fmt.Println()
}
