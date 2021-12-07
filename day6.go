package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(fishes []int, days int) []int {
	for d := 1; d <= days; d++ {
		newFishes := []int{}
		// fmt.Println("Begin day:", fishes)
		for f := 0; f < len(fishes); f++ {
			fishes[f]--
			if fishes[f] < 0 {
				fishes[f] = 6
				newFishes = append(newFishes, 8)
			}
		}
		fishes = append(fishes, newFishes...)
		// fmt.Println("End day:", fishes)
	}
	return fishes
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
	newFishes := multiply(fishes, 80)
	fmt.Println(len(newFishes))
}
