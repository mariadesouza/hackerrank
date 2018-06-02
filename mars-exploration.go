/*

Mars Exploration
Letters in some of the SOS messages are altered by cosmic radiation during transmission. Given the signal received by Earth as a string, , determine how many letters of Sami's SOS have been changed by radiation.

Input Format

There is one line of input: a single string, .

Note: As the original message is just SOS repeated  times, 's length will be a multiple of .

Constraints

 will contain only uppercase English letters.
Output Format

Print the number of letters in Sami's message that were altered by cosmic radiation.

Sample Input 0

SOSSPSSQSSOR
Sample Output 0

3
Sample Input 1

SOSSOT
Sample Output 1

1
Explanation

Sample 0

 = SOSSPSSQSSOR, and signal length . Sami sent  SOS messages (i.e.: ).

Expected signal: SOSSOSSOSSOS
Recieved signal: SOSSPSSQSSOR

We print the number of changed letters, which is .

Sample 1

 = SOSSOT, and signal length . Sami sent  SOS messages (i.e.: ).

Expected Signal: SOSSOS
Received Signal: SOSSOT

We print the number of changed letters, which is .
*/
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var a string
	fmt.Scanln(&a)

	strlen := utf8.RuneCountInString(a)

	expected := strings.Repeat("SOS", strlen/3)

	//fmt.Println(a, expected)

	var count int
	for i := 0; i < strlen; i++ {
		if a[i] != expected[i] {
			count++
		}
	}

	fmt.Println(count)

}
