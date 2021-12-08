/**
https://adventofcode.com/2021/day/8
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		outputDigits := strings.Split(strings.Split(scanner.Text(), "|")[1], " ")
		for _, digit := range outputDigits {
			digit = strings.TrimSpace(digit)
			switch len(digit) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	fmt.Println(fmt.Sprintf("Digits with unique number of segments: %d", count))
}
