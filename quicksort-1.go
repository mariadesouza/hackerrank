/* Quick Sort
Step 1: Divide
Choose some pivot element, , and partition your unsorted array, , into three smaller arrays: , , and , where each element in , each element in , and each element in .

Challenge
Given  and , partition  into , , and  using the Divide instructions above. Then print each element in  followed by each element in , followed by each element in  on a single line. Your output should be space-separated.

Note: There is no need to sort the elements in-place; you can create two lists and stitch them together at the end.

Input Format

The first line contains  (the size of ).
The second line contains  space-separated integers describing  (the unsorted array). The first integer (corresponding to ) is your pivot element, .

Constraints

All elements will be unique.
Multiple answer can exists for the given test case. Print any one of them.
Output Format

On a single line, print the partitioned numbers (i.e.: the elements in , then the elements in , and then the elements in ). Each integer should be separated by a single space.

Sample Input

5
4 5 3 7 2
Sample Output

3 2 4 5 7
Explanation


Pivot: .
; ;

, so it's added to .
; ;

, so it's added to .
; ;

, so it's added to .
; ;

, so it's added to .
; ;

We then print the elements of , followed by , followed by , we get: 3 2 4 5 7.

This example is only one correct answer based on the implementation shown, but it is not the only correct answer (e.g.: another valid solution would be 2 3 4 5 7).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	for {
		inNumbers := scanner.Text()
		arr := strings.Split(inNumbers, " ")
		if len(arr) == n {
			err := quickSort(arr)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
	}

}

func quickSort(a []string) error {
	n := len(a)
	p, err := strconv.Atoi(a[0])
	if err != nil {
		return err
	}

	var rightArr, leftArr, equalArr []string
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(a[i])
		if err != nil {
			return err
		}
		if num < p {
			leftArr = append(leftArr, a[i])
		} else if num > p {
			rightArr = append(rightArr, a[i])
		} else {
			equalArr = append(equalArr, a[i])
		}
	}

	resArr := append(leftArr, equalArr...)
	resArr = append(resArr, rightArr...)

	fmt.Println(strings.Join(resArr, " "))

	return nil
}
