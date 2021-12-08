/**
https://adventofcode.com/2021/day/8
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// entry represents one line of the input.
type entry struct {
	patterns []string
	output   []string
}

func main() {
	entries, err := load()
	if err != nil {
		log.Fatalf("unabled to read the input: %v", err)
	}

	var sum int
	for _, entry := range entries {
		digits := map[int]string{}

		// Associate unique numbers to their patterns.
		for _, pattern := range entry.patterns {
			switch len(pattern) {
			case 2:
				// Number 1
				digits[1] = pattern
			case 4:
				// Number 4
				digits[4] = pattern
			case 3:
				// Number 7
				digits[7] = pattern
			case 7:
				// Number 8
				digits[8] = pattern
			}
		}

		fiveSegments := removeCommonCharacters(digits[4], digits[1])
		for _, pattern := range entry.patterns {
			// Number 3
			if len(pattern) == 5 && containsCharacters(pattern, digits[1]) {
				digits[3] = pattern
			}
			// Number 5
			if len(pattern) == 5 && containsCharacters(pattern, fiveSegments) {
				digits[5] = pattern
			}
			// Number 6
			if len(pattern) == 6 && !containsCharacters(pattern, digits[1]) {
				digits[6] = pattern
			}
			// Number 9
			if len(pattern) == 6 && containsCharacters(pattern, digits[1]) && containsCharacters(pattern, fiveSegments) {
				digits[9] = pattern
			}
		}

		for _, pattern := range entry.patterns {
			// Number 0
			if len(pattern) == 6 && pattern != digits[6] && pattern != digits[9] {
				digits[0] = pattern
			}
			// Number 2
			if len(pattern) == 5 && pattern != digits[3] && pattern != digits[5] {
				digits[2] = pattern
			}
		}

		// Sort the patterns and the outputs alphabetically, so we can compare them.
		var outputString string
		for _, output := range entry.output {
			switch sortStringByCharacter(output) {
			case sortStringByCharacter(digits[0]):
				outputString += "0"
			case sortStringByCharacter(digits[1]):
				outputString += "1"
			case sortStringByCharacter(digits[2]):
				outputString += "2"
			case sortStringByCharacter(digits[3]):
				outputString += "3"
			case sortStringByCharacter(digits[4]):
				outputString += "4"
			case sortStringByCharacter(digits[5]):
				outputString += "5"
			case sortStringByCharacter(digits[6]):
				outputString += "6"
			case sortStringByCharacter(digits[7]):
				outputString += "7"
			case sortStringByCharacter(digits[8]):
				outputString += "8"
			case sortStringByCharacter(digits[9]):
				outputString += "9"
			default:
				log.Fatalf("unable to find the pattern: %s", output)
			}
		}

		outputNumber, err := strconv.Atoi(outputString)
		if err != nil {
			log.Fatalf("failed to convert the output string: %s", outputString)
		}
		sum += outputNumber
	}

	fmt.Println(fmt.Sprintf("Sum: %d", sum))
}

// load reads the input from STDIN.
func load() ([]entry, error) {
	var entries []entry
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var record entry
		row := strings.Split(strings.TrimSpace(scanner.Text()), "|")
		if len(row) != 2 {
			return nil, fmt.Errorf("unknown line: %v", scanner.Text())
		}
		// Patterns
		record.patterns = strings.Split(strings.TrimSpace(row[0]), " ")
		// Outputs
		record.output = strings.Split(strings.TrimSpace(row[1]), " ")

		entries = append(entries, record)
	}

	return entries, nil
}

// removeCommonCharacters removes the common characters.
func removeCommonCharacters(from, to string) string {
	var output string
	for _, i := range from {
		var found bool
		for _, j := range to {
			if string(i) == string(j) {
				found = true
			}
		}
		if !found {
			output += string(i)
		}
	}
	return output
}

// containsCharacters checks if "input" contains all the characters in "str"
func containsCharacters(input, str string) bool {
	contains := true
	for _, i := range str {
		if !strings.ContainsRune(input, i) {
			contains = false
		}
	}
	return contains
}

// ByRune represents a string as a slice of runes.
// ByRune implements the sort interface.
type ByRune []rune
func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

// stringToRuneSlice converts a string to a slice of runes.
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// sortStringByCharacter sorts a string by characters.
func sortStringByCharacter(s string) string {
	var r ByRune = stringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}