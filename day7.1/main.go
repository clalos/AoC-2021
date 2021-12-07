/**
https://adventofcode.com/2021/day/7
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	positions, err := loadValues()
	if err != nil {
		log.Fatalf("failed to read the input file: %v", err)
	}

	sort.Ints(positions)
	maxValue := positions[len(positions)-1]
	minValue := positions[0]

	fuelCosts := map[int]int{}
	for curValue := minValue; curValue <= maxValue; curValue++ {
		for i := 0; i < len(positions); i++ {
			steps := int(math.Abs(float64(curValue - positions[i])))
			fuelCosts[curValue] += triangleNumber(steps)
		}
	}

	var bestFuelCost, bestPosition int
	for i, value := range fuelCosts {
		if bestFuelCost == 0 || value < bestFuelCost {
			bestFuelCost = value
			bestPosition = i
		}
	}

	fmt.Println(fmt.Sprintf("Best fuel cost: %d", bestFuelCost))
	fmt.Println(fmt.Sprintf("Best position: %d", bestPosition))
}

// loadValues reads the input from STDIN.
func loadValues() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringValues := strings.Split(scanner.Text(), ",")
	output := make([]int, 0, len(stringValues))
	for _, stringValue := range stringValues {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return nil, err
		}
		output = append(output, intValue)
	}
	return output, nil
}

// triangleNumber calculates the triangle number of the given value.
func triangleNumber(n int) int {
	if n == 0 {
		return 0
	}
	return n + triangleNumber(n-1)
}
