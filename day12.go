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

func canVisit(path []*Cave, next Cave) bool {
	// TODO: Add loop detection
	if next.large {
		return true
	}
	// Check if we just came from this node
	if len(path) > 0 && path[len(path)-1].name == next.name {
		return false
	}
	// Check if small cave has been visited already
	for _, c := range path {
		if c.name == next.name {
			return false
		}
	}
	return true
}

type Path []*Cave

func findPath(caveMap map[string]*Cave, currLoc string, end string, currentPath Path, allPaths *[]Path) {
	curr := caveMap[currLoc]
	fmt.Println("Path:", caveNames(currentPath), "Current: ", curr.name)
	for _, next := range curr.neighbors {
		if curr.name == "start" {
			fmt.Println("START PATH:", caveNames(currentPath))
		}
		fmt.Println("Checking next", next.name, canVisit(currentPath, *next))
		switch {
		case next.name == end:
			// Found end, add it to current path and add to all paths
			path := append(currentPath, next)
			*allPaths = append(*allPaths, path)
			fmt.Println("FOUND PATH", caveNames(path), "Paths so far", len(*allPaths))
		case canVisit(currentPath, *next):
			// Can visit neighbor, add it to current path and keep walkling/recur
			path := append(currentPath, next)
			findPath(caveMap, next.name, end, path, allPaths)
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

func caveNames(caves []*Cave) []string {
	names := make([]string, len(caves))
	for i, c := range caves {
		names[i] = c.name
	}
	return names
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
	caveMap, err := slurpDay12("input/day12small.txt")
	if err != nil {
		panic(err)
	}
	prnMap(caveMap)
	allPath := []Path{}
	start := caveMap["start"]
	findPath(caveMap, "start", "end", Path{start}, &allPath)
	for _, p := range allPath {
		fmt.Println(caveNames(p))
	}
	//fmt.Println(allPaths(caveMap, "start", "end"))
}
