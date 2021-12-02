package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
https://adventofcode.com/2021/day/1
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var depths []int

	for scanner.Scan() {
		read, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, read)
	}

	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}

	fmt.Println(fmt.Sprintf("%d increases", increases))
}
