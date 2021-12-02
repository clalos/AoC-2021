/**
https://adventofcode.com/2021/day/2
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

// Movement holds the movement.
type Movement int

const (
	Forward Movement = iota + 1
	Up
	Down
)

func main() {
	var aim, depth, horizontalPos int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		read := strings.TrimSpace(scanner.Text())

		movement, err := readMovement(read)
		if err != nil {
			log.Fatal(err)
		}
		value, err := readValue(read)
		if err != nil {
			log.Fatal(err)
		}

		switch movement {
		case Forward:
			horizontalPos += value
			depth += aim * value
		case Up:
			aim -= value
		case Down:
			aim += value
		}
	}

	fmt.Println(fmt.Sprintf("Horizontal pos: %d", horizontalPos))
	fmt.Println(fmt.Sprintf("Depth: %d", depth))
	fmt.Println(fmt.Sprintf("Result: %d", horizontalPos*depth))
}

func readMovement(s string) (Movement, error) {
	switch {
	case strings.Contains(s, "forward"):
		return Forward, nil
	case strings.Contains(s, "up"):
		return Up, nil
	case strings.Contains(s, "down"):
		return Down, nil
	default:
		return 0, fmt.Errorf("unknown position")
	}
}

func readValue(s string) (int, error) {
	stringValue := strings.Split(s, " ")
	if len(stringValue) <= 1 {
		return 0, fmt.Errorf("invalid input")
	}
	value, err := strconv.Atoi(stringValue[1])
	if err != nil {
		return 0, err
	}

	return value, nil
}
