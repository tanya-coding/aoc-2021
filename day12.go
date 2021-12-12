package main

import (
	"fmt"
	"strings"
	"unicode"
)

// start
// end
// node
//  large: true/false
//  neighbors (bi-directional)

// something to uniquely identify path, or collect path and check before going in if next item can be visited given a path so far
// need to detect deadends

// data structure:
//   collecting path: creating new collections as we branch
//   check if given path so far and the next item, item can be walked
//     if it can be add it to path and walk all its neighbours (except the node you just came from).

type Cave struct {
	name      string
	large     bool
	neighbors []*Cave
}

func canVisit(path []string, next Cave) bool {
	// TODO: Add loop detection
	if next.large {
		return true
	}
	// Check if we just came from this node
	// if len(path) > 0 && path[len(path)-1] == next.name {
	// 	return false
	// }
	// Check if small cave has been visited already
	for _, c := range path {
		if c == next.name {
			return false
		}
	}
	return true
}

type Path []string

func findPaths(caveMap map[string]*Cave, currLoc string, end string, currentPath Path, allPaths *[]Path) {
	curr := caveMap[currLoc]
	path := make(Path, len(currentPath))
	copy(path, currentPath)
	for _, next := range curr.neighbors {
		switch {
		case next.name == end:
			// Found end, add it to current path and add to all paths
			path := append(path, next.name)
			// fmt.Println("FOUND PATH", currentCopy, "Current path", currentPath, "Paths so far", len(*allPaths))
			*allPaths = append(*allPaths, path)
		case canVisit(currentPath, *next):
			// Can visit neighbor, add it to current path and keep walkling/recur
			path := append(path, next.name)
			findPaths(caveMap, next.name, end, path, allPaths)
		}
	}
}

func findOrNew(caveMap map[string]*Cave, name string) *Cave {
	if c, ok := caveMap[name]; ok {
		return c
	} else {
		return &Cave{name: name, large: unicode.IsUpper(rune(name[0])), neighbors: []*Cave{}}
	}
}

func slurpDay12(path string) (map[string]*Cave, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	caves := map[string]*Cave{}
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

func prnMap(caveMap map[string]*Cave) {
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
	allPaths := []Path{}
	findPaths(caveMap, "start", "end", Path{"start"}, &allPaths)
	for _, p := range allPaths {
		fmt.Println(p)
	}
	fmt.Println("Found paths:", len(allPaths))
}
