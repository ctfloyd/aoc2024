package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func readInput(input string) [][]rune {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}
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

type Pos struct {
	row, col int
}

type Border struct {
	origin, border Pos
}

type Plot struct {
	identifier rune
	positions  []Pos
}

func isInRange(pos Pos, grid [][]rune) bool {
	return pos.row >= 0 && pos.col >= 0 && pos.row < len(grid) && pos.col < len(grid[pos.row])
}

func findInRangeNeighbors(pos Pos, grid [][]rune) []Pos {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	neighbors := []Pos{}
	for _, dir := range dirs {
		neighbor := Pos{row: pos.row + dir[0], col: pos.col + dir[1]}
		if isInRange(neighbor, grid) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func explorePlot(pos Pos, grid [][]rune) Plot {
	plot := Plot{identifier: grid[pos.row][pos.col]}

	explored := []Pos{}
	queue := []Pos{pos}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if !slices.Contains(explored, p) {
			if grid[p.row][p.col] == plot.identifier {
				plot.positions = append(plot.positions, p)
				neighbors := findInRangeNeighbors(p, grid)
				queue = append(queue, neighbors...)
			}

			explored = append(explored, p)
		}
	}

	return plot
}

func findPlots(grid [][]rune) []Plot {
	plots := []Plot{}
	plotted := []Pos{}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			pos := Pos{row, col}
			if !slices.Contains(plotted, pos) {
				plot := explorePlot(pos, grid)
				plotted = append(plotted, plot.positions...)
				plots = append(plots, plot)
			}
		}
	}

	return plots
}

func adjacentPositions(pos Pos) []Pos {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	positions := []Pos{}
	for _, dir := range dirs {
		positions = append(positions, Pos{row: pos.row + dir[0], col: pos.col + dir[1]})
	}
	return positions
}

func adjacent(pos Pos, plot Plot, inPlot bool) []Pos {
	adjacent := []Pos{}
	for _, p := range adjacentPositions(pos) {
		if inPlot && slices.Contains(plot.positions, p) {
			adjacent = append(adjacent, p)
		} else if !inPlot && !slices.Contains(plot.positions, p) {
			adjacent = append(adjacent, p)
		}
	}
	return adjacent
}

func calcPlot(plot Plot) int {
	area := len(plot.positions)

	perimeter := 0
	for _, pos := range plot.positions {
		perimeter += 4 - len(adjacent(pos, plot, true))
	}

	return area * perimeter
}

func calc(plots []Plot) int {
	sum := 0
	for _, plot := range plots {
		sum += calcPlot(plot)
	}
	return sum
}

func findBorderPoints(plot Plot) []Border {
	borders := []Border{}
	for _, pos := range plot.positions {
		adjacent := adjacent(pos, plot, false)
		for _, adj := range adjacent {
			borders = append(borders, Border{origin: pos, border: adj})
		}
	}
	return borders
}

func visitBorder(border Border, borders []Border, visited []Border) []Border {
	if slices.Contains(visited, border) || !slices.Contains(borders, border) {
		return visited
	}

	visited = append(visited, border)

	if border.origin.row == border.border.row {
		visited = visitBorder(Border{
			origin: Pos{row: border.origin.row - 1, col: border.origin.col},
			border: Pos{row: border.border.row - 1, col: border.border.col},
		}, borders, visited)
		return visitBorder(Border{
			origin: Pos{row: border.origin.row + 1, col: border.origin.col},
			border: Pos{row: border.border.row + 1, col: border.border.col},
		}, borders, visited)
	} else {
		visited = visitBorder(Border{
			origin: Pos{row: border.origin.row, col: border.origin.col - 1},
			border: Pos{row: border.border.row, col: border.border.col - 1},
		}, borders, visited)
		return visitBorder(Border{
			origin: Pos{row: border.origin.row, col: border.origin.col + 1},
			border: Pos{row: border.border.row, col: border.border.col + 1},
		}, borders, visited)
	}
}

func countSides(plot Plot) int {
	borders := findBorderPoints(plot)
	visited := []Border{}

	sides := 0
	for _, border := range borders {
		if slices.Contains(visited, border) {
			continue
		}
		sides += 1
		visited = visitBorder(border, borders, visited)
	}
	return sides
}

func calcPlot2(plot Plot) int {
	area := len(plot.positions)
	sides := countSides(plot)
	fmt.Printf("Found %v sides for plot %s\n", sides, string(plot.identifier))
	return area * sides
}

func calc2(plots []Plot) int {
	sum := 0
	for _, plot := range plots {
		sum += calcPlot2(plot)
	}
	return sum
}

func partOne(input string) {
	grid := readInput(input)
	plots := findPlots(grid)
	fmt.Printf("Output %v\n", calc(plots))
}

func partTwo(input string) {
	grid := readInput(input)
	plots := findPlots(grid)
	fmt.Printf("Output %v\n", calc2(plots))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
