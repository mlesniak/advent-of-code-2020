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
	return strings.Split(readFile(filename), "\n")
}

func readFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func readGroupedLines(filename string) [][]string {
	var result [][]string

	groups := strings.Split(readFile(filename), "\n\n")
	for _, group := range groups {
		result = append(result, strings.Split(group, "\n"))
	}

	return result
}

type Grid struct {
	Height int
	Width  int
	Data   [][]byte
}

// format: row, column
func read2D(filename string) Grid {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return Grid{
		Height: len(lines),
		Width:  len(lines[0]),
		Data:   grid,
	}
}

func linesToNumbers(numbers []string) []int {
	var res []int

	for _, number := range numbers {
		i, _ := strconv.Atoi(number)
		res = append(res, i)
	}

	return res
}
