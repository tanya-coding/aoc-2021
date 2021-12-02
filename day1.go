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

	sums := make([]int, l)
	for i := 0; i < l; i++ {
		sums[i] = 0
		if i+size <= l {
			for j := 0; j < size; j++ {
				sums[i] = sums[i] + input[i+j]
			}
		}
	}
	return countIncreases(sums)
}

func day1() {
	inputPath := "/Users/tanya/Downloads/aoc-day1-1.txt"
	input, err := readInput(inputPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(countIncreases(input))
	fmt.Println(countSlidingWindow(input, 3))
}
