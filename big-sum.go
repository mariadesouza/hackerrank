package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var arrLen int
	fmt.Scanln(&arrLen)

	scanner := bufio.NewScanner(os.Stdin)

	var aa []string
	scanner.Scan()
	for {
		strNums := scanner.Text()
		aa = strings.Split(strNums, " ")
		if len(aa) == arrLen {
			break
		}
	}

	var total int64
	for i := 0; i < arrLen; i++ {
		var num int64
		num, err := strconv.ParseInt(aa[i], 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		total += num
	}

	fmt.Println(total)
}
