package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	grid  *Grid
	id    int
	sides []*tile
}

func (t *tile) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*** %v\n", t.id))
	for i := 0; i <= 3; i++ {
		id := "</>"
		if t.sides[i] != nil {
			id = fmt.Sprintf("%d", t.sides[i].id)
		}
		sb.WriteString(fmt.Sprintf("%v -> %v\n", i, id))
	}

	return sb.String()
}

func day20() {
	tiles := readTiles()

	for tileIndex := range tiles {
		tile := &tiles[tileIndex]

		for candIndex := range tiles {
			if candIndex == tileIndex {
				continue
			}
			candidate := &tiles[candIndex]

		nextSide:
			for candSide := 0; candSide <= 3; candSide++ {
				if candidate.sides[candSide] != nil {
					continue
				}

				// Try all possibilities to match against the candidates side.
				for flip := 0; flip <= 1; flip++ {
					for rotate := 0; rotate <= 3; rotate++ {
						if matchTile(candidate, tile, candSide) {
							linkTile(tile, candidate, (candSide+2)%4)
							continue nextSide
						}
						tile.rotate()
					}
					tile.flip()
				}
			}
		}
	}

	//analysis(tiles)

	corners := computeProduct(tiles)
	println(corners)
}

func analysis(tiles []tile) {
	println(len(tiles))

	howManyPointToMe := make(map[int]int)
	for _, tile := range tiles {
		for _, side := range tile.sides {
			if side == nil {
				continue
			}
			howManyPointToMe[side.id]++
		}
	}

	perSide := make(map[int]int)
	for k, v := range howManyPointToMe {
		fmt.Printf("id=%d <- #%d\n", k, v)
		perSide[v]++
	}
	println()
	for k, v := range perSide {
		fmt.Printf("sides=%d <- #%d\n", k, v)
	}
}

func fixed(candidate *tile) bool {
	for _, side := range candidate.sides {
		if side != nil {
			return true
		}
	}

	return false
}

func tileID(tile *tile) string {
	if tile == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%d", tile.id)
}

func computeProduct(tiles []tile) []tile {
	var prod int64 = 1

	var cornerTiles []tile

	for _, tile := range tiles {
		count := 0
		for _, side := range tile.sides {
			if side != nil {
				count++
			}
		}
		if count == 2 {
			//fmt.Printf("*** %v\n", tile.id)
			//for i := 0; i <= 3; i++ {
			//	id := "</>"
			//	if tile.sides[i] != nil {
			//		id = fmt.Sprintf("%d", tile.sides[i].id)
			//	}
			//	fmt.Printf("%v -> %v\n", i, id)
			//}
			prod *= int64(tile.id)
			cornerTiles = append(cornerTiles, tile)
		}
	}

	println(prod)
	return cornerTiles
}

func linkTile(t *tile, candidate *tile, tileSide int) {
	switch tileSide {
	case 0:
		if t.sides[0] == nil {
			t.sides[0] = candidate
		}
		if candidate.sides[2] == nil {
			candidate.sides[2] = t
		}
	case 2:
		if t.sides[2] == nil {
			t.sides[2] = candidate
		}
		if candidate.sides[0] == nil {
			candidate.sides[0] = t
		}
	case 3:
		if t.sides[3] == nil {
			t.sides[3] = candidate
		}
		if candidate.sides[1] == nil {
			candidate.sides[1] = t
		}
	case 1:
		if t.sides[1] == nil {
			t.sides[1] = candidate
		}
		if candidate.sides[3] == nil {
			candidate.sides[3] = t
		}
	}
}

func matchTile(t *tile, candidate *tile, orientation int) bool {
	// Find side for tile
	// Find side for candidate
	// Compare sides

	var sideTile []byte
	var sideCandidate []byte

	switch orientation {
	case 0:
		sideTile = t.grid.Data[0]
		sideCandidate = candidate.grid.Data[candidate.grid.Height-1]
	case 2:
		sideTile = t.grid.Data[t.grid.Height-1]
		sideCandidate = candidate.grid.Data[0]
	case 3:
		sideTile = make([]byte, t.grid.Height)
		for i := 0; i < t.grid.Height; i++ {
			sideTile[i] = t.grid.Data[i][0]
		}
		sideCandidate = make([]byte, candidate.grid.Height)
		for i := 0; i < candidate.grid.Height; i++ {
			sideCandidate[i] = candidate.grid.Data[i][candidate.grid.Width-1]
		}
	case 1:
		sideTile = make([]byte, t.grid.Height)
		for i := 0; i < t.grid.Height; i++ {
			sideTile[i] = t.grid.Data[i][t.grid.Width-1]
		}
		sideCandidate = make([]byte, candidate.grid.Height)
		for i := 0; i < candidate.grid.Height; i++ {
			sideCandidate[i] = candidate.grid.Data[i][0]
		}
	}

	return bytes.Compare(sideTile, sideCandidate) == 0
}

// flip flips a tile horizontally
func (t *tile) flip() {
	var ng Grid
	ng.Height = t.grid.Width
	ng.Width = t.grid.Height
	ng.Data = make([][]byte, t.grid.Height)
	for row := 0; row < t.grid.Height; row++ {
		ng.Data[row] = make([]byte, t.grid.Width)
	}

	for row := 0; row < t.grid.Height; row++ {
		for col := 0; col < t.grid.Width; col++ {
			ng.Data[t.grid.Height-row-1][col] = t.grid.Data[row][col]
		}
	}
	t.grid = &ng

	tmp := t.sides[0]
	t.sides[0] = t.sides[2]
	t.sides[2] = tmp
}

//func (g *Grid) flipVertical() *Grid {
//	// Format row / col.
//	var ng Grid
//	ng.Height = g.Width
//	ng.Width = g.Height
//	ng.Data = make([][]byte, g.Height)
//	for row := 0; row < g.Height; row++ {
//		ng.Data[row] = make([]byte, g.Width)
//	}
//
//	for row := 0; row < g.Height; row++ {
//		for col := 0; col < g.Width; col++ {
//			ng.Data[row][g.Width-col-1] = g.Data[row][col]
//		}
//	}
//
//	return &ng
//}

func (t *tile) rotate() {
	var ng Grid
	ng.Height = t.grid.Width
	ng.Width = t.grid.Height
	ng.Data = make([][]byte, t.grid.Height)
	for row := 0; row < t.grid.Height; row++ {
		ng.Data[row] = make([]byte, t.grid.Width)
	}
	for row := 0; row < t.grid.Height; row++ {
		for col := 0; col < t.grid.Width; col++ {
			ng.Data[col][t.grid.Width-row-1] = t.grid.Data[row][col]
		}
	}
	t.grid = &ng

	sides := make([]*tile, 4)
	sides[0] = t.sides[3]
	sides[1] = t.sides[0]
	sides[2] = t.sides[1]
	sides[3] = t.sides[2]
	t.sides = sides
}

func readTiles() []tile {
	var tiles []tile
	groups := readGroupedLines("input/20.txt")

	for _, group := range groups {
		ns := strings.Split(group[0], " ")[1]
		id, _ := strconv.Atoi(ns[:len(ns)-1])
		grid := parseGrid(group[1:])

		tiles = append(tiles, tile{
			grid:  &grid,
			id:    id,
			sides: make([]*tile, 4),
		})
	}

	return tiles
}

// 1171 2473 3079
// 1489      2311
// 2971 2729 1951

// 2729 2971
//      1489

// 1171 2473
// 1489

// 2311
// 1951 2729

//      2311
// 2473 3079
