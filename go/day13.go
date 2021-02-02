package main

import (
	modmath "github.com/deanveloper/modmath/v1/bigmod"
	"math/big"
	"strconv"
	"strings"
)

func day13() {
	timeplan := readLines("input/13.txt")

	var busTimes []int
	for _, busTime := range strings.Split(timeplan[1], ",") {
		if busTime == "x" {
			busTimes = append(busTimes, -1)
			continue
		}
		bt, _ := strconv.Atoi(busTime)
		busTimes = append(busTimes, bt)
	}

	var eqs []modmath.CrtEntry
	for i, t := range busTimes {
		if t == -1 {
			continue
		}
		eqs = append(eqs, modmath.CrtEntry{A: big.NewInt(int64(-i)), N: big.NewInt(int64(t))})
	}

	solution := modmath.SolveCrtMany(eqs)
	println(solution.Int64())
}
