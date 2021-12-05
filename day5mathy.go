package main

import (
	"fmt"
)

func between(p int, p1 int, p2 int) bool {
	if p1 <= p2 {
		return p1 <= p && p <= p2
	} else {
		return p2 <= p && p <= p1
	}
}

func analyzeDanger(lines [][4]int) int {
	dangerMap := map[[2]int]int{}
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			x1 := lines[i][0]
			y1 := lines[i][1]
			x2 := lines[i][2]
			y2 := lines[i][3]
			u1 := lines[j][0]
			v1 := lines[j][1]
			u2 := lines[j][2]
			v2 := lines[j][3]

			fmt.Println("Analyzing lines:", lines[i], "and", lines[j], x1, y1, x2, y2, u1, v1, u2, v2)
			switch {
			case x1 == x2 && v1 == v2 && x1 == v1: // Overlapping horizontal lines
				ys := -1
				ye := -1
				if between(y1, v1, v2) {
					ys = y1
				} else if between(v1, y1, y2) {
					ys = v1
				}
				if between(y2, v1, v2) {
					ye = y2
				} else if between(v2, y1, y2) {
					ye = v2
				}
				fmt.Println("Overlapping horizontal line:", x1, ys, x1, ye)
			}
			if !(x1 == x2 && u1 == u2 ||
				y1 == y2 && v1 == v2 ||
				x1 != x2 && y1 != y2 && u1 != u2 && v1 != v2) { // If lines are not parallel
				x := -1 * ((x1-x2)*(u1*v2-u2*v1) - (u2-u1)*(x2*y1-x1*y2)) / ((v1-v2)*(x1-x2) - (u2-u1)*(y2-y1))
				y := -1 * (u1*v2*y1 - u1*v2*y2 - u2*v1*y1 + u2*v1*y2 - v1*x1*y2 + v1*x2*y1 + v2*x1*y2 - v2*x2*y1) / (-1*u1*y1 + u1*y2 + u2*y1 - u2*y2 + v1*x1 - v1*x2 - v2*x1 + v2*x2)
				if between(x, x1, x2) && between(x, u1, u2) && between(y, y1, y2) && between(y, v1, v2) {
					fmt.Println("Intersection:", x, ",", y)
					dangerMap[[2]int{x, y}]++
				}
			}
		}
	}
	danger := 0
	for c, d := range dangerMap {
		if d >= 2 {
			danger++
			fmt.Println(c, ", danger:", d)
		}
	}
	return danger
	// fmt.Println("Added vent:", line)
	// fmt.Println("Map:", floorMap.ventMap)
}

func slurpDay5Mathy(path string) ([][4]int, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}

	lines := [][4]int{}
	for _, s := range input {
		x1, y1, x2, y2 := parseCoordinates(s)
		lines = append(lines, [4]int{x1, y1, x2, y2})
	}
	return lines, nil
}

func day5mathy() {
	fmt.Println("\nDay 5 *******************")
	floorMap, err := slurpDay5Mathy("input/day5small.txt")
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
	fmt.Println(analyzeDanger(floorMap))
}
