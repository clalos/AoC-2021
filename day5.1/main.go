/**
https://adventofcode.com/2021/day/5
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// segment represents a segment.
type segment struct {
	startX int
	startY int
	endX   int
	endY   int
}

// isDiagonal returns true if the segment is at 45 degrees.
func (s *segment) is45Degree() bool {
	xLength := s.startX - s.endX
	yLength := s.startY - s.endY
	return math.Abs(float64(xLength)) == math.Abs(float64(yLength))
}

// isDiagonal returns true if the segment is diagonal (non-vertical/horizontal)
func (s *segment) isDiagonal() bool {
	return s.startX != s.endX && s.startY != s.endY
}

// isHorizontal returns true if the segment is horizontal.
func (s *segment) isHorizontal() bool {
	return !s.isDiagonal() && s.startY == s.endY
}

// isVertical returns true if the segment is vertical.
func (s *segment) isVertical() bool {
	return !s.isDiagonal() && s.startX == s.endX
}

// main is this program entrypoint.
func main() {
	segments, err := loadSegments()
	if err != nil {
		log.Fatal(err)
	}

	diagram := initDiagram()
	for _, segment := range segments {

		if segment.is45Degree() {
			points := int(math.Abs(float64(segment.startX-segment.endX)) + 1)
			x := segment.startX
			y := segment.startY
			for i := 0; i < points; i++ {
				diagram[x][y]++
				if segment.startX < segment.endX {
					x++
				} else {
					x--
				}
				if segment.startY < segment.endY {
					y++
				} else {
					y--
				}
			}
		}

		if segment.isHorizontal() {
			// From right to left.
			if segment.startX > segment.endX {
				for i := segment.endX; i <= segment.startX; i++ {
					diagram[i][segment.startY]++
				}
				continue
			}

			// From left to write.
			for i := segment.startX; i <= segment.endX; i++ {
				diagram[i][segment.startY]++
			}
		}

		if segment.isVertical() {
			// From top to bottom.
			if segment.startY > segment.endY {
				for i := segment.endY; i <= segment.startY; i++ {
					diagram[segment.startX][i]++
				}
				continue
			}

			// From bottom to top
			for i := segment.startY; i <= segment.endY; i++ {
				diagram[segment.startX][i]++
			}
		}
	}

	fmt.Println(fmt.Sprintf("%d overlaps", countOverlaps(diagram)))
}

// loadSegments loads the segment from STDIN to memory.
func loadSegments() ([]segment, error) {
	var segments []segment
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())

		coordinates := strings.Split(row, "->")
		startCoordinates := strings.Split(strings.TrimSpace(coordinates[0]), ",")
		endCoordinates := strings.Split(strings.TrimSpace(coordinates[1]), ",")

		startX, err := strconv.Atoi(startCoordinates[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert start X coordinate: %s", startCoordinates[0])
		}
		startY, err := strconv.Atoi(startCoordinates[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert start Y coordinate: %s", startCoordinates[1])
		}
		endX, err := strconv.Atoi(endCoordinates[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert end X coordinate: %s", endCoordinates[0])
		}
		endY, err := strconv.Atoi(endCoordinates[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert end Y coordinate: %s", endCoordinates[1])
		}

		segments = append(segments, segment{
			startX: startX,
			startY: startY,
			endX:   endX,
			endY:   endY,
		})
	}

	return segments, nil
}

// countOverlaps counts the number of overlaps (value >= 2).
func countOverlaps(diagram [][]int) int {
	var overlaps int
	for x := 0; x < len(diagram); x++ {
		for y := 0; y < len(diagram[x]); y++ {
			if diagram[x][y] >= 2 {
				overlaps++
			}
		}
	}

	return overlaps
}

// initDiagram initialise a matrix with all zeros.
func initDiagram() [][]int {
	diagram := make([][]int, 1000)
	for i := range diagram {
		diagram[i] = make([]int, 1000)
	}
	return diagram
}
