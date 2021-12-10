/**
https://adventofcode.com/2021/day/6
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
	fishes, err := loadTimers()
	if err != nil {
		log.Fatalf("failed to load the fishes: %v", err)
	}

	// One counter for each timer.
	var counters [9]uint

	// Initial setup.
	for _, fish := range fishes {
		counters[fish]++
	}

	for day := 0; day < 256; day++ {
		zeros := counters[0]
		counters[0] = counters[1]
		counters[1] = counters[2]
		counters[2] = counters[3]
		counters[3] = counters[4]
		counters[4] = counters[5]
		counters[5] = counters[6]
		counters[6] = counters[7] + zeros
		counters[7] = counters[8]
		counters[8] = zeros
	}

	// Sum counters.
	var totalFishes uint
	for i := 0; i < 9; i++ {
		totalFishes += counters[i]
	}

	fmt.Println(fmt.Sprintf("Total amount of fish: %d", totalFishes))
}

// loadTimers loads the ages from STDIN.
func loadTimers() ([]uint, error) {
	var lanternfishes []uint
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringTimers := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	for _, timer := range stringTimers {
		age, err := strconv.ParseUint(timer, 10, 64)
		if err != nil {
			return nil, err
		}
		lanternfishes = append(lanternfishes, uint(age))
	}
	return lanternfishes, nil
}
