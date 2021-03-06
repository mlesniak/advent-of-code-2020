package main

import (
	"fmt"
	"os"
	"strings"
)

type Code []Command

type Command struct {
	Command  string
	Argument int
}

const (
	ResultLoop = iota
	ResultTerminated
)

type Result int

func (r Result) String() string {
	switch r {
	case ResultLoop:
		return "LOOP"
	case ResultTerminated:
		return "FINISHED"
	default:
		return "<?>"
	}
}

func day8() {
	lines := readLines("input/8.txt")
	code := parseCode(lines)

	pos := 0
	for {
		buffer := make([]Command, len(code))
		copy(buffer, code)

		// We know the initial code is wrong.
		// Find next position:
		for {
			if buffer[pos].Command == "jmp" {
				buffer[pos].Command = "nop"
				pos++
				break
			}
			if buffer[pos].Command == "nop" {
				buffer[pos].Command = "jmp"
				pos++
				break
			}
			pos++
		}

		result, value := executeCode(buffer)
		fmt.Printf("%v -> %d\n", result, value)

		if result == ResultTerminated {
			break
		}
	}
}

func executeCode(code Code) (Result, int) {
	// We do not know yet if this machine will be used in later
	// exercises, hence YAGNI.
	seenIPs := make(map[int]struct{})
	ip := 0
	acc := 0
	for {
		fmt.Printf("%02d %02d\n", ip, acc)

		if ip >= len(code) {
			return ResultTerminated, acc
		}

		if _, seen := seenIPs[ip]; seen {
			return ResultLoop, acc
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

func parseCode(lines []string) Code {
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
