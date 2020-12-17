package main

import "fmt"

type coordinate struct {
	x, y, z int
}

func day17() {
	initialCells := read2D("input/17.txt")

	grid := make(map[coordinate]bool)

	z := 0
	for y := 0; y < initialCells.Height; y++ {
		for x := 0; x < initialCells.Height; x++ {
			if initialCells.Data[y][x] == '#' {
				grid[coordinate{x, y, z}] = true
			}
		}
	}

	for k := range grid {
		fmt.Printf("%v\n", k)
	}
}
