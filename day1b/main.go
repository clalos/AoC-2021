/**
https://adventofcode.com/2021/day/1
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
	scanner := bufio.NewScanner(os.Stdin)
	var depths []int

	for scanner.Scan() {
		read, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			panic(err)
		}
		depths = append(depths, read)
	}

	increases := 0
	for i := 0; i < len(depths)-3; i++ {
		if depths[i]+depths[i+1]+depths[i+2] <
			depths[i+1]+depths[i+2]+depths[i+3] {
			increases++
		}
	}

	fmt.Println(fmt.Sprintf("%d increases", increases))
}
