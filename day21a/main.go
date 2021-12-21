package main

import (
	"container/ring"
	"fmt"
)

func main() {
	p1StartingPosition := 1
	p2StartingPosition := 5
	// Create the board.
	board1 := initRing(10)
	board2 := initRing(10)
	die := initRing(100)

	// Init scores and positions.
	var p1Score, p2Score, rolls int
	board1 = board1.Move(p1StartingPosition - 1)
	board2 = board2.Move(p2StartingPosition - 1)

	// Play game.
	turn := false
	for p1Score < 1000 && p2Score < 1000 {
		turn = !turn
		movement := die.Value.(int) + die.Next().Value.(int) + die.Next().Next().Value.(int)
		if turn {
			board1 = board1.Move(movement)
			p1Score += board1.Value.(int)
		} else {
			board2 = board2.Move(movement)
			p2Score += board2.Value.(int)
		}
		die = die.Move(3)
		rolls += 3
	}

	if p1Score >= 1000 {
		fmt.Println("Result:", p2Score*rolls)
		return
	}
	fmt.Println("Result:", p1Score*rolls)
}

// initRing initialises a ring of integers.
func initRing(size int) *ring.Ring {
	board := ring.New(size)
	for i := 1; i <= board.Len(); i++ {
		board.Value = i
		board = board.Next()
	}
	return board
}
