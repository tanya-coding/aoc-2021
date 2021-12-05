package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lineToNums(line string, parse func(string) []string) []int {
	numStrs := parse(line)
	nums := make([]int, len(numStrs))
	for i, s := range numStrs {
		num, _ := strconv.Atoi(s) // Assuming input is valid
		nums[i] = num
	}
	return nums
}

type Board struct {
	items   map[int][2]int  // board items: number -> coordinate
	marked  map[[2]int]bool // marked numbers: coordinate -> true when marked
	markedX map[int]int     // count of X coordinate marked
	markedY map[int]int     // count of Y coordinate marked
}

func prn(board Board) {
	fmt.Println("items\t", board.items)
	fmt.Println("marked\t", board.marked)
	fmt.Println("markedX\t", board.markedX)
	fmt.Println("markedY\t", board.markedY)
}

func mark(board *Board, num int) bool {
	if coord, contains := board.items[num]; contains {
		board.marked[coord] = true
		x := coord[0]
		y := coord[1]
		board.markedX[x]++
		board.markedY[y]++
		return board.markedX[x] == 5 || board.markedY[y] == 5
	}
	return false
}

func resetBoard(board *Board) {
	board.marked = map[[2]int]bool{}
	board.markedX = map[int]int{}
	board.markedY = map[int]int{}
}

func sumUnmarked(board Board) int {
	sum := 0
	for num, coord := range board.items {
		if !board.marked[coord] {
			sum += num
		}
	}
	return sum
}

func newBoard(lines []string) *Board {
	board := Board{items: map[int][2]int{}, marked: map[[2]int]bool{}, markedX: map[int]int{}, markedY: map[int]int{}}
	for y := 0; y < 5; y++ {
		nums := lineToNums(lines[y], func(s string) []string {
			return strings.Fields(s)
		})
		for x, num := range nums {
			board.items[num] = [2]int{x, y}
		}
	}
	return &board
}

func slurpDay4(path string) ([]int, []*Board, error) {
	input, err := slurp(path)
	if err != nil {
		return nil, nil, err
	}
	numbers := lineToNums(input[0], func(s string) []string {
		return strings.Split(s, ",")
	})
	numBoards := (len(input) - 1) / 6
	boards := make([]*Board, numBoards)

	for i := 0; i < numBoards; i++ {
		idx := i*6 + 1
		boards[i] = newBoard(input[idx+1 : idx+6])
	}
	return numbers, boards, nil
}

func drawFirstWins(numbers []int, boards []*Board) int {
	for _, num := range numbers {
		for _, board := range boards {
			if mark(board, num) {
				fmt.Println("Bingo. Number", num)
				prn(*board)
				return sumUnmarked(*board) * num
			}
		}
	}
	return 0
}

func drawLastWins(numbers []int, boards []*Board) int {
	wins := map[int]bool{}
	var winningBoard Board
	winningNum := 0
	for _, num := range numbers {
		for bi, board := range boards {
			// Skip previous winners
			if !wins[bi] && mark(board, num) {
				winningBoard = *board
				winningNum = num
				wins[bi] = true
			}
		}
	}
	fmt.Println("Bingo. Number", winningNum)
	prn(winningBoard)
	return sumUnmarked(winningBoard) * winningNum
}

func resetBoards(boards []*Board) {
	for _, b := range boards {
		resetBoard(b)
	}
}

func day4() {
	fmt.Println("\nDay 4 *******************")
	numbers, boards, err := slurpDay4("input/day4.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(drawFirstWins(numbers, boards))
	resetBoards(boards)
	fmt.Println(drawLastWins(numbers, boards))
}
