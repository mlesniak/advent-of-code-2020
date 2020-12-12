package main

import (
	"fmt"
	"math"
	"strconv"
)

type point struct {
	x int
	y int
}

type ship struct {
	direction int
	x         int
	y         int
}

func (s ship) String() string {
	return fmt.Sprintf("{x=%v y=%v dir=%v}", s.x, s.y, s.direction)
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

	waypoint := point{
		x: 10,
		y: 1,
	}

	for _, instr := range instructions {
		fmt.Printf("\nship=%v waypoint=%v\n", ship, waypoint)
		fmt.Printf("%v\n", instr)
		i := parseInstruction(instr)
		switch i.command {
		case "N":
			waypoint.y += i.argument
		case "S":
			waypoint.y -= i.argument
		case "E":
			waypoint.x += i.argument
		case "W":
			waypoint.x -= i.argument
		case "L":
			rad := float64(i.argument) * math.Pi / 180.0
			nx := float64(waypoint.x)*math.Cos(rad) - float64(waypoint.y)*math.Sin(rad)
			ny := float64(waypoint.y)*math.Cos(rad) + float64(waypoint.x)*math.Sin(rad)
			waypoint.x = int(math.Round(nx))
			waypoint.y = int(math.Round(ny))
		case "R":
			rad := float64(-i.argument) * math.Pi / 180.0
			nx := float64(waypoint.x)*math.Cos(rad) - float64(waypoint.y)*math.Sin(rad)
			ny := float64(waypoint.y)*math.Cos(rad) + float64(waypoint.x)*math.Sin(rad)
			waypoint.x = int(math.Round(nx))
			waypoint.y = int(math.Round(ny))
		case "F":
			ship.x += i.argument * waypoint.x
			ship.y += i.argument * waypoint.y
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
