package main

import (
	"bufio"
	"fmt"
	"os"
)

var MATCH = []rune{'X', 'M', 'A', 'S'}
var MATCH_TWO = []rune{'M', 'A', 'S'}
var MATCH_TWO_REV = []rune{'S', 'A', 'M'}

func readInput(input string) [][]rune {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	var grid [][]rune
	for scanner.Scan() {
		lineString := scanner.Text()

		var lineRune []rune
		for _, ch := range lineString {
			lineRune = append(lineRune, ch)
		}
		grid = append(grid, lineRune)
	}

	return grid
}

func matchOffset(grid [][]rune, i int, j int, match []rune, idx int, iOff int, jOff int) bool {
	if idx == len(match) {
		return true
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}

	ch := match[idx]
	if grid[i][j] != ch {
		return false
	}

	return matchOffset(grid, i+iOff, j+jOff, match, idx+1, iOff, jOff)
}

func match(grid [][]rune, match []rune, i int, j int) int {
	numMatches := 0
	for iOff := -1; iOff <= 1; iOff++ {
		for jOff := -1; jOff <= 1; jOff++ {
			if matchOffset(grid, i, j, match, 0, iOff, jOff) {
				numMatches++
			}
		}
	}
	return numMatches
}

func diags(grid [][]rune, i int, j int, match []rune) bool {
	downRight := matchOffset(grid, i, j, match, 0, 1, 1)
	if downRight {
		downLeft := matchOffset(grid, i, j+2, match, 0, 1, -1)
		upRight := matchOffset(grid, i+2, j, match, 0, -1, 1)
		if downLeft || upRight {
			return true
		}
	}
	return false
}

func partOne(input string) {
	grid := readInput(input)

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += match(grid, MATCH, i, j)
		}
	}

	fmt.Printf("Output: %d\n", sum)

}

func partTwo(input string) {
	grid := readInput(input)

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if diags(grid, i, j, MATCH_TWO) {
				sum++
			} else if diags(grid, i, j, MATCH_TWO_REV) {
				sum++
			}
		}
	}

	fmt.Printf("Output: %d\n", sum)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
