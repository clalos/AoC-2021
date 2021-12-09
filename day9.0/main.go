/**
https://adventofcode.com/2021/day/9
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	heightmap, err := load()
	if err != nil {
		log.Fatalf("failed to read the input: %v", err)
	}

	var sum int
	for i := 1; i < len(heightmap)-1; i++ {
		for j := 1; j < len(heightmap[i])-1; j++ {
			point := heightmap[i][j]
			if point < heightmap[i-1][j] &&
				point < heightmap[i+1][j] &&
				point < heightmap[i][j-1] &&
				point < heightmap[i][j+1] {
				sum += point + 1
			}
		}
	}

	fmt.Println(fmt.Sprintf("Sum of risk levels: %d", sum))
}

// load reads the matrix from STDIN.
// It wraps the input matrix in a frame of 9 values.
func load() ([][]int, error) {
	var heightmap [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var row []int
		for pos, stringRune := range scanner.Text() {
			value, err := strconv.Atoi(string(stringRune))
			if err != nil {
				return nil, fmt.Errorf("failed to conver the number %s", string(stringRune))
			}
			if pos == 0 {
				row = append(row, 9)
			}
			row = append(row, value)
		}
		row = append(row, 9)
		heightmap = append(heightmap, row)
	}
	nineRow := [][]int{make([]int, len(heightmap[0]))}

	heightmap = append(nineRow, heightmap...)
	heightmap = append(heightmap, make([]int, len(heightmap[0])))

	for j := 0; j < len(heightmap[0]); j++ {
		heightmap[0][j] = 9
		heightmap[len(heightmap)-1][j] = 9
	}

	return heightmap, nil
}
