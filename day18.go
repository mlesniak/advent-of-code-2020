package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day18() {
	expressions := readLines("input/18.txt")

	for _, expr := range expressions {
		fmt.Printf("%v = %d\n", expr, eval(expr))
	}
}

func eval(expr string) int {
	parts := strings.Split(expr, " ")

	result := 0
	op := "+"

	for i := 0; i < len(parts); i += 2 {
		n2 := toNum(parts[i])
		result = compute(result, op, n2)
		if i != len(parts)-1 {
			op = parts[i+1]
		}
	}

	return result
}

func compute(n1 int, op string, n2 int) int {
	switch op {
	case "+":
		return n1 + n2
	case "*":
		return n1 * n2
	default:
		panic("not supported")
	}
}

func toNum(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
