package main

import (
	"fmt"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minMaxOf(vars ...int) (int, int) {
	minVal := vars[0]
	maxVal := vars[0]
	for _, v := range vars {
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}
	return minVal, maxVal
}

// Aligns crabs at given position and calculates fuel spent
func align(crabs []int, pos int) int {
	fuel := 0
	for _, c := range crabs {
		distance := max(c, pos) - min(c, pos)
		// Using formula https://www.cuemath.com/sum-of-natural-numbers-formula/
		fuel += distance * (distance + 1) / 2
	}
	return fuel
}

func bestPosition(crabs []int) (int, int) {
	minPos, maxPos := minMaxOf(crabs...)
	best := minPos
	fuel := align(crabs, minPos)
	for i := minPos + 1; i <= maxPos; i++ {
		f := align(crabs, i)
		if f <= fuel {
			fuel = f
			best = i
		}
	}
	return best, fuel
}

func slurpDay7(path string) ([]int, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	strs := strings.Split(input[0], ",")
	crabs := make([]int, len(strs))
	for i, t := range strs {
		crabs[i], _ = strconv.Atoi(t)
	}
	return crabs, nil
}

func day7() {
	fmt.Println("\nDay 7 *******************")
	crabs, err := slurpDay6("input/day7.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(bestPosition(crabs))
}
