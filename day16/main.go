package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
)

type Direction struct {
	oy, ox, idx int
}

var (
	NORTH = Direction{oy: -1, ox: 0, idx: 0}
	EAST  = Direction{oy: 0, ox: 1, idx: 1}
	SOUTH = Direction{oy: 1, ox: 0, idx: 2}
	WEST  = Direction{oy: 0, ox: -1, idx: 3}
	DIRS  = []Direction{NORTH, EAST, WEST, SOUTH}
)

func opposite(d1, d2 int) bool {
	return d1 == NORTH.idx && d2 == SOUTH.idx ||
		d1 == SOUTH.idx && d2 == NORTH.idx ||
		d1 == EAST.idx && d2 == WEST.idx ||
		d1 == WEST.idx && d2 == EAST.idx
}

func readInput(input string) [][]rune {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func findStartAndEnd(grid [][]rune) ([2]int, [2]int) {
	start, end := [2]int{-1, -1}, [2]int{-1, -1}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'S' {
				start = [2]int{row, col}
			} else if grid[row][col] == 'E' {
				end = [2]int{row, col}
			}
		}
	}
	return start, end
}

type Node struct {
	row, col, dir, score int
	path                 [][2]int
}

func dji(start, end [2]int, dir Direction, grid [][]rune) (int, [][][2]int) {
	scores := make(map[int][][][2]int)

	h := []Node{{row: start[0], col: start[1], dir: dir.idx, score: 0, path: [][2]int{{start[0], start[1]}}}}
	v := [][3]int{}

	for len(h) > 0 {
		n := h[0]
		h = h[1:]
		slices.SortFunc(h, func(a, b Node) int {
			return a.score - b.score
		})

		if n.row == end[0] && n.col == end[1] {
			scores[n.score] = append(scores[n.score], n.path)
		}

		for _, d := range DIRS {
			nn := [2]int{n.row + d.oy, n.col + d.ox}
			if grid[nn[0]][nn[1]] != '#' && !opposite(d.idx, n.dir) {
				c := 1
				if d.idx != n.dir {
					c += 1000
				}
				if !slices.Contains(v, [3]int{nn[0], nn[1], d.idx}) {
					pa := append([][2]int{}, append(n.path, nn)...)
					h = append(h, Node{row: nn[0], col: nn[1], dir: d.idx, score: n.score + c, path: pa})
				}
			}
		}
		k := [3]int{n.row, n.col, n.dir}
		v = append(v, k)
	}

	keys := slices.Collect(maps.Keys(scores))
	sort.Ints(keys)

	return keys[0], scores[keys[0]]
}

func score(grid [][]rune) (int, [][][2]int) {
	start, end := findStartAndEnd(grid)
	score, paths := dji(start, end, EAST, grid)
	return score, paths
}

func countBestTiles(paths [][][2]int) int {
	t := [][2]int{}
	for _, p := range paths {
		for _, ti := range p {
			if !slices.Contains(t, ti) {
				t = append(t, ti)
			}
		}
	}
	return len(t)
}

func partOne(input string) {
	grid := readInput(input)
	score, _ := score(grid)
	fmt.Printf("result %v\n", score)
}

func partTwo(input string) {
	grid := readInput(input)
	_, paths := score(grid)
	tiles := countBestTiles(paths)
	fmt.Printf("result %v\n", tiles)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
