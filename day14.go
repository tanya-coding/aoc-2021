package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Pair [2]rune

type InsertionRule struct {
	pair   Pair
	insert rune
}

type Polymer struct {
	items  *list.List               // Polymer letters linked list
	index  map[rune][]*list.Element // Track locations of letters in the linked list
	counts map[rune]int             // Count letters
}

func applyRules(polymer Polymer, rules []InsertionRule) Polymer {
	operations := map[*list.Element]rune{}
	for _, rule := range rules {
		l1 := rule.pair[0]
		l2 := rule.pair[1]
		elements := polymer.index[l1]
		for _, elem := range elements {
			if elem.Next() != nil && elem.Next().Value == l2 {
				// Found pair, add operation to insert letter
				operations[elem] = rule.insert
			}
		}
	}
	// Apply operations
	for el, letter := range operations {
		newEl := polymer.items.InsertAfter(letter, el)
		polymer.index[letter] = append(polymer.index[letter], newEl)
		polymer.counts[letter]++
	}
	return polymer
}

func slurpDay14(path string) (Polymer, []InsertionRule) {
	input, err := slurp(path)
	if err != nil {
		panic(err)
	}

	polymer := list.New()
	index := map[rune][]*list.Element{}
	counts := map[rune]int{}
	for _, letter := range input[0] {
		elem := polymer.PushBack(letter)
		index[letter] = append(index[letter], elem)
		counts[letter]++
	}
	rules := []InsertionRule{}
	for i := 2; i < len(input); i++ {
		ruleParts := strings.Split(input[i], " -> ")

		l1 := rune(ruleParts[0][0])
		l2 := rune(ruleParts[0][1])
		rule := InsertionRule{pair: Pair{l1, l2}, insert: rune(ruleParts[1][0])}
		rules = append(rules, rule)
	}
	return Polymer{items: polymer, index: index, counts: counts}, rules
}

func minAndMax(polymer Polymer) (int, int) {
	minCnt := -1
	maxCnt := 0
	for _, c := range polymer.counts {
		if minCnt > c || minCnt < 0 {
			minCnt = c
		}
		if maxCnt < c {
			maxCnt = c
		}
	}
	return minCnt, maxCnt
}

func toStr(polymer Polymer) string {
	var sb strings.Builder
	for el := polymer.items.Front(); el != nil; el = el.Next() {
		sb.WriteString(string(rune(el.Value.(int32))))
	}
	return sb.String()
}

func day14() {
	polymer, rules := slurpDay14("input/day14.txt")
	for i := 0; i < 10; i++ {
		applyRules(polymer, rules)
	}
	min, max := minAndMax(polymer)
	fmt.Println(min, max)
	fmt.Println(max - min)
}
