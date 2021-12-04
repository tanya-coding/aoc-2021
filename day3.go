package main

import (
	"fmt"
	"strconv"
)

// Returns
//   1st: slice where each element is a rune slice, e.g. ['0', '1', '1', '0']
//   2nd: error
func slurpDay3(path string) ([]([]rune), error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	codes := make([]([]rune), len(input))
	for i, line := range input {
		codes[i] = []rune(line)
	}
	return codes, nil
}

// Counts frequencies of each bit in the list at its position
// Returns a map in shape {position: {'0': zero-frequency, '1': one-frequency}}
// e.g. {0: {'0': 123, '1': 876}, 1: {'0': 654, '1': 233}, 2: {'0': 34, '1': 98}, ...}
func countFrequencies(codes []([]rune)) map[int](map[rune]int) {
	frequencies := map[int](map[rune]int){}
	for _, code := range codes {
		for pos, bit := range code {
			if frequencies[pos] == nil {
				frequencies[pos] = map[rune]int{}
			}
			frequencies[pos][bit] += 1
		}
	}
	return frequencies
}

func powerConsumption(freq map[int](map[rune]int)) int64 {
	maxLen := len(freq)
	gr := make([]rune, maxLen)
	for i := 0; i < len(freq); i++ {
		if freq[i]['0'] > freq[i]['1'] {
			gr[i] = '0'
		} else {
			gr[i] = '1'
		}
	}
	gamma, _ := strconv.ParseInt(string(gr), 2, 64) // Ignoring error since we control the input
	mask := 1<<maxLen - 1                           // Get mask of all 1s of maxLength: e.g. for maxLenth 3: 111
	epsilon := gamma ^ int64(mask)
	return gamma * epsilon
}

func day3() {
	fmt.Println("\nDay 3 *******************")
	codes, err := slurpDay3("input/day3.txt")
	if err != nil {
		panic(err)
	}
	freq := countFrequencies(codes)
	fmt.Println(powerConsumption(freq))
}
