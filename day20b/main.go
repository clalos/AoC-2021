/**
https://adventofcode.com/2021/day/20
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	// Load Image Enhancement Algorithm.
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	row := s.Text()
	iea := make([]rune, 0, len(row))
	for _, char := range row {
		iea = append(iea, char)
	}
	// Skip empty line.
	s.Scan()

	// Load input image.
	var image [][]bool
	for s.Scan() {
		row := s.Text()
		imageRow := make([]bool, 0, len(row))
		for _, char := range row {
			imageRow = append(imageRow, char == '#')
		}
		image = append(image, imageRow)
	}

	// Apply IEA
	countPixels := 0
	bg := false
	bgFlashes := iea[0] == '#' && iea[len(iea)-1] == '.'
	for times := 0; times < 50; times++ {
		image = addFrame(image, 1, bg)
		countPixels = 0
		output := initialiseSlice(len(image))
		for x := 0; x < len(image); x++ {
			for y := 0; y < len(image[0]); y++ {
				b := bitset.New(64)
				if x+1 < len(image) {
					if y+1 < len(image[0]) {
						b.SetTo(0, image[x+1][y+1])
					} else {
						b.SetTo(0, bg)
					}

					b.SetTo(1, image[x+1][y])

					if y-1 >= 0 {
						b.SetTo(2, image[x+1][y-1])
					} else {
						b.SetTo(2, bg)
					}
				} else {
					b.SetTo(0, bg)
					b.SetTo(1, bg)
					b.SetTo(2, bg)
				}

				if y+1 < len(image[0]) {
					b.SetTo(3, image[x][y+1])
				} else {
					b.SetTo(3, bg)
				}
				b.SetTo(4, image[x][y])
				if y-1 >= 0 {
					b.SetTo(5, image[x][y-1])
				} else {
					b.SetTo(5, bg)
				}

				if x-1 >= 0 {
					if y+1 < len(image[0]) {
						b.SetTo(6, image[x-1][y+1])
					} else {
						b.SetTo(6, bg)
					}

					b.SetTo(7, image[x-1][y])
					if y-1 >= 0 {
						b.SetTo(8, image[x-1][y-1])
					} else {
						b.SetTo(8, bg)
					}

				} else {
					b.SetTo(6, bg)
					b.SetTo(7, bg)
					b.SetTo(8, bg)
				}

				if iea[b.Bytes()[0]] == '#' {
					countPixels++
					output[x][y] = true
				}
			}
		}
		image = output
		if bgFlashes {
			bg = !bg
		}
	}

	fmt.Println("Pixels lit: ", countPixels)
}

// initialiseSlice preallocate a slice of the given size
func initialiseSlice(size int) [][]bool {
	a := make([][]bool, size)
	for i := range a {
		a[i] = make([]bool, size)
	}
	return a
}

// addFrame adds a frame around the m matrix with the given defaultValue.
func addFrame(m [][]bool, size int, defaultValue bool) [][]bool {
	output := initialiseSlice(size*2 + len(m))
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			output[x+size][y+size] = m[x][y]
		}
	}

	for x := 0; x < len(output); x++ {
		for y := 0; y < len(output[0]); y++ {
			if x < size || x > size+len(m)-1 || y < size || y > size+len(m[0])-1 {
				output[x][y] = defaultValue
			}
		}
	}

	return output
}
