package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

type Direction struct {
	ch    rune
	xOff  int
	yOff  int
	right *Direction
}

type Guard struct {
	dir *Direction
	x   int
	y   int
}

var (
	UP         = Direction{ch: '^', xOff: 0, yOff: -1}
	LEFT       = Direction{ch: '<', xOff: -1, yOff: 0, right: &UP}
	DOWN       = Direction{ch: 'v', xOff: 0, yOff: 1, right: &LEFT}
	RIGHT      = Direction{ch: '>', xOff: 1, yOff: 0, right: &DOWN}
	DIRECTIONS = []*Direction{&UP, &DOWN, &LEFT, &RIGHT}
)

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

func initializeGuard(grid [][]rune) Guard {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '.' && grid[y][x] != '#' {
				for _, dir := range DIRECTIONS {
					if grid[y][x] == dir.ch {
						return Guard{
							dir: dir,
							y:   y,
							x:   x,
						}
					}
				}
			}
		}
	}
	panic("Could not initialize guard.")
}

func inBounds(x, y int, grid [][]rune) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y])
}

func isWall(x, y int, grid [][]rune) bool {
	return grid[y][x] == '#'
}

func countObstructionsMakingLoops(guard Guard, grid [][]rune, steps [][2]int) int {
	count := 0

	initGuard := guard
	for _, step := range steps {
		x, y := step[0], step[1]
		r := grid[y][x]
		if r != '.' {
			continue
		}

		grid[y][x] = '#'
		if isLoop(guard, grid) {
			count++
		}
		grid[y][x] = r
		guard = initGuard
	}
	return count
}

func isLoop(guard Guard, grid [][]rune) bool {
	walls := [][3]int{}

	for inBounds(guard.x, guard.y, grid) {
		newX := guard.x + guard.dir.xOff
		newY := guard.y + guard.dir.yOff

		if !inBounds(newX, newY, grid) {
			break
		}

		if isWall(newX, newY, grid) {
			guard.dir = guard.dir.right
			wall := [3]int{newX, newY, int(guard.dir.ch)}
			if slices.Contains(walls, wall) {
				return true
			}
			walls = append(walls, wall)
			continue
		}

		guard.x = newX
		guard.y = newY
	}

	return false
}

func findSteps(guard Guard, grid [][]rune) [][2]int {
	steps := [][2]int{}

	for inBounds(guard.x, guard.y, grid) {
		newX := guard.x + guard.dir.xOff
		newY := guard.y + guard.dir.yOff

		if !inBounds(newX, newY, grid) {
			break
		}

		if isWall(newX, newY, grid) {
			guard.dir = guard.dir.right
			continue
		}

		if !slices.Contains(steps, [2]int{newX, newY}) {
			steps = append(steps, [2]int{newX, newY})
		}
		guard.x = newX
		guard.y = newY
	}

	return steps
}

func partOne(input string) {
	grid := readInput(input)
	guard := initializeGuard(grid)
	steps := findSteps(guard, grid)
	fmt.Printf("Output: %d\n", len(steps))
}

func partTwo(input string) {
	now := time.Now()
	grid := readInput(input)
	guard := initializeGuard(grid)
	steps := findSteps(guard, grid)
	cnt := countObstructionsMakingLoops(guard, grid, steps)
	fmt.Printf("Output: %d, Took: %d ms\n", cnt, time.Since(now).Milliseconds())
}

func main() {
	UP.right = &RIGHT
	partOne("input.txt")
	partTwo("input.txt")
}
