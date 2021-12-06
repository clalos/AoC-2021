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

// lanternfish represents a timer.
type lanternfish int

// timer returns the lanternfish's timer.
func (l *lanternfish) timer() int {
	if l == nil {
		return 0
	}
	return int(*l)
}

// setTimer sets the timer's value.
func (l *lanternfish) setTimer(t int) {
	*l = lanternfish(t)
}

// newDay decrease the fish internal timer.
// Spawn a new lanternfish if the timer is 0.
func (l *lanternfish) newDay() *lanternfish {
	if l.timer() == 0 {
		l.setTimer(6)
		newBorn := lanternfish(8)
		return &newBorn
	}
	l.setTimer(l.timer() - 1)
	return nil
}

func main() {
	fishes, err := loadTimers()
	if err != nil {
		log.Fatalf("failed to load the fishes: %v", err)
	}

	// Increase days.
	for day := 0; day < 80; day++ {
		for _, fish := range fishes {
			newBorn := fish.newDay()
			if newBorn != nil {
				fishes = append(fishes, newBorn)
			}
		}
	}

	fmt.Println(fmt.Sprintf("Total amount of fish: %d", len(fishes)))
}

// loadTimers loads the ages from STDIN.
func loadTimers() ([]*lanternfish, error) {
	var lanternfishes []*lanternfish
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringTimers := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	for _, timer := range stringTimers {
		age, err := strconv.Atoi(timer)
		if err != nil {
			return nil, err
		}
		lanternfish := lanternfish(age)
		lanternfishes = append(lanternfishes, &lanternfish)
	}
	return lanternfishes, nil
}
