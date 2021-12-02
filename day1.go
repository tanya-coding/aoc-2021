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
	l := len(input)
	numWindows := l - l%size

	sums := make([]int, numWindows)
	increased := 0
	for i := 0; i < numWindows; i++ {
		sums[i] = 0
		// Calculate sliding window (looking forward)
		for j := 0; j < size; j++ {
			sums[i] = sums[i] + input[i+j]
		}
		// Check if new sum is an increase
		if i > 0 && sums[i] > sums[i-1] {
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
