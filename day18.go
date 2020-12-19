package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day18() {
	expressions := readLines("input/18.txt")

	sum := 0
	for _, expr := range expressions {
		res := eval(expr)
		fmt.Printf("%v = %d\n", expr, res)
		sum += res
	}
	println(sum)
}

func eval(expr string) int {
	for strings.Contains(expr, "(") {
		start, end, value := evalSub(expr)
		var prefix string
		if start == 0 {
			prefix = ""
		} else {
			prefix = expr[:start-1] + " "
		}
		expr = prefix + fmt.Sprintf("%d", value) + expr[end+1:]
	}

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

func evalSub(expr string) (int, int, int) {
	// Find first bracket
	start := 0
	for i := range expr {
		if expr[i] == '(' {
			start = i
			break
		}
	}

	// Find matching bracket
	count := 0
	end := 0
	for j := start + 1; j < len(expr); j++ {
		if expr[j] == ')' && count == 0 {
			end = j
			break
		}
		if expr[j] == ')' {
			count--
		}
		if expr[j] == '(' {
			count++
		}
	}

	result := eval(expr[start+1 : end])
	return start, end, result
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
