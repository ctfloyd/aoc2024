package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
)

type Antenna struct {
	frequency rune
	row       int
	col       int
}

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

func parseGrid(grid [][]rune) []Antenna {
	antennas := []Antenna{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != '.' {
				antennas = append(antennas, Antenna{
					frequency: grid[row][col],
					row:       row,
					col:       col,
				})
			}
		}
	}
	return antennas
}

func groupByFreq(antennas []Antenna) map[rune][]Antenna {
	m := make(map[rune][]Antenna)
	for _, a := range antennas {
		m[a.frequency] = append(m[a.frequency], a)
	}
	return m
}

func isInRange(pos [2]int, rows, cols int) bool {
	return pos[0] >= 0 && pos[0] < rows && pos[1] >= 0 && pos[1] < cols
}

func isAnyInRange(an [][2]int, rows, cols int) bool {
	r := false
	for _, a := range an {
		r = r || isInRange(a, rows, cols)
	}
	return r
}

func calculateAntinodes(a, b Antenna, rows, cols int, ignoreDistance bool) [][2]int {
	ans := [][2]int{}

	rowDistance := int(math.Abs(float64(a.row - b.row)))
	colDistance := int(math.Abs(float64(a.col - b.col)))

	if a.row > b.row {
		tmp := b
		b = a
		a = tmp
	}

	if a.col < b.col {
		if ignoreDistance {
			i := 1
			finished := false
			for !finished {
				answers := [][2]int{
					{a.row - rowDistance*i, a.col - colDistance*i},
					{a.row + rowDistance*i, a.col + colDistance*i},
					{b.row - rowDistance*i, b.col - colDistance*i},
					{b.row + rowDistance*i, b.col + colDistance*i},
				}
				ans = append(ans, answers...)
				finished = !isAnyInRange(answers, rows, cols)
				i += 1
			}
		} else {
			ans = append(ans, [2]int{a.row - rowDistance, a.col - colDistance})
			ans = append(ans, [2]int{b.row + rowDistance, b.col + colDistance})
		}
	} else {
		if ignoreDistance {
			i := 1
			finished := false
			for !finished {
				answers := [][2]int{
					{a.row - rowDistance*i, a.col + colDistance*i},
					{a.row + rowDistance*i, a.col - colDistance*i},
					{b.row - rowDistance*i, b.col + colDistance*i},
					{b.row + rowDistance*i, b.col - colDistance*i},
				}
				ans = append(ans, answers...)
				finished = !isAnyInRange(answers, rows, cols)
				i += 1
			}
		} else {
			ans = append(ans, [2]int{a.row - rowDistance, a.col + colDistance})
			ans = append(ans, [2]int{b.row + rowDistance, b.col - colDistance})
		}
	}

	fmt.Printf("For antennas %v, %v, found ans %v\n", a, b, ans)

	return ans
}

func findAntinodes(a, b Antenna, rows, cols int, ignoreDistance bool) [][2]int {
	ans := [][2]int{}
	for _, an := range calculateAntinodes(a, b, rows, cols, ignoreDistance) {
		if isInRange(an, rows, cols) {
			ans = append(ans, an)
		}
	}
	return ans
}

func addAntinodes(allAntinodes [][2]int, antinodes [][2]int) [][2]int {
	for _, an := range antinodes {
		if !slices.Contains(allAntinodes, an) {
			allAntinodes = append(allAntinodes, an)
		}
	}
	return allAntinodes
}

func solveForFrequency(antennas []Antenna, rows, cols int, ignoreDistance bool) [][2]int {
	n := len(antennas)

	fmt.Printf("freq %s\n", string(antennas[0].frequency))
	allAntinodes := [][2]int{}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			a, b := antennas[i], antennas[j]
			if a != b {
				antinodes := findAntinodes(a, b, rows, cols, ignoreDistance)
				allAntinodes = addAntinodes(allAntinodes, antinodes)
			}
		}
	}
	return allAntinodes
}

func solve(antennasByFreq map[rune][]Antenna, rows, cols int, ignoreDistance bool) [][2]int {
	ans := [][2]int{}

	k := maps.Keys(antennasByFreq)
	for key := range k {
		ans = addAntinodes(ans, solveForFrequency(antennasByFreq[key], rows, cols, ignoreDistance))
	}
	return ans
}

func printGrid(grid [][]rune, result [][2]int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '.' && slices.Contains(result, [2]int{row, col}) {
				fmt.Print("#")
			} else {
				fmt.Print(string(grid[row][col]))
			}
		}
		fmt.Println()
	}
}

func partOne(input string) {
	grid := readInput(input)
	antennas := parseGrid(grid)
	antennasByFreq := groupByFreq(antennas)
	result := solve(antennasByFreq, len(grid), len(grid[0]), false)
	fmt.Printf("result %v\n", len(result))
}

func partTwo(input string) {
	grid := readInput(input)
	antennas := parseGrid(grid)
	antennasByFreq := groupByFreq(antennas)
	result := solve(antennasByFreq, len(grid), len(grid[0]), true)
	printGrid(grid, result)
	fmt.Printf("result %v\n", len(result))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
