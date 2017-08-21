package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	//Alice
	var a []string
	scanner.Scan()
	for {
		aScores := scanner.Text()
		a = strings.Split(aScores, " ")
		if len(a) == 3 {
			break
		}
	}
	//Bob
	var b []string
	scanner.Scan()
	for {
		bScores := scanner.Text()
		b = strings.Split(bScores, " ")
		if len(b) == 3 {
			break
		}
	}

	aScore, bScore := 0, 0
	for i := 0; i < 3; i++ {
		aa, _ := strconv.Atoi(a[i])
		bb, _ := strconv.Atoi(b[i])
		if aa > bb {
			aScore += 1
		} else {
			if aa < bb {
				bScore += 1
			}
		}
	}
	fmt.Printf("%d %d", aScore, bScore)
}
