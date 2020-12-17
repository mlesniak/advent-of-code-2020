package main

import (
	"fmt"
	"strings"
)

type coordinate struct {
	x, y, z int
}

type Grid3D map[coordinate]bool

func (g *Grid3D) String() string {
	var sb strings.Builder

	for k := range *g {
		sb.WriteString(fmt.Sprintf("%v\n", k))
	}

	return sb.String()
}

func day17() {
	initialCells := read2D("input/17.txt")

	grid := make(Grid3D)

	z := 0
	for y := 0; y < initialCells.Height; y++ {
		for x := 0; x < initialCells.Height; x++ {
			if initialCells.Data[y][x] == '#' {
				grid[coordinate{x, y, z}] = true
			}
		}
	}

	println(grid.String())
}
