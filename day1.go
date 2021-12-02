package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	input := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, intVar)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

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

func day1() {
	input, err := readInput("/Users/tanya/Downloads/aoc-day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(countIncreases(input))
	fmt.Println(countSlidingWindow(input, 3))
}
