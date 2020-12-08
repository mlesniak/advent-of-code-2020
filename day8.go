package main

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Command  string
	Argument int
}

func day8() {
	lines := readLines("input/8.txt")
	code := parseCode(lines)

	for _, c := range code {
		fmt.Println(c)
	}

	seenIPs := make(map[int]struct{})

	// We do not know yet if this machine will be used in later
	// exercises, hence YAGNI.
	ip := 0
	acc := 0
	for {
		fmt.Printf("%02d %02d\n", ip, acc)
		if _, seen := seenIPs[ip]; seen {
			println(acc)
			break
		}
		seenIPs[ip] = struct{}{}

		switch c := code[ip]; {
		case c.Command == "nop":
			ip++
		case c.Command == "jmp":
			ip += c.Argument
		case c.Command == "acc":
			acc += c.Argument
			ip++
		}
		//readChar()
	}
}

func readChar() {
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func parseCode(lines []string) []Command {
	var code []Command
	for _, line := range lines {
		if strings.Trim(line, " \t") == "" {
			continue
		}

		var c Command
		fmt.Sscanf(line, "%s %d", &c.Command, &c.Argument)
		code = append(code, c)
	}

	return code
}
