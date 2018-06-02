/*
two strings are anagrams of each other
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1 string
	fmt.Scan(&str1)

	var str2 string
	fmt.Scan(&str2)

	if isAnagram(str1, str2) {
		println("Is Anagram")
	} else {
		println("Is Not Anagram")
	}
	println(str1, str2)
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)
	if strings.Join(s1, "") != strings.Join(s2, "") {
		return false
	}
	return true
}
