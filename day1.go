package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input-1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs), "\n")
	var numbers []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1+n2 == 2020 {
				println(n1 * n2)
				return
			}
		}
	}
}
