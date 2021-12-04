package main

import (
	"fmt"
	"math"
	"strconv"
)

//func parseDay3(string)

// func slurpDay3(path string) (int, []int64, error) {
// 	input, err := slurp(path)
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	codes := make([]int64, len(input))
// 	max := 0
// 	for i, line := range input {
// 		if len(line) > max {
// 			max = len(line)
// 		}
// 		if num, err := strconv.ParseInt(line, 2, 64); err != nil {
// 			return 0, nil, err
// 		} else {
// 			codes[i] = num
// 		}
// 	}
// 	return max, codes, nil
// }

func slurpDay3(path string) (int, []([]rune), error) {
	input, err := slurp(path)
	if err != nil {
		return 0, nil, err
	}
	codes := make([]([]rune), len(input))
	max := 0
	for i, line := range input {
		if len(line) > max {
			max = len(line)
		}
		codes[i] = []rune(line)
	}
	return max, codes, nil
}

func countFrequencies(codes []([]rune), maxLen int) map[int](map[rune]int) {
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

func gamma(freq map[int](map[rune]int), maxLen int) int64 {
	gr := make([]rune, maxLen)
	for i := 0; i < maxLen; i++ {
		if freq[i]['0'] > freq[i]['1'] {
			gr[i] = '0'
		} else {
			gr[i] = '1'
		}
	}
	gs := string(gr)
	g, _ := strconv.ParseInt(gs, 2, 64)
	mask := math.Pow(2, float64(maxLen)) - 1
	e := g ^ int64(mask)
	return g * e
}

func day3() {
	fmt.Println("\nDay 3 *******************")
	maxLen, codes, err := slurpDay3("input/day3.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(codes, maxLen)
	freq := countFrequencies(codes, maxLen)

	// fmt.Println(freq)
	fmt.Println(gamma(freq, maxLen))
}
