package main

import "fmt"

func day11() {
	grid := readGrid("input/11.txt")

	for _, row := range grid {
		fmt.Printf("%v\n", row)
	}
}

func readGrid(filename string) [][]byte {
	lines := readLines(filename)
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}
