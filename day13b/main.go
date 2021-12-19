/**
https://adventofcode.com/2021/day/13
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type foldingType int

const (
	FoldUp foldingType = iota + 1
	FoldLeft
)

type coord struct {
	y, x int
}

type instruction struct {
	direction   foldingType
	foldingLine int
}

func main() {
	dots := map[coord]struct{}{}
	var instructions []instruction

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			instr := strings.Split(strings.Split(line, " ")[2], "=")
			value, _ := strconv.Atoi(instr[1])
			switch instr[0] {
			case "y":
				instructions = append(instructions, instruction{direction: FoldUp, foldingLine: value})
			case "x":
				instructions = append(instructions, instruction{direction: FoldLeft, foldingLine: value})
			}
			continue
		}

		stringCoord := strings.Split(line, ",")
		x, _ := strconv.Atoi(stringCoord[0])
		y, _ := strconv.Atoi(stringCoord[1])
		dots[coord{y, x}] = struct{}{}
	}

	for _, fold := range instructions {
		switch fold.direction {
		case FoldUp:
			for index := range dots {
				if index.y > fold.foldingLine {
					delete(dots, index)
					dots[foldY(fold.foldingLine, index)] = struct{}{}
				}
			}
		case FoldLeft:
			for index := range dots {
				if index.x > fold.foldingLine {
					delete(dots, index)
					dots[foldX(fold.foldingLine, index)] = struct{}{}
				}
			}
		}
	}

	// Print result
	for y := 0; y < 6; y++ {
		for x := 0; x < 45; x++ {
			if _, found := dots[coord{y, x}]; found {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

// foldY returns the c coordinate folded up.
func foldY(foldingLine int, c coord) coord {
	return coord{foldingLine - (c.y - foldingLine), c.x}
}

// foldX returns the c coordinate folded left.
func foldX(foldingLine int, c coord) coord {
	return coord{c.y, foldingLine - (c.x - foldingLine)}
}
