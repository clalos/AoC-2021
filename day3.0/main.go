/**
https://adventofcode.com/2021/day/3
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	// We assume all the lines have the same length of the first one.
	var gamma, epsilon string
	for i := 0; i < len(lines[0]); i++ {
		var zeros, ones int

		for _, line := range lines {
			if string(line[i]) == "0" {
				zeros++
				continue
			}
			ones++
		}

		switch {
		case zeros > ones:
			gamma += "0"
			epsilon += "1"
		case zeros < ones:
			gamma += "1"
			epsilon += "0"
		default:
			panic(fmt.Sprintf("Even count for the column %d", i))
		}
	}

	gammaRate, err := strconv.ParseUint(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonRate, err := strconv.ParseUint(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Gamma rate: %d", gammaRate))
	fmt.Println(fmt.Sprintf("Epsilon rate: %d", epsilonRate))
	fmt.Println(fmt.Sprintf("Power consuption: %d", gammaRate*epsilonRate))
}
