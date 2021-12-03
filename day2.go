package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	unit      int
}

type Location struct {
	depth    int
	position int
}

type Location2 struct {
	depth    int
	position int
	aim      int
}

func parseDay2(input []string) ([]Instruction, error) {
	parsed := []Instruction{}
	for _, line := range input {
		s := strings.Fields(line)
		u, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, Instruction{direction: s[0], unit: u})
	}
	return parsed, nil
}

func slurpDay2(path string) ([]Instruction, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, err
	}
	return parseDay2(input)
}

func advance(location Location, instruction Instruction) Location {
	switch instruction.direction {
	case "down":
		location.depth += instruction.unit
	case "up":
		location.depth -= instruction.unit
	case "forward":
		location.position += instruction.unit
	}
	return location
}

func getLocation(instructions []Instruction) int {
	location := Location{}
	for _, i := range instructions {
		location = advance(location, i)
	}
	return location.depth * location.position
}

func advance2(location Location2, instruction Instruction) Location2 {
	switch instruction.direction {
	case "down":
		location.aim += instruction.unit
	case "up":
		location.aim -= instruction.unit
	case "forward":
		location.position += instruction.unit
		location.depth += location.aim * instruction.unit
	}
	return location
}

func getLocation2(instructions []Instruction) int {
	location := Location2{}
	for _, i := range instructions {
		location = advance2(location, i)
	}
	return location.depth * location.position
}

func day2() {
	fmt.Println("\nDay 2 *******************")
	instructions, err := slurpDay2("day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(getLocation(instructions))
	fmt.Println(getLocation2(instructions))
}
