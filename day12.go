package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Cave struct {
	name      string
	large     bool
	neighbors []Cave
}

// Visitor rules - part 1
func canVisit(caveMap map[string]Cave, path Path, next Cave) bool {
	if next.large {
		return true
	}
	// Check if small cave has been visited already
	for _, c := range path {
		if c == next.name {
			return false
		}
	}
	return true
}

// Visitor rules - part 2
func canVisit2(caveMap map[string]Cave, path Path, next Cave) bool {
	if next.large {
		return true
	}
	if next.name == "start" || next.name == "end" {
		return false
	}
	// Check if small cave has been visited already
	// We could change Path type to collect these stats
	// so we don't need to recalculate it each time
	smallCnt := map[string]int{}
	limitReached := false
	for _, cn := range path {
		c := caveMap[cn]
		if !c.large {
			curr := smallCnt[cn]
			smallCnt[cn] = curr + 1
			if smallCnt[cn] > 1 {
				limitReached = true
			}
		}
	}
	return smallCnt[next.name] < 1 || smallCnt[next.name] == 1 && !limitReached
}

type Path []string

type CanVisitFunc func(caveMap map[string]Cave, path Path, next Cave) bool

func findPaths(canVisit CanVisitFunc, caveMap map[string]Cave, currLoc string, end string, currentPath Path, allPaths []Path) []Path {
	curr := caveMap[currLoc]
	for _, next := range curr.neighbors {
		switch {
		case next.name == end:
			// Found end, add new path
			path := append(currentPath, end)
			allPaths = append(allPaths, path)
		case canVisit(caveMap, currentPath, next):
			// Can visit neighbor, add it to current path and keep walkling/recur
			path := append(currentPath, next.name)
			allPaths = findPaths(canVisit, caveMap, next.name, end, path, allPaths)
		}
	}
	return allPaths
}

func findOrNew(caveMap map[string]Cave, name string) Cave {
	if c, ok := caveMap[name]; ok {
		return c
	} else {
		return Cave{name: name, large: unicode.IsUpper(rune(name[0])), neighbors: []Cave{}}
	}
}

func slurpDay12(path string) (map[string]Cave, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	caves := map[string]Cave{}
	for _, conn := range input {
		strs := strings.Split(conn, "-")
		n1 := strs[0]
		n2 := strs[1]
		c1 := findOrNew(caves, n1)
		c2 := findOrNew(caves, n2)
		c1.neighbors = append(c1.neighbors, c2)
		c2.neighbors = append(c2.neighbors, c1)
		caves[n1] = c1
		caves[n2] = c2
	}
	return caves, nil
}

func prnMap(caveMap map[string]Cave) {
	for n, c := range caveMap {
		names := make([]string, len(c.neighbors))
		for i, c := range c.neighbors {
			names[i] = c.name
		}
		fmt.Println(n, ": large?", c.large, ", neighbors:", names)
	}
}

func day12() {
	fmt.Println("\nDay 12 *******************")
	caveMap, err := slurpDay12("input/day12.txt")
	if err != nil {
		panic(err)
	}
	prnMap(caveMap)
	allPaths := findPaths(canVisit, caveMap, "start", "end", Path{"start"}, []Path{})
	// for _, p := range allPaths {
	// 	fmt.Println(p)
	// }
	fmt.Println("Found paths:", len(allPaths))

	allPaths2 := findPaths(canVisit2, caveMap, "start", "end", Path{"start"}, []Path{})
	// for _, p := range allPaths {
	// 	fmt.Println(p)
	// }
	fmt.Println("Found paths:", len(allPaths2))
}
