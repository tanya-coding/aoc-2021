package main

import (
	"fmt"
	"strconv"
)

// Returns
//   1st: slice where each element is a rune slice, e.g. ['0', '1', '1', '0']
//   2nd: error
func slurpDay3(path string) ([][]rune, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	codes := make([][]rune, len(input))
	for i, line := range input {
		codes[i] = []rune(line)
	}
	return codes, nil
}

// Counts frequencies of each bit in the list at its position
// Returns an array where item represents frequency of '0' and '1' at that position:
// e.g. [{'0': 123, '1': 876}, {'0': 654, '1': 233}, {'0': 34, '1': 98}, ...}
func countFrequencies(codes [][]rune) []map[rune]int {
	// We will assumes all codes are the same length
	frequencies := make([]map[rune]int, len(codes[0]))
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

func powerConsumption(freq []map[rune]int) int64 {
	codeLen := len(freq)
	gr := make([]rune, codeLen)
	for i, f := range freq {
		if f['0'] > f['1'] {
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

func OxGenBitCriteria(freq []map[rune]int, pos int) rune {
	if freq[pos]['1'] >= freq[pos]['0'] {
		return '1'
	}
	return '0'
}

func CO2ScrubBitCriteria(freq []map[rune]int, pos int) rune {
	if freq[pos]['0'] <= freq[pos]['1'] {
		return '0'
	}
	return '1'
}

type RuneFilter func([]rune) bool

type BitCriteria func([]map[rune]int, int) rune

func Filter(bitCriteria BitCriteria, freq []map[rune]int, pos int) RuneFilter {
	criteria := bitCriteria(freq, pos)
	return func(code []rune) bool {
		return criteria == code[pos]
	}
}

func dbg(codes [][]rune) {
	fmt.Println("Remaining", len(codes))
	for i := 0; i < 5; i++ {
		if i < len(codes) {
			fmt.Println(string(codes[i]))
		}
	}
	fmt.Println("...")
}

func MatchingCode(bitCriteria BitCriteria, freq []map[rune]int, codes [][]rune, pos int) []rune {
	// fmt.Println("Processing pos", pos, "Bit criteria: ", string(bitCriteria(freq, pos)), "Frequencies:", freq[pos])
	matches := Filter(bitCriteria, freq, pos)
	remaining := [][]rune{}
	for _, code := range codes {
		if matches(code) {
			remaining = append(remaining, code)
		}
	}
	// dbg(remaining)
	if len(remaining) == 0 {
		panic("Oops, didn't find a match")
	}
	if len(remaining) == 1 {
		return remaining[0]
	}
	return MatchingCode(bitCriteria, countFrequencies(remaining), remaining, pos+1)
}

func oxGenRating(codes [][]rune, freq []map[rune]int) int64 {
	code := MatchingCode(OxGenBitCriteria, freq, codes, 0)
	rating, _ := strconv.ParseInt(string(code), 2, 64)
	return rating
}

func co2ScrubRating(codes [][]rune, freq []map[rune]int) int64 {
	code := MatchingCode(CO2ScrubBitCriteria, freq, codes, 0)
	rating, _ := strconv.ParseInt(string(code), 2, 64)
	return rating
}

func day3() {
	fmt.Println("\nDay 3 *******************")
	codes, err := slurpDay3("input/day3.txt")
	if err != nil {
		panic(err)
	}
	freq := countFrequencies(codes)
	fmt.Println(powerConsumption(freq))
	fmt.Println(oxGenRating(codes, freq) * co2ScrubRating(codes, freq))
}
