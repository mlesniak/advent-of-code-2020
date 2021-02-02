package main

import "fmt"

func day3() {
	grid := read2D("input/3.txt")

	// Slopes
	slopes := []struct {
		dx int
		dy int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, slope := range slopes {
		trees := computeSlope(grid, slope.dx, slope.dy)
		product *= trees
	}
	fmt.Printf("%v\n", product)
}

func computeSlope(grid Grid, dx int, dy int) int {
	x := 0
	y := 0
	count := 0

	for y < grid.Height {
		objectAt := grid.Data[y][x%grid.Width]
		if objectAt == '#' {
			count++
		}

		x += dx
		y += dy
	}

	return count
}
