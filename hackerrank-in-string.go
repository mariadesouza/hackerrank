/*

HackerRank in a String!

 Practice
 Compete
 Jobs
 Rank
 Leaderboard
  mariadesouza
Dashboard  Algorithms  Strings  HackerRank in a String!
Badge Progress(Details)
Points: 376.00 Rank: 106036
Your HackerRank in a String! submission got 20.00 points.
Try the Next Challenge   |   Try a Random Challenge
×
HackerRank in a String!
by shashank21j
Problem
Submissions
Leaderboard
Discussions
Editorial
We say that a string, , contains the word hackerrank if a subsequence of the characters in  spell the word hackerrank. For example, haacckkerrannkk does contain hackerrank, but haacckkerannk does not (the characters all appear in the same order, but it's missing a second r).

More formally, let  be the respective indices of h, a, c, k, e, r, r, a, n, k in string . If  is true, then  contains hackerrank.

You must answer  queries, where each query consists of a string, . For each query, print YES on a new line if contains hackerrank; otherwise, print NO instead.

Input Format

The first line contains an integer denoting  (the number of queries).
Each line of the  subsequent lines contains a single string denoting .

Constraints

Output Format

For each query, print YES on a new line if  contains hackerrank; otherwise, print NO instead.

Sample Input 0

2
hereiamstackerrank
hackerworld
Sample Output 0

YES
NO
Explanation 0

We perform the following  queries:


The characters of hackerrank are bolded in the string above. Because the string contains all the characters in hackerrank in the same exact order as they appear in hackerrank, we print YES on a new line.
 does not contain the last three characters of hackerrank, so we print NO on a new line.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var a []string
	var line string
	for i := 0; i < n; i++ {
		fmt.Scanln(&line)
		a = append(a, line)
	}

	expected := "hackerrank"

	//fmt.Println(a, expected)

	for i := 0; i < n; i++ {
		strlen := utf8.RuneCountInString(a[i])
		strcount := 0
		for k := 0; k < strlen; k++ {
			if a[i][k] = expected[strcount] {
				strcount++
			}
			if strcount == 10 {
                break
      }
		}
		if strcount == 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
