/**
https://adventofcode.com/2021/day/4
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

// bingo holds a bingo game
type bingo struct {
	winningNumbers []int
	boards         []board
}

// square represents a board square
type square struct {
	x     int
	y     int
	value int
}

// board is a bingo board
type board struct {
	squares []square
}

func main() {
	// Load bingo from STDIN.
	var bingoGame bingo
	err := bingoGame.load()
	if err != nil {
		log.Fatal(err)
	}

	// Find winner board.
	var drawnNumbers []int
	var winningBoards []board
	totalBoards := len(bingoGame.boards)
	for _, drawn := range bingoGame.winningNumbers {
		drawnNumbers = append(drawnNumbers, drawn)
		for i := 0; i < len(bingoGame.boards); i++ {
			if bingoGame.boards[i].isWinning(drawnNumbers) {
				winningBoards = append(winningBoards, bingoGame.boards[i])

				// Remove itself from the boards in-game
				bingoGame.boards = append(bingoGame.boards[:i], bingoGame.boards[i+1:]...)
				i--
			}

			// As soon as all boards are winning, calculate the score for the last one.
			if len(winningBoards) == totalBoards {
				score := winningBoards[len(winningBoards)-1].getScore(drawnNumbers)
				fmt.Println(fmt.Sprintf("Score: %d", score))
				return
			}
		}

	}

}

// load reads the bingo game from stdin.
func (b *bingo) load() error {
	scanner := bufio.NewScanner(os.Stdin)
	var x, y, i int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// First line are drawn numbers
		if i == 0 {
			drawnNumbers := strings.Split(line, ",")
			for _, stringNum := range drawnNumbers {
				value, err := strconv.Atoi(stringNum)
				if err != nil {
					return fmt.Errorf("failed to convert the drawn number: %s", stringNum)
				}
				b.winningNumbers = append(b.winningNumbers, value)
			}
			i++
			continue
		}

		// Empty line.
		// Starts a new board and reset the indexes.
		if line == "" {
			x = 0
			y = 0
			b.boards = append(b.boards, board{})
			continue
		}

		// Board line.
		boardLine := strings.Split(line, " ")
		for _, stringNum := range boardLine {
			if strings.TrimSpace(stringNum) == "" {
				continue
			}
			value, err := strconv.Atoi(stringNum)
			if err != nil {
				return fmt.Errorf("failed to convert the board number: %s", stringNum)
			}

			// Append squares to the last board.
			b.boards[len(b.boards)-1].squares = append(b.boards[len(b.boards)-1].squares, square{x, y, value})
			y++
		}

		// Increment x and reset y
		x++
		y = 0
	}
	return nil
}

// isWinning returns true if the board is winning.
func (b *board) isWinning(winningNumbers []int) bool {
	// Get the winning squares.
	var markedSquares []square
	for _, square := range b.squares {
		if contains(winningNumbers, square.value) {
			markedSquares = append(markedSquares, square)
		}
	}

	// Check for 5 repeated x or y occurrences.
	xOccurrences := map[int]int{}
	yOccurrences := map[int]int{}
	for _, square := range markedSquares {
		xOccurrences[square.x] = xOccurrences[square.x] + 1
		yOccurrences[square.y] = yOccurrences[square.y] + 1
	}

	for _, occ := range xOccurrences {
		if occ == 5 {
			return true
		}
	}

	for _, occ := range yOccurrences {
		if occ == 5 {
			return true
		}
	}

	return false
}

// getScore calculates the score.
// Sum of the losing numbers multiplied the last drawn number.
func (b *board) getScore(winningNumbers []int) int {
	// Get losing squares.
	var unmarkedSquare []square
	for _, square := range b.squares {
		if !contains(winningNumbers, square.value) {
			unmarkedSquare = append(unmarkedSquare, square)
		}
	}

	// Sum unmarked squares.
	var sumUnmarked int
	for _, unmarked := range unmarkedSquare {
		sumUnmarked += unmarked.value
	}

	return sumUnmarked * winningNumbers[len(winningNumbers)-1]
}

// contains returns true if "search" is present in "n".
func contains(n []int, search int) bool {
	for i := 0; i < len(n); i++ {
		if n[i] == search {
			return true
		}
	}
	return false
}
