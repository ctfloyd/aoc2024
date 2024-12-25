package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"time"
)

func readInput(input string) [][]rune {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func findStart(grid [][]rune) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'S' {
				return row, col
			}
		}
	}
	return -1, -1
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func tillEnd(position [2]int, grid [][]rune, visited [][2]int) ([][2]int, map[[2]int]int) {
	visited = append(visited, position)
	if grid[position[0]][position[1]] == 'E' {
		m := make(map[[2]int]int)
		for i, p := range visited {
			m[p] = len(visited) - i - 1
		}
		return visited, m
	}
	for _, d := range dirs {
		nn := [2]int{position[0] + d[0], position[1] + d[1]}
		if !slices.Contains(visited, nn) && nn[0] >= 0 && nn[1] >= 0 && nn[0] < len(grid) && nn[1] < len(grid[nn[0]]) {
			if grid[nn[0]][nn[1]] != '#' {
				p, te := tillEnd(nn, grid, visited)
				if len(te) > 0 {
					return p, te
				}
			}
		}
	}
	return [][2]int{}, map[[2]int]int{}
}

var memo = make(map[int][][2]int)

func findAllPointsNDistanceAway(pos [2]int, distance int) [][2]int {
	offsets := [][2]int{}
	if o, ok := memo[distance]; ok {
		offsets = o
	} else {
		og := distance
		q := [][2]int{{0, 0}}
		nq := [][2]int{}
		for distance > 0 {
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				for _, d := range dirs {
					nn := [2]int{p[0] + d[0], p[1] + d[1]}
					if !slices.Contains(offsets, nn) {
						offsets = append(offsets, nn)
						nq = append(nq, nn)
					}
				}
			}
			q = nq
			nq = [][2]int{}
			distance -= 1
		}
		memo[og] = offsets
	}

	points := [][2]int{}
	for _, of := range offsets {
		points = append(points, [2]int{pos[0] + of[0], pos[1] + of[1]})
	}
	return points
}

func calculateActualDistance(pos, option [2]int) int {
	return int(math.Abs(float64(pos[0]-option[0]))) + int(math.Abs(float64(pos[1]-option[1])))
}

func solve(path [][2]int, tillEnd map[[2]int]int, distance int, totalLength int) []int {
	lengths := []int{}
	for i, pos := range path {
		options := findAllPointsNDistanceAway(pos, distance)
		for _, option := range options {
			actualDistance := calculateActualDistance(pos, option)
			if v, ok := tillEnd[option]; ok && v+i+actualDistance < totalLength {
				lengths = append(lengths, v+i+actualDistance)
			}
		}
	}
	return lengths
}

func countSaved(pathLength int, allLengths []int) map[int]int {
	m := make(map[int]int)
	for _, l := range allLengths {
		saved := pathLength - l
		if saved > 0 {
			m[saved]++
		}
	}
	return m
}

func partOne(input string) {
	grid := readInput(input)
	startRow, startCol := findStart(grid)

	start := [2]int{startRow, startCol}
	path, tillEnd := tillEnd(start, grid, [][2]int{})
	lengths := solve(path, tillEnd, 2, tillEnd[start])
	saved := countSaved(tillEnd[start], lengths)

	sum := 0
	for k, v := range saved {
		if k >= 100 {
			sum += v
		}
	}
	fmt.Printf("result %v\n", sum)
}

func partTwo(input string) {
	grid := readInput(input)

	now := time.Now()
	startRow, startCol := findStart(grid)
	start := [2]int{startRow, startCol}
	path, tillEnd := tillEnd(start, grid, [][2]int{})
	lengths := solve(path, tillEnd, 20, tillEnd[start])
	saved := countSaved(tillEnd[start], lengths)

	sum := 0
	for k, v := range saved {
		if k >= 100 {
			sum += v
		}
	}
	fmt.Printf("result %v %s\n", sum, time.Since(now))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
