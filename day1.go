package main

import (
	"fmt"
	"strconv"
)

func countIncreases(input []int) int {
	increased := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increased++
		}
	}
	return increased
}

func countSlidingWindow(input []int, size int) int {
	currentWindow := 0
	increased := 0

	for i := 0; i < len(input); i++ {
		if i < size {
			// not yet a complete window
			currentWindow += input[i]
			continue
		}

		lastWindow := currentWindow
		// remove trailing element and add current element to maintain window
		currentWindow += input[i] - input[i-size]
		if currentWindow > lastWindow {
			increased++
		}
	}

	return increased
}

func parseDay1(input []string) ([]int, error) {
	parsed := []int{}
	for _, line := range input {
		intVar, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, intVar)
	}
	return parsed, nil
}

func slurpDay1(path string) ([]int, error) {
	input, err := slurp("day1.txt")
	if err != nil {
		return nil, err
	}
	parsed, err := parseDay1(input)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

func day1() {
	input, err := slurpDay1("day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(countIncreases(input))
	fmt.Println(countSlidingWindow(input, 3))
}
