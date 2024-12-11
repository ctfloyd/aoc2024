package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Position struct {
	row, col int
}

var DIRECTIONS = [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func readInput(input string) [][]int {
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

	var grid [][]int
	for scanner.Scan() {
		lineString := scanner.Text()

		var line []int
		for _, ch := range lineString {
			line = append(line, int(ch-'0'))
		}
		grid = append(grid, line)
	}

	return grid
}

func findValidNeighbors(position Position, elevation int, grid [][]int) []Position {
	validNeighbors := []Position{}
	for _, offset := range DIRECTIONS {
		newPosition := Position{row: position.row + offset[0], col: position.col + offset[1]}
		if newPosition.col >= 0 && newPosition.row >= 0 && newPosition.row < len(grid) && newPosition.col < len(grid[newPosition.row]) {
			if grid[newPosition.row][newPosition.col] == elevation+1 {
				validNeighbors = append(validNeighbors, newPosition)
			}
		}
	}
	return validNeighbors
}

func combineSummits(summits, others []Position) []Position {
	for _, summit := range others {
		if !slices.Contains(summits, summit) {
			summits = append(summits, summit)
		}
	}
	return summits
}

func navigate(position Position, grid [][]int, uniqueTrailScore bool) []Position {
	elevation := grid[position.row][position.col]

	if elevation == 9 {
		return []Position{position}
	}

	summits := []Position{}
	neighbors := findValidNeighbors(position, elevation, grid)
	for _, neighbor := range neighbors {
		ns := navigate(neighbor, grid, uniqueTrailScore)
		if !uniqueTrailScore {
			summits = combineSummits(summits, ns)
		} else {
			summits = append(summits, ns...)
		}
	}
	return summits
}

func solve(grid [][]int, uniqueTrailPosition bool) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				sum += len(navigate(Position{row: i, col: j}, grid, uniqueTrailPosition))
			}
		}
	}
	return sum
}

func partOne(input string) {
	grid := readInput(input)
	fmt.Printf("result: %v\n", solve(grid, false))
}

func partTwo(input string) {
	grid := readInput(input)
	fmt.Printf("result: %v\n", solve(grid, true))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
