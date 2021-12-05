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
	// fmt.Println("Added vent:", line)
	// fmt.Println("Map:", floorMap.ventMap)
}

// Parse line instruction in format 'x1,y1 -> x2,y2' and
// return (x1, y1, x2, y2)
func parseCoordinates(ls string) (int, int, int, int) {
	parts := strings.Fields(ls)
	coord1 := strings.Split(parts[0], ",")
	coord2 := strings.Split(parts[2], ",")
	x1, _ := strconv.Atoi(coord1[0])
	y1, _ := strconv.Atoi(coord1[1])
	x2, _ := strconv.Atoi(coord2[0])
	y2, _ := strconv.Atoi(coord2[1])

	return x1, y1, x2, y2
}

// Parse line instruction into list of coordinates
// that represents a line
// args:
//	ls - line string in format 'x1,y1 -> x2,y2', e.g. 309,347 -> 309,464
// returns:
// 	line points: [[x1, y1], [xn, yn], ... [x2, y2] ]
func strToLine(ls string) [][2]int {
	x1, y1, x2, y2 := parseCoordinates(ls)
	line := [][2]int{}
	xDir := 0
	if x1 > x2 {
		xDir = -1
	} else if x1 < x2 {
		xDir = 1
	}
	yDir := 0
	if y1 > y2 {
		yDir = -1
	} else if y1 < y2 {
		yDir = 1
	}
	x := x1
	y := y1
	for {
		line = append(line, [2]int{x, y})
		if x == x2 && y == y2 {
			return line
		}
		x += xDir
		y += yDir
	}
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
	// fmt.Println(floorMap.ventMap)
	// We can walk coordinates instead of using precalculated value:
	// d := 0
	// for _, danger := range floorMap.ventMap {
	// 	if danger >= 2 {
	// 		d++
	// 	}
	// }
	// fmt.Println(d)
	fmt.Println(floorMap.danger)
}
