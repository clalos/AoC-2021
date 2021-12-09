/**
https://adventofcode.com/2021/day/9
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// point represents a value in the matrix.
type point struct {
	value   int
	touched bool
}

func main() {
	heightmap, err := load()
	if err != nil {
		log.Fatalf("failed to read the input: %v", err)
	}

	var basinSizes []int
	for x := 1; x < len(heightmap)-1; x++ {
		for y := 1; y < len(heightmap[x])-1; y++ {
			curValue := heightmap[x][y].value

			// Look for the lowest point.
			if curValue < heightmap[x-1][y].value &&
				curValue < heightmap[x+1][y].value &&
				curValue < heightmap[x][y-1].value &&
				curValue < heightmap[x][y+1].value {

				// Calculate the basin size starting from this lowest point.
				basinSizes = append(basinSizes, basinSize(heightmap, x, y))
			}
		}
	}

	// Sort the basins and pick the last 3.
	sort.Ints(basinSizes)
	fmt.Println(fmt.Sprintf("Result: %d", basinSizes[len(basinSizes)-3]*basinSizes[len(basinSizes)-2]*basinSizes[len(basinSizes)-1]))
}

// basinSize calculates the size of the basin
// from the given lowest point coordinates
func basinSize(heightmap [][]point, x int, y int) int {
	if heightmap[x][y].touched || heightmap[x][y].value == 9 {
		return 0
	}
	heightmap[x][y].touched = true
	return 1 +
		basinSize(heightmap, x-1, y) +
		basinSize(heightmap, x+1, y) +
		basinSize(heightmap, x, y-1) +
		basinSize(heightmap, x, y+1)
}

// load reads the matrix from STDIN.
// It wraps the input matrix in a frame of 9 values.
func load() ([][]point, error) {
	var heightmap [][]point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var row []point
		for pos, stringRune := range scanner.Text() {
			value, err := strconv.Atoi(string(stringRune))
			if err != nil {
				return nil, fmt.Errorf("failed to conver the number %s", string(stringRune))
			}
			if pos == 0 {
				row = append(row, point{value: 9, touched: true})
			}
			row = append(row, point{value: value, touched: false})
		}
		row = append(row, point{value: 9, touched: true})
		heightmap = append(heightmap, row)
	}
	nineRow := [][]point{make([]point, len(heightmap[0]))}

	heightmap = append(nineRow, heightmap...)
	heightmap = append(heightmap, make([]point, len(heightmap[0])))

	for j := 0; j < len(heightmap[0]); j++ {
		heightmap[0][j] = point{value: 9, touched: false}
		heightmap[len(heightmap)-1][j] = point{value: 9, touched: false}
	}

	return heightmap, nil
}
