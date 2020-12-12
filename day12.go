package main

import (
	"fmt"
	"math"
	"strconv"
)

type ship struct {
	direction int
	x         int
	y         int
}

func (s ship) String() string {
	return fmt.Sprintf("x=%v y=%v dir=%v", s.x, s.y, s.direction)
}

type instruction struct {
	command  string
	argument int
}

func day12() {
	instructions := readLines("input/12.txt")

	ship := ship{
		direction: 0,
		x:         0,
		y:         0,
	}

	for _, instr := range instructions {
		fmt.Printf("\n%v\n", ship)
		fmt.Printf("%v\n", instr)
		i := parseInstruction(instr)
		switch i.command {
		case "N":
			ship.y += i.argument
		case "S":
			ship.y -= i.argument
		case "E":
			ship.x += i.argument
		case "W":
			ship.x -= i.argument
		case "L":
			ship.direction += i.argument
		case "R":
			ship.direction -= i.argument
		case "F":
			dx := int(math.Cos(float64(ship.direction) * math.Pi / 180))
			dy := int(math.Sin(float64(ship.direction) * math.Pi / 180))
			ship.x += dx * i.argument
			ship.y += dy * i.argument
		}
	}

	fmt.Printf("%v\n", ship)
	println(int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))))
}

func parseInstruction(instr string) instruction {
	c := instr[0]
	a, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic(err)
	}
	return instruction{
		command:  string(c),
		argument: a,
	}
}
