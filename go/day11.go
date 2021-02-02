package main

import "fmt"

func day11() {
	grid := readGrid("input/11.txt")
	displayGrid(grid)
	println("---------------------------")

	numChanges := 1
	cycles := 0
	for numChanges > 0 {
		cycles++
		grid, numChanges = computeCycle(grid)
		displayGrid(grid)
		println("---------------------------")
	}
	println(cycles)

	seatCount := 0
	for _, row := range grid {
		for _, col := range row {
			if col == '#' {
				seatCount++
			}
		}
	}
	println(seatCount)
}

func computeCycle(grid [][]byte) ([][]byte, int) {
	numChanges := 0
	tmp := make([][]byte, len(grid))

	for y := range grid {
		tmp[y] = make([]byte, len(grid[y]))
		for x := range grid[y] {
			seat := grid[y][x]
			if seat == '.' {
				tmp[y][x] = '.'
				continue
			}

			//occupiedSeats := computeOccupiedNeighbourSeats(grid, y, x)
			occupiedSeats := computeVisibleOccupiedSeats(grid, y, x)
			if seat == 'L' && occupiedSeats == 0 {
				tmp[y][x] = '#'
				numChanges++
				continue
			}

			if seat == '#' && occupiedSeats >= 5 {
				tmp[y][x] = 'L'
				numChanges++
				continue
			}

			tmp[y][x] = seat
		}
	}

	return tmp, numChanges
}

func computeVisibleOccupiedSeats(grid [][]byte, y int, x int) int {
	dxs := []int{-1, 0, +1}
	dys := []int{-1, 0, +1}

	count := 0
	for _, dx := range dxs {
		for _, dy := range dys {
			if dx == 0 && dy == 0 {
				continue
			}
			if isSeatOccupiedInLineOfSight(grid, y, x, dy, dx) {
				count++
			}
		}
	}

	return count
}

func isSeatOccupiedInLineOfSight(grid [][]byte, y int, x int, dy int, dx int) bool {
	width := len(grid[0])
	height := len(grid)

	for {
		y = y + dy
		x = x + dx

		// Still valid?
		if !(x >= 0 && x <= width-1 && y >= 0 && y <= height-1) {
			return false
		}

		// Check value at position
		if grid[y][x] == '.' {
			continue
		}
		if grid[y][x] == 'L' {
			return false
		}
		if grid[y][x] == '#' {
			return true
		}
	}
}

func computeOccupiedNeighbourSeats(grid [][]byte, y int, x int) int {
	count := 0

	width := len(grid[0])
	height := len(grid)

	if y >= 1 && grid[y-1][x] == '#' {
		count++
	}
	if y >= 1 && x <= width-2 && grid[y-1][x+1] == '#' {
		count++
	}
	if x <= width-2 && grid[y][x+1] == '#' {
		count++
	}
	if y <= height-2 && x <= width-2 && grid[y+1][x+1] == '#' {
		count++
	}
	if y <= height-2 && grid[y+1][x] == '#' {
		count++
	}
	if y <= height-2 && x >= 1 && grid[y+1][x-1] == '#' {
		count++
	}
	if x >= 1 && grid[y][x-1] == '#' {
		count++
	}
	if y >= 1 && x >= 1 && grid[y-1][x-1] == '#' {
		count++
	}

	return count
}

func readGrid(filename string) [][]byte {
	lines := readLines(filename)
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func displayGrid(grid [][]byte) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		println()
	}
}
