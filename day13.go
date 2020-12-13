package main

import (
	"strconv"
	"strings"
)

func day13() {
	timeplan := readLines("input/13.txt")

	startTime, _ := strconv.Atoi(timeplan[0])
	var busTimes []int
	for _, busTime := range strings.Split(timeplan[1], ",") {
		if busTime == "x" {
			continue
		}
		bt, _ := strconv.Atoi(busTime)
		busTimes = append(busTimes, bt)
	}

	var solution int
	t := startTime
loop:
	for {
		for _, bt := range busTimes {
			if t%bt == 0 {
				solution = (t - startTime) * bt
				break loop
			}
		}
		t++
	}

	println(solution)
}
