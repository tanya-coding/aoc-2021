package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert simple array of fish timers into a short array
// where each item represents a geneation of fishes with the same
// timer, e.g.
// 0: num-fishes-with-timer-0
// 1: num-fishes-with-timer-1
// ...
// 8: num-fishes-with-timer-8
func compact(fishes []int) [9]int {
	compactFishes := [9]int{}
	for _, f := range fishes {
		compactFishes[f]++
	}
	return compactFishes
}

func multiply(generations [9]int, days int) [9]int {
	for d := 1; d <= days; d++ {
		reproduced := generations[0] // How many will produce offspring this day
		for gen := 1; gen < len(generations); gen++ {
			generations[gen-1] = generations[gen] // Promote generation
		}
		generations[8] = reproduced  // New offspring
		generations[6] += reproduced // Reset timer of fishes that produced offspring
	}
	return generations
}

func countFishes(generations [9]int) int {
	total := 0
	for _, c := range generations {
		total += c
	}
	return total
}

func slurpDay6(path string) ([]int, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	timers := strings.Split(input[0], ",")
	fishes := make([]int, len(timers))
	for i, t := range timers {
		fishes[i], _ = strconv.Atoi(t)
	}
	return fishes, nil
}

func day6() {
	fmt.Println("\nDay 6 *******************")
	fishes, err := slurpDay6("input/day6.txt")
	if err != nil {
		panic(err)
	}
	generations := multiply(compact(fishes), 256)
	fmt.Println(countFishes(generations))
}
