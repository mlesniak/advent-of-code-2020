package main

import (
	"fmt"
	"math"
	"sort"
)

func day10() {
	adapters := readNumbers("input/10.txt")

	//d10part1(adapters)
	d10part2(adapters)
}

func d10part2(adapters []int) {
	sort.Ints(adapters)
	counter := make(map[int]int) // Cool property of maps: for non-existing values we have 0

	adapters = append(adapters, adapters[len(adapters)-1]+3)
	counter[0] = 1

	// Walk through the array and see how many ways we had to get to the particular jolt
	// using memoization. This is way easier than doing combinatorics.
	for _, jolt := range adapters {
		counter[jolt] = counter[jolt-1] + counter[jolt-2] + counter[jolt-3]
	}

	println(counter[adapters[len(adapters)-1]])
}

func d10part1(adapters []int) {
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
