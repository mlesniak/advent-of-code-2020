package main

import "fmt"

func day3() {
	grid := read2D("input/3.txt")
	trees := 0

	// Slope
	dx := 3
	dy := 1

	x := 0
	y := 0
	for y < grid.Height {
		objectAt := grid.Data[y][x%grid.Width]
		if objectAt == '#' {
			trees++
		}

		x += dx
		y += dy
	}

	fmt.Printf("Trees: %v\n", trees)
}
