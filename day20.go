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

	//analysis(tiles)  1471 1637 3877 3407 1721 1783 2377 2309 1753 2797 2971 2677

	corners := computeProduct(tiles)
	// dir2 := 3

	// Walk in first direction
	start := &corners[0]
	fmt.Printf("%v\n", start)
	dir1 := 0
	dir2 := 3

	//start = start.sides[3]
	//start.rotate()
	//start.rotate()

	ids := make([][]int, 12)
	for i := 0; i < 12; i++ {
		ids[i] = make([]int, 12)
	}
	y := 0
	x := 0

	for {
		cur := start
		for cur != nil {
			ids[y][x] = cur.id
			fmt.Printf("%v ", cur.id)
			next := cur.sides[dir1]
			if next == nil {
				break
			}
			// Determine next side to follow
			for i, sideNode := range next.sides {
				if sideNode == nil {
					continue
				}
				if sideNode.id == cur.id {
					dir1 = (i + 2) % 4
				}
			}
			cur = next
			x++
		}
		y++
		x = 0
		fmt.Println() // 1471 2801 1759 1193 2749 2153 3821 1951 1789 1439 2137

		tmp := start.sides[dir2]
		if tmp == nil {
			break
		}

		for {
			if tmp.sides[2] == nil {
				break
			}
			//fmt.Printf("Rotating %v\n", tmp.id)
			tmp.rotate()
		}
		for i, sideNode := range tmp.sides {
			if sideNode == nil {
				continue
			}
			if sideNode.id == start.id {
				dir2 = (i + 2) % 4
			}
		}
		dir1 = 0
		start = tmp
	}

	tileMap := make(map[int]*tile)
	for _, t := range tiles {
		tmp := t
		tileMap[t.id] = &tmp
		removeBorder(&t)
	}

	// Create large string from map. 96 x 96 grid

	var sb strings.Builder
	for row := 0; row < 12; row++ {
		for line := 0; line < 8; line++ {
			for col := 0; col < 12; col++ {
				t := tileMap[ids[row][col]]
				sb.Write(t.grid.Data[line])
			}
			sb.WriteString("\n")
		}
	}
	image := sb.String()
	lines := strings.Split(image, "\n")
	lines = lines[0 : len(lines)-1]
	g := parseGrid(lines)
	ti := tile{
		grid:  &g,
		id:    0,
		sides: make([]*tile, 4),
	}
	//ti.flip()
	//ti.rotate()
	//ti.rotate()
	//ti.rotate()
	fmt.Printf("%s\n", ti.grid.String())

	num := strings.Count(ti.grid.String(), "#")
	fmt.Printf("Number of #: %d\n", num)
	sm := 14
	count := 2
	fmt.Printf("%d\n", num-sm*count)

	//rx := regexp.MustCompile(`(?s)##`)
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
