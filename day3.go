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
// Returns an array where item represents frequency of '0' and '1' at that position:
// e.g. [{'0': 123, '1': 876}, {'0': 654, '1': 233}, {'0': 34, '1': 98}, ...}
func countFrequencies(codes []([]rune)) [](map[rune]int) {
	// We will assumes all codes are the same length
	frequencies := make([](map[rune]int), len(codes[0]))
	for _, code := range codes {
		for pos, bit := range code {
			if frequencies[pos] == nil {
				frequencies[pos] = map[rune]int{}
			}
			frequencies[pos][bit]++
		}
	}
	return frequencies
}

func powerConsumption(freq [](map[rune]int)) int64 {
	codeLen := len(freq)
	gr := make([]rune, codeLen)
	for i := 0; i < codeLen; i++ {
		if freq[i]['0'] > freq[i]['1'] {
			gr[i] = '0'
		} else {
			gr[i] = '1'
		}
	}
	gamma, _ := strconv.ParseInt(string(gr), 2, 64) // Ignoring error since we control the input
	mask := int64(1)<<codeLen - 1                   // Get mask of all 1s of maxLength: e.g. for maxLength 3: 111
	epsilon := gamma ^ mask                         // Bitwise XOR gives us epsilon
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
