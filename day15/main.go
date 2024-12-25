package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Position struct {
	row, col int
}

type Wall struct {
	position Position
}

type Box struct {
	width    int
	position Position
}

func (box *Box) isPushable(movement rune, warehouse *Warehouse) bool {
	ro, co := offset(movement)
	pushable := true

	if ro != 0 {
		if !warehouse.isSpanEmpty(box.position.row+ro, box.position.col, box.position.col+(box.width-1)) {
			// If moving vertically, check there are no walls for the box's span.
			// And that other boxes in its way are also pushable.
			for i := 0; i < box.width; i++ {
				wall := warehouse.getWall(box.position.row+ro, box.position.col+i)
				if wall != nil {
					pushable = false
				}
				b := warehouse.getBox(box.position.row+ro, box.position.col+i)
				if b != nil {
					pushable = pushable && b.isPushable(movement, warehouse)
				}
			}
		}
	} else {
		col := 0
		if co < 0 {
			col = box.position.col - 1
		} else {
			col = box.position.col + box.width
		}
		// Check there are no walls in the one horizontal direction.
		b := warehouse.getBox(box.position.row, col)
		wall := warehouse.getWall(box.position.row, col)
		if wall != nil {
			pushable = false
		} else if b != nil {
			pushable = pushable && b.isPushable(movement, warehouse)
		}
	}

	return pushable
}

func (box *Box) push(movement rune, warehouse *Warehouse) {
	ro, co := offset(movement)
	if ro != 0 {
		if !warehouse.isSpanEmpty(box.position.row+ro, box.position.col, box.position.col+(box.width-1)) {
			for i := 0; i < box.width; i++ {
				box := warehouse.getBox(box.position.row+ro, box.position.col+i)
				if box != nil {
					box.push(movement, warehouse)
				}
			}
		}
		box.position.row += ro
	}

	if co != 0 {
		if co > 0 {
			if !warehouse.isEmpty(box.position.row, box.position.col+box.width) {
				box := warehouse.getBox(box.position.row, box.position.col+box.width)
				box.push(movement, warehouse)
			}
		} else {
			if !warehouse.isEmpty(box.position.row, box.position.col-1) {
				box := warehouse.getBox(box.position.row, box.position.col-1)
				box.push(movement, warehouse)
			}
		}
		box.position.col += co
	}

}

type Robot struct {
	position Position
}

type Warehouse struct {
	walls []Wall
	boxes []Box
}

func (warehouse *Warehouse) isSpanEmpty(row, startCol, endCol int) bool {
	for i := 0; i <= (endCol - startCol); i++ {
		wall := warehouse.getWall(row, startCol+i)
		box := warehouse.getBox(row, startCol+i)
		if wall != nil || box != nil {
			return false
		}
	}
	return true
}

func (warehouse *Warehouse) isEmpty(row, col int) bool {
	wall := warehouse.getWall(row, col)
	box := warehouse.getBox(row, col)
	if wall != nil || box != nil {
		return false
	}
	return true
}

func (warehouse *Warehouse) getWall(row, col int) *Wall {
	for i := range warehouse.walls {
		w := &warehouse.walls[i]
		if w.position.row == row && w.position.col == col {
			return w
		}
	}
	return nil
}

func (warehouse *Warehouse) getBox(row, col int) *Box {
	for i := range warehouse.boxes {
		b := &warehouse.boxes[i]
		if b.position.row == row && (col >= b.position.col && col <= b.position.col+b.width-1) {
			return b
		}
	}
	return nil

}

func (robot *Robot) move(movement rune, warehouse *Warehouse) {
	or, oc := offset(movement)

	wall := warehouse.getWall(robot.position.row+or, robot.position.col+oc)
	box := warehouse.getBox(robot.position.row+or, robot.position.col+oc)
	if wall != nil && box != nil {
		fmt.Printf("INVALID! Warehouse report %d,%d is both a wall and a box!", robot.position.row+or, robot.position.col+oc)
	}

	if wall != nil {
		return
	}

	if box != nil {
		if !box.isPushable(movement, warehouse) {
			return
		}
		box.push(movement, warehouse)
	}

	robot.position.row += or
	robot.position.col += oc
}

func readInput(input string, expand bool) ([][]rune, []rune) {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	var grid [][]rune
	var movements []rune

	parseMovements := false
	for scanner.Scan() {
		lineString := scanner.Text()
		if len(lineString) == 0 {
			parseMovements = true
			continue
		}
		if parseMovements {
			movements = append(movements, []rune(lineString)...)
		} else {
			line := []rune{}
			if expand {
				for _, ch := range lineString {
					if ch != 'O' && ch != '@' {
						line = append(line, ch)
						line = append(line, ch)
					} else if ch == '@' {
						line = append(line, ch)
						line = append(line, '.')
					} else {
						line = append(line, '[')
						line = append(line, ']')
					}
				}

				grid = append(grid, line)
			} else {
				grid = append(grid, []rune(lineString))
			}
		}
	}

	return grid, movements
}

func parseGrid(grid [][]rune) (Warehouse, Robot) {
	robot := Robot{}
	walls := []Wall{}
	boxes := []Box{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			ch := grid[row][col]
			if ch == '#' {
				walls = append(walls, Wall{position: Position{row, col}})
			} else if ch == '@' {
				robot = Robot{position: Position{row, col}}
			} else if ch == 'O' {
				boxes = append(boxes, Box{width: 1, position: Position{row, col}})
			} else if ch == '[' {
				boxes = append(boxes, Box{width: 2, position: Position{row, col}})
			}
		}
	}
	return Warehouse{walls, boxes}, robot
}

func offset(movement rune) (int, int) {
	if movement == '^' {
		return -1, 0
	} else if movement == '>' {
		return 0, 1
	} else if movement == '<' {
		return 0, -1
	} else if movement == 'v' {
		return 1, 0
	}
	fmt.Printf("invalid movement %v\n", movement)
	return -1, -1
}

func doMoves(robot *Robot, warehouse *Warehouse, movements []rune, w, h int) {
	for _, movement := range movements {
		robot.move(movement, warehouse)
	}
}

func printGrid(w, h int, warehouse *Warehouse, robot *Robot) {
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			wall := warehouse.getWall(row, col)
			if wall != nil {
				fmt.Print("#")
			}
			box := warehouse.getBox(row, col)
			if box != nil {
				if box.width == 2 {
					fmt.Print("[]")
					col++
				} else {
					fmt.Print("O")
				}
			}
			if box == nil && wall == nil {
				if robot.position.row == row && robot.position.col == col {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func score(warehouse *Warehouse) int {
	sum := 0
	for _, box := range warehouse.boxes {
		sum += 100*box.position.row + box.position.col
	}
	return sum
}

func partOne(input string) {
	grid, movements := readInput(input, false)
	warehouse, robot := parseGrid(grid)
	doMoves(&robot, &warehouse, movements, len(grid[0]), len(grid))
	fmt.Printf("result %v\n", score(&warehouse))
}

func partTwo(input string) {
	grid, movements := readInput(input, true)
	warehouse, robot := parseGrid(grid)
	now := time.Now()
	doMoves(&robot, &warehouse, movements, len(grid[0]), len(grid))
	fmt.Printf("result %v (%s ms)\n", score(&warehouse), time.Since(now))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
