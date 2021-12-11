/**
https://adventofcode.com/2021/day/11
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// octopus holds the octopus properties.
type octopus struct {
	energy  int
	flashed bool
}

func main() {
	grid, err := load()
	if err != nil {
		log.Fatal(err)
	}

	for step := 0; step < 500; step++ {
		var flashes int
		increaseEnergy(grid)
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y].energy--
				flashes += flash(grid, x, y)
			}
		}
		if flashes == 100 {
			fmt.Println(fmt.Sprintf("All flash at step: %d", step+1))
			return
		}
		reset(grid)
	}
}

// flash propagates a flash from the given position. It updates the input grid.
// flash returns the number of flashes occurred.
func flash(grid [][]octopus, x int, y int) int {
	if x < 0 || x > 9 || y < 0 || y > 9 {
		// Out of grid case.
		return 0
	}

	grid[x][y].energy++
	if grid[x][y].energy > 9 && !grid[x][y].flashed {
		grid[x][y].flashed = true
		return 1 +
			flash(grid, x-1, y) +
			flash(grid, x+1, y) +
			flash(grid, x, y-1) +
			flash(grid, x, y+1) +
			flash(grid, x+1, y+1) +
			flash(grid, x-1, y-1) +
			flash(grid, x+1, y-1) +
			flash(grid, x-1, y+1)
	}

	return 0
}

// increaseEnergy increases all energies by 1.
func increaseEnergy(grid [][]octopus) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			grid[x][y].energy++
		}
	}
}

// Reset resets all flashed octopuses to energy 0.
func reset(grid [][]octopus) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y].flashed {
				grid[x][y].energy = 0
				grid[x][y].flashed = false
			}
		}
	}
}

// load reads the input grid from STDIN.
func load() ([][]octopus, error) {
	var grid [][]octopus
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var row []octopus
		for _, num := range line {
			value, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, fmt.Errorf("failed to convert the rune %s", string(num))
			}
			row = append(row, octopus{energy: value})
		}
		grid = append(grid, row)
	}
	return grid, nil
}

