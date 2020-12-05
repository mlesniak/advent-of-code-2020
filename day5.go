package main

import (
	"strconv"
	"strings"
)

type seat = int

func day5() {
	lines := readLines("input/5.txt")
	seats := parseSeats(lines)

	maxSeat := 0
	for _, s := range seats {
		if s > maxSeat {
			maxSeat = s
		}
	}
	println(maxSeat)
}

func parseSeats(lines []string) []seat {
	var seats []seat

	for _, line := range lines {
		s1 := strings.ReplaceAll(line, "F", "0")
		s2 := strings.ReplaceAll(s1, "B", "1")
		s3 := strings.ReplaceAll(s2, "R", "1")
		l := strings.ReplaceAll(s3, "L", "0")

		row := l[:7]
		col := l[7:]

		ri, err := strconv.ParseInt(row, 2, 64)
		if err != nil {
			panic(err)
		}
		ci, err := strconv.ParseInt(col, 2, 64)

		seatID := ri*8 + ci
		seats = append(seats, seat(seatID))
	}

	return seats
}
