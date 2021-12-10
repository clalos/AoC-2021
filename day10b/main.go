/**
https://adventofcode.com/2021/day/10
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var incompleteLines []string
	for scanner.Scan() {
		line := scanner.Text()
		var stack lifoStack
		corrupted := false
	charScanner:
		for _, char := range line {
			switch {
			case isOpen(char):
				newAllocation := char
				stack.Push(&newAllocation)
			case isClose(char):
				lastOpen := *stack.Pop()
				if char != openPairs[lastOpen] {
					corrupted = true
					break charScanner
				}
			default:
				log.Fatalf("unknown character found: %s", string(char))
			}
		}
		if !corrupted {
			incompleteLines = append(incompleteLines, line)
		}
	}

	var scores []int
	for _, line := range incompleteLines {
		var stack lifoStack
		for _, char := range line {
			switch {
			case isOpen(char):
				newAllocation := char
				stack.Push(&newAllocation)
			case isClose(char):
				stack.Pop()
			default:
				log.Fatalf("unknown character found: %s", string(char))
			}
		}

		fmt.Print(fmt.Sprintf("%s - Complete by adding ", line))
		var score int
		missing := stack.Unwrap()
		for i := len(missing) - 1; i >= 0; i-- {
			fmt.Print(string(openPairs[*missing[i]]))
			switch openPairs[*missing[i]] {
			case ')':
				score = score*5 + 1
			case ']':
				score = score*5 + 2
			case '}':
				score = score*5 + 3
			case '>':
				score = score*5 + 4
			}
		}
		scores = append(scores, score)
		fmt.Println()
	}

	sort.Ints(scores)
	fmt.Println(fmt.Sprintf("Score %d", scores[len(scores)/2]))
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

// lifoStack is a LIFO stack.
type lifoStack []*rune

// Push adds a node to the stack.
func (l *lifoStack) Push(n *rune) {
	*l = append(*l, n)
}

// Pop removes and returns a node from the stack in last to first order.
func (l *lifoStack) Pop() (n *rune) {
	x := l.Len() - 1
	n = (*l)[x]
	*l = (*l)[:x]
	return
}

// Last returns the last element in the stack. It does not pop.
func (l *lifoStack) Last() (n *rune) {
	x := l.Len() - 1
	n = (*l)[x]
	return
}

// Len returns the stack's size.
func (l *lifoStack) Len() int {
	return len(*l)
}

// Unwrap returns the underlying native type.
func (l *lifoStack) Unwrap() []*rune {
	return *l
}
