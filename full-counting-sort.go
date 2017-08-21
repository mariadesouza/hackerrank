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
