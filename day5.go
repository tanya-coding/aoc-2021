package main

import (
	"fmt"
	"strconv"
	"strings"
)

type FloorMap struct {
	ventMap map[[2]int]int
	danger  int
}

func (floorMap *FloorMap) addVent(line [][2]int) {
	for _, coord := range line {
		floorMap.ventMap[coord]++
		if floorMap.ventMap[coord] == 2 {
			floorMap.danger++
		}
	}
}

func (floorMap FloorMap) countDanger() int {
	return floorMap.danger
}

// Parse line instruction into list of coordinates that represent
// a horizontal or vertical line
// args:
//	ls - line string in format x1,y1 -> x2,y2, e.g. 309,347 -> 309,464
func strToLine(ls string) [][2]int {
	parts := strings.Fields(ls)
	coord1 := strings.Split(parts[0], ",")
	coord2 := strings.Split(parts[2], ",")
	x1, _ := strconv.Atoi(coord1[0])
	y1, _ := strconv.Atoi(coord1[1])
	x2, _ := strconv.Atoi(coord2[0])
	y2, _ := strconv.Atoi(coord2[1])
	line := [][2]int{}
	// Normalize lines (from smaller x1,y1 to larger x2,y2) since direction doesn't matter
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}
	// Keeping it simple for now
	if x1 == x2 || y1 == y2 {
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				line = append(line, [2]int{x, y})
			}
		}
	}
	return line
}

func slurpDay5(path string) (*FloorMap, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	floorMap := FloorMap{ventMap: map[[2]int]int{}}
	for _, s := range input {
		line := strToLine(s)
		floorMap.addVent(line)
	}
	return &floorMap, nil
}

func day5() {
	fmt.Println("\nDay 5 *******************")
	floorMap, err := slurpDay5("input/day5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(floorMap.countDanger())
}
