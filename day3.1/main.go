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

	oxygenRate, err := calculateRate(lines, 0, true)
	if err != nil {
		log.Fatalf("Oxygen rate failed to return: %v", err)
	}

	co2Rate, err := calculateRate(lines, 0, false)
	if err != nil {
		log.Fatalf("CO2 rate failed to return: %v", err)
	}

	fmt.Println(fmt.Sprintf("Oxygen rate: %d", oxygenRate))
	fmt.Println(fmt.Sprintf("CO2 rate: %d", co2Rate))
	fmt.Println(fmt.Sprintf("Life supporting rating: %d", oxygenRate*co2Rate))
}

func calculateRate(input []string, pos int, withMostCommon bool) (uint64, error) {
	if len(input) == 1 {
		rate, err := strconv.ParseUint(input[0], 2, 64)
		if err != nil {
			return 0, err
		}
		return rate, nil
	}

	if pos == len(input[0]) {
		return 0, fmt.Errorf("failed to determine the output. Multiple values left in the list: %v", input)
	}

	keep := map[string][]string{}
	for _, line := range input {
		if string(line[pos]) == "0" {
			keep["0"] = append(keep["0"], line)
			continue
		}
		keep["1"] = append(keep["1"], line)
	}

	switch {
	case len(keep["0"]) > len(keep["1"]):
		if withMostCommon {
			return calculateRate(keep["0"], pos+1, true)
		}
		return calculateRate(keep["1"], pos+1, false)
	default:
		if withMostCommon {
			return calculateRate(keep["1"], pos+1, true)
		}
		return calculateRate(keep["0"], pos+1, false)
	}
}
