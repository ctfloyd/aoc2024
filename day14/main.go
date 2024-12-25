package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Robot struct {
	x, y, vx, vy int
}

func readInput(input string) []Robot {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fi)

	r, _ := regexp.Compile(`-*\d+`)
	robots := []Robot{}
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindAllString(line, -1)
		x, _ := strconv.Atoi(m[0])
		y, _ := strconv.Atoi(m[1])
		vx, _ := strconv.Atoi(m[2])
		vy, _ := strconv.Atoi(m[3])
		robots = append(robots, Robot{x, y, vx, vy})
	}
	return robots
}

func (r *Robot) step(w, h int) {
	x := r.x + r.vx
	y := r.y + r.vy

	if x < 0 {
		x = w + x
	}

	if y < 0 {
		y = h + y
	}

	if x >= w {
		x = x - w
	}

	if y >= h {
		y = y - h
	}

	r.x = x
	r.y = y
}

func quads(robots []Robot, w, h int) (int, int, int, int) {
	mw, mh := w/2, h/2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for i := 0; i < len(robots); i++ {
		r := robots[i]
		if r.x != mw && r.y != mh {
			if r.x < mw {
				if r.y < mh {
					q1++
				} else {
					q3++
				}
			} else {
				if r.y < mh {
					q2++
				} else {
					q4++
				}
			}
		}
	}

	return q1, q2, q3, q4
}

func pg(robots []Robot, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			cnt := 0
			for _, r := range robots {
				if r.x == x && r.y == y {
					cnt++
				}
			}
			if cnt == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", cnt)
			}
		}
		fmt.Println()
	}
}

func isPossibleTree(robots []Robot) bool {
	m := [][2]int{}
	for _, r := range robots {
		pos := [2]int{r.x, r.y}
		if !slices.Contains(m, pos) {
			m = append(m, [2]int{r.x, r.y})
		} else {
			return false
		}
	}
	return true
}

func solve(robots []Robot, w, h, steps int) int {
	for i := 0; i < steps; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].step(w, h)
		}
		if isPossibleTree(robots) {
			pg(robots, w, h)
			return i + 1
		}
	}
	q1, q2, q3, q4 := quads(robots, w, h)
	return q1 * q2 * q3 * q4
}

func partOne(input string) {
	robots := readInput(input)
	value := solve(robots, 101, 103, 100)
	fmt.Printf("value %v\n", value)
}

func partTwo(input string) {
	robots := readInput(input)
	value := solve(robots, 101, 103, 0xFFFFFFFF)
	fmt.Printf("value %v\n", value)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
