package main

import (
	"fmt"
	"math"
	"strconv"
)

// Returns
//   1st: maximum code length
//   2nd: slice where each element is a rune slice, e.g. ['0', '1', '1', '0']
func slurpDay3(path string) (int, []([]rune), error) {
	input, err := slurp(path)
	if err != nil {
		return 0, nil, err
	}
	codes := make([]([]rune), len(input))
	maxLen := 0
	for i, line := range input {
		if len(line) > maxLen {
			maxLen = len(line)
		}
		codes[i] = []rune(line)
	}
	return maxLen, codes, nil
}

// Counts frequencies of each bit in the list at its position
// Returns a map in shape {position: {'0': zero-frequency, '1': one-frequency}}
// e.g. {0: {'0': 123, '1': 876}, 1: {'0': 654, '1': 233}, 2: {'0': 34, '1': 98}, ...}
func countFrequencies(codes []([]rune), maxLen int) map[int](map[rune]int) {
	frequencies := map[int](map[rune]int){}
	for _, code := range codes {
		// In sample input all items are the same length but if they weren't
		// this would take care of starting at the right position
		offset := maxLen - len(code)
		for pos, bit := range code {
			if frequencies[offset+pos] == nil {
				frequencies[offset+pos] = map[rune]int{}
			}
			frequencies[offset+pos][bit] += 1
		}
	}
	return frequencies
}

func powerConsumption(freq map[int](map[rune]int), maxLen int) int64 {
	gr := make([]rune, maxLen)
	for i := 0; i < maxLen; i++ {
		if freq[i]['0'] > freq[i]['1'] {
			gr[i] = '0'
		} else {
			gr[i] = '1'
		}
	}
	gamma, _ := strconv.ParseInt(string(gr), 2, 64) // Ignoring error since we control the input
	mask := math.Pow(2, float64(maxLen)) - 1
	epsilon := gamma ^ int64(mask)
	return gamma * epsilon
}

func day3() {
	fmt.Println("\nDay 3 *******************")
	maxLen, codes, err := slurpDay3("input/day3.txt")
	if err != nil {
		panic(err)
	}
	freq := countFrequencies(codes, maxLen)
	fmt.Println(powerConsumption(freq, maxLen))
}
