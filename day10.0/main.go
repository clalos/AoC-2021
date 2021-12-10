/**
https://adventofcode.com/2021/day/10
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	openPairs = map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
		'<': '>',
	}

	closePairs = map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
		'>': '<',
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var illegalChars []rune
	for scanner.Scan() {
		line := scanner.Text()
		stack := NewLifo()
	charScanner:
		for _, char := range line {
			switch {
			case isOpen(char):
				stack.Push(char)
			case isClose(char):
				// Check if the last opened parenthesis correspond to this close parenthesis
				lastOpened := stack.Last()
				if char != openPairs[lastOpened] {
					illegalChars = append(illegalChars, char)
					fmt.Println(fmt.Sprintf("%s - Expected %s, but found %s instead", line, string(openPairs[lastOpened]), string(char)))
					break charScanner
				}
				stack.Pop()
			default:
				log.Fatalf("unknown character found: %s", string(char))
			}
		}
	}

	// Calculate points
	var score int
	for _, c := range illegalChars {
		switch c {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}

	fmt.Println(fmt.Sprintf("Final score: %d", score))
}

// isOpen returns true if the given rune is an open parenthesis.
func isOpen(r rune) bool {
	_, found := openPairs[r]
	return found
}

// isClose returns true if the given rune is a closed parenthesis.
func isClose(r rune) bool {
	_, found := closePairs[r]
	return found
}

// NewLifo returns a new lifo stack.
func NewLifo() *Lifo {
	return &Lifo{}
}

// Lifo is a basic LIFO stack that resizes as needed.
type Lifo struct {
	nodes []rune
	count int
}

// Push adds a rune to the stack.
func (s *Lifo) Push(n rune) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Lifo) Pop() rune {
	if s.count == 0 {
		return rune(0)
	}
	s.count--
	return s.nodes[s.count]
}

// Last returns the last element in the stack, it doesn't pop.
func (s *Lifo) Last() rune {
	if s.count == 0 {
		return rune(0)
	}
	return s.nodes[s.count-1]
}
