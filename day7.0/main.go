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

	medianValue := median(positions)

	// Calculate distance from the median.
	// Sum all the distances.
	var totalFuel int
	for i := 0; i < len(positions); i++ {
		fuel := positions[i] - medianValue
		totalFuel += int(math.Abs(float64(fuel)))
	}

	fmt.Println(fmt.Sprintf("Best position: %d", medianValue))
	fmt.Println(fmt.Sprintf("Total fuel: %d", totalFuel))
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

// Median calculates the median of the values store in the input slice.
func median(input []int) int {
	sort.Ints(input)
	medianNumber := len(input) / 2

	// Odd
	if medianNumber%2 != 0 {
		return input[medianNumber]
	}

	// Even
	return (input[medianNumber-1] + input[medianNumber]) / 2
}
