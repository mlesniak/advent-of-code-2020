package main

import (
	"fmt"
	"math"
)

func day10() {
	adapters := readNumbers("input/10.txt")

	currentJolt := 0
	jolt1 := 0
	jolt3 := 0

	for len(adapters) > 0 {
		minJolt := math.MaxInt64
		index := 0

		for i, jolt := range adapters {
			if jolt <= currentJolt+3 {
				if jolt < minJolt {
					index = i
					minJolt = jolt
				}
			}
		}
		fmt.Printf("Found jolt %d at %d\n", minJolt, index)
		if minJolt == currentJolt+1 {
			jolt1++
		}
		if minJolt == currentJolt+3 {
			jolt3++
		}

		currentJolt = minJolt
		adapters = append(adapters[:index], adapters[index+1:]...)
	}

	// Add final adapter
	jolt3++

	println(jolt1)
	println(jolt3)
	println(jolt1 * jolt3)
}
