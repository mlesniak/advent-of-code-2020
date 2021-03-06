package main

import (
	"fmt"
	"strings"
)

type coordinate struct {
	x, y, z, w int
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
	w := 0
	for y := 0; y < initialCells.Height; y++ {
		for x := 0; x < initialCells.Height; x++ {
			if initialCells.Data[y][x] == '#' {
				grid[coordinate{x, y, z, w}] = true
			}
		}
	}

	widthX := len(initialCells.Data[0])
	widthY := len(initialCells.Data)
	widthZ := 1
	widthW := 1

	for cycle := 0; cycle < 6; cycle++ {
		// Compute a single cycle
		tmp := make(Grid3D)
		for x := -widthX; x <= widthX; x++ {
			for y := -widthY; y <= widthY; y++ {
				for z := -widthZ; z <= widthZ; z++ {
					for w := -widthW; w <= widthW; w++ {
						// Compute number of neighbours
						activeNeighbours := 0
						for dx := -1; dx <= 1; dx++ {
							for dy := -1; dy <= 1; dy++ {
								for dz := -1; dz <= 1; dz++ {
									for dw := -1; dw <= 1; dw++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											continue
										}
										x2 := x + dx
										y2 := y + dy
										z2 := z + dz
										w2 := w + dw
										if grid[coordinate{x2, y2, z2, w2}] {
											activeNeighbours++
										}
									}
								}
							}
						}

						//fmt.Printf("%v %v %v / %v -> %d\n", x, y, z,
						//	grid[coordinate{x, y, z, w}],
						//	activeNeighbours)
						if grid[coordinate{
							x: x,
							y: y,
							z: z,
							w: w,
						}] && (activeNeighbours == 2 || activeNeighbours == 3) {
							tmp[coordinate{
								x: x,
								y: y,
								z: z,
								w: w,
							}] = true
						}
						if !grid[coordinate{
							x: x,
							y: y,
							z: z,
							w: w,
						}] && activeNeighbours == 3 {
							tmp[coordinate{
								x: x,
								y: y,
								z: z,
								w: w,
							}] = true
						}
					}
				}
			}
		}

		widthX += 1
		widthY += 1
		widthZ += 1
		widthW += 1
		grid = tmp
	}

	println(len(grid))
}
