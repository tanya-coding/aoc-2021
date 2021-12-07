package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert simple array of fishes timers into a short array
// where each item represents a bucket of fishes with the same
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

func multiply(fishes [9]int, days int) [9]int {
	for d := 1; d <= days; d++ {
		// fmt.Println("Begin day:", fishes)
		reproduced := fishes[0]
		for gen := 1; gen < len(fishes); gen++ {
			fishes[gen-1] = fishes[gen]
		}
		fishes[8] = reproduced
		fishes[6] += reproduced
		// fmt.Println("End day:", fishes)
	}
	return fishes
}

func countFishes(fishes [9]int) int {
	total := 0
	for _, c := range fishes {
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
