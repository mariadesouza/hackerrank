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

	buf := bufio.NewReader(os.Stdin)

	line, _ := buf.ReadString('\n')

	arr := make([]int, n)

	for _, w := range strings.Fields(line) {
		temp, _ := strconv.Atoi(w)
		arr[temp]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < arr[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
