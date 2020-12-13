package main

import (
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

	for t := 0; ; t++ {
		found := true
		for i := 0; i < len(busTimes); i++ {
			if busTimes[i] == -1 {
				// ignore, since 'x'
				continue
			}

			okTime := t + i
			if okTime%busTimes[i] != 0 {
				// Not viable
				found = false
				break
			}
		}

		if found {
			println(t)
			break
		}
	}
}
