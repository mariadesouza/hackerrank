package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter Len:")
	var arrLen int
	fmt.Scanln(&arrLen)

	var a []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter %d Numbers separated by a space:\n", arrLen)
		scanner.Scan()
		inNumbers := scanner.Text()
		fmt.Println(inNumbers)
		a = strings.Split(inNumbers, " ")

		fmt.Println(len(a), arrLen)
		if len(a) == arrLen {
			break
		}
	}

	total := 0
	for i := range a {
		numeric, _ := strconv.Atoi(a[i])
		total += numeric
	}
	fmt.Println("Total:", total)
}
