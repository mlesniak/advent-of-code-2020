package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readNumbers(filename string) []int {
	bs, err := ioutil.ReadFile(filename)
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
	return numbers
}

func readLines(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	return lines
}
