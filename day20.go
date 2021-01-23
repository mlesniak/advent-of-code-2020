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

func orientationToString(o int) string {
	switch o {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	}

	panic("unknown orientation")
}

func (t *tile) String() string {
	//return fmt.Sprintf("%v\n%v\n", t.id, t.grid.String())
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

// 2971     1489     1171
// 2729              2473
//

// 1951    2311
// 2729

// 2971    1489
// 2729

// 2311    3079
//         2473

// 1951    2311    3079
// 2729    1427    2473
// 2971    1489    1171

func day20() {
	tiles := readTiles()

	for tileIndex := range tiles {
		tile := &tiles[tileIndex]
		//fmt.Printf("- TILE %v----------------------------------------------\n", tile.id)

		for candIndex := range tiles {
			if candIndex == tileIndex {
				continue
			}
			candidate := &tiles[candIndex]
			//fmt.Printf("? %v\n", candidate.id)

		nextSide:
			for candSide := 0; candSide <= 3; candSide++ {
				// Try all possibilities to match against the candidates side.
				for flip := 0; flip <= 1; flip++ {
					for rotate := 0; rotate <= 3; rotate++ {
						if matchTile(candidate, tile, candSide) {
							//fmt.Printf("! match candSide=%v to tileSide=%v\n", candSide, (candSide+2)%4)
							//fmt.Printf("tile=\n%v\ncand=%v\n", tile, candidate)
							linkTile(tile, candidate, (candSide+2)%4)
							continue nextSide
						}

						//fmt.Printf("R90\n")
						tile.rotate()
					}
					//fmt.Printf("FH\n")
					tile.flip()
				}
			}
		}

		//fmt.Printf("*** %v\n", tile.id)
		//for i := 0; i <= 3; i++ {
		//	id := "</>"
		//	if tile.sides[i] != nil {
		//		id = fmt.Sprintf("%d", tile.sides[i].id)
		//	}
		//	fmt.Printf("%v -> %v\n", i, id)
		//}
	}

	corners := computeProduct(tiles)

	// Remove borders.
	for i := range tiles {
		tile := &tiles[i]
		removeBorder(tile)
	}

	//fmt.Printf("%d\n%s\n", tiles[0].id, tiles[0].grid.String())
	//return

	// create combined image.
	seen := make(map[int]struct{})
	start := &corners[0]
	var dir1 int
	var dir2 int
	first := false
	for i := 0; i <= 3; i++ {
		if start.sides[i] != nil {
			if !first {
				dir1 = i
				first = true
				continue
			} else {
				dir2 = i
				break
			}
		}
	}

	seen[start.id] = struct{}{}

	for start != nil {
		head := start
		for head != nil {
			fmt.Printf("%v\n", head.id)
			if head.sides[dir1] != nil {
				_, found := seen[head.sides[dir1].id]
				if !found {
					head = head.sides[dir1]
					seen[head.id] = struct{}{}
					continue
				} else {
					head = head.sides[(dir1+2)%4]
					if head == nil {
						continue
					}
					seen[head.id] = struct{}{}
					continue
				}
			} else if head.sides[(dir1+2)%4] != nil {
				_, found := seen[head.sides[(dir1+2)%4].id]
				if !found {
					head = head.sides[(dir1+2)%4]
					seen[head.id] = struct{}{}
					continue
				} else {
					head = head.sides[dir1]
					if head == nil {
						continue
					}
					seen[head.id] = struct{}{}
					continue
				}
			}
		}

		if start.sides[dir2] != nil {
			_, found := seen[start.sides[dir2].id]
			if !found {
				start = start.sides[dir2]
				seen[start.id] = struct{}{}
				continue
			} else {
				start = start.sides[(dir2+2)%4]
				if start == nil {
					continue
				}
				continue
			}
		} else if start.sides[(dir2+2)%4] != nil {
			_, found := seen[start.sides[(dir2+2)%4].id]
			if !found {
				start = start.sides[(dir2+2)%4]
				seen[start.id] = struct{}{}
				continue
			} else {
				start = start.sides[dir2]
				if start == nil {
					continue
				}
				continue
			}
		}
	}

}

func removeBorder(t *tile) {
	t.grid.Data = t.grid.Data[1 : len(t.grid.Data)-1]
	for i := range t.grid.Data {
		t.grid.Data[i] = t.grid.Data[i][1 : len(t.grid.Data[i])-1]
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
		//fmt.Printf("%v = %v\n", tile.id, count)
		if count == 2 {
			fmt.Printf("*** %v\n", tile.id)
			for i := 0; i <= 3; i++ {
				id := "</>"
				if tile.sides[i] != nil {
					id = fmt.Sprintf("%d", tile.sides[i].id)
				}
				fmt.Printf("%v -> %v\n", i, id)
			}
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
