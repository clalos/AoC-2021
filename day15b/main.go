/**
https://adventofcode.com/2021/day/15
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	levels, err := load()
	if err != nil {
		log.Fatalf("failed to read the input: %v", err)
	}

	rowCount := len(levels)
	columnCount := len(levels[0])

	graph := dijkstra.NewGraph()
	for x := 0; x < rowCount; x++ {
		for y := 0; y < columnCount; y++ {

			vertexID := coordinatesToID(columnCount, x, y)
			if _, err := graph.GetVertex(vertexID); err != nil {
				graph.AddVertex(vertexID)
			}

			if x-1 >= 0 {
				createArc(graph, vertexID, coordinatesToID(columnCount, x-1, y), levels[x-1][y])
			}
			if y-1 >= 0 {
				createArc(graph, vertexID, coordinatesToID(columnCount, x, y-1), levels[x][y-1])
			}
			if x+1 < rowCount {
				createArc(graph, vertexID, coordinatesToID(columnCount, x+1, y), levels[x+1][y])
			}
			if y+1 < columnCount {
				createArc(graph, vertexID, coordinatesToID(columnCount, x, y+1), levels[x][y+1])
			}
		}
	}

	best, err := graph.Shortest(coordinatesToID(columnCount, 0, 0), coordinatesToID(columnCount, rowCount-1, columnCount-1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lowest total risk: ", best.Distance)
}

// coordinatesToID returns a unique ID for the given coordinates.
func coordinatesToID(totalColumns, x, y int) int {
	return totalColumns*x + y
}

// createArc creates a new arch for the given vertices IDs.
func createArc(graph *dijkstra.Graph, fromID int, toID, weight int) {
	if _, err := graph.GetVertex(toID); err != nil {
		graph.AddVertex(toID)
	}
	err := graph.AddArc(fromID, toID, int64(weight))
	if err != nil {
		log.Fatal(err)
	}
}

// load reads the input from STDIN.
func load() ([][]int, error) {
	var levels [][]int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		var row []int
		for i := 0; i < len(line); i++ {
			value, err := strconv.Atoi(string(line[i]))
			if err != nil {
				return nil, err
			}
			row = append(row, value)
		}
		levels = append(levels, row)
	}

	var output [][]int
	for incr := 0; incr < 5; incr++ {
		for x := 0; x < len(levels); x++ {
			var row []int
			for columnIncr := 0; columnIncr < 5; columnIncr++ {
				for y := 0; y < len(levels[0]); y++ {
					weight := levels[x][y] + columnIncr + incr
					if weight > 9 {
						weight = weight - 9
					}
					row = append(row, weight)
				}
			}
			output = append(output, row)
		}
	}

	return output, nil
}
