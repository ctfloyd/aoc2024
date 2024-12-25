package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"
)

func readInput(input string) [][2]int {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)
	r, _ := regexp.Compile(`\d+`)
	i := [][2]int{}
	for scanner.Scan() {
		matches := r.FindAllString(scanner.Text(), -1)
		i1, _ := strconv.Atoi(matches[0])
		i2, _ := strconv.Atoi(matches[1])
		i = append(i, [2]int{i1, i2})
	}
	return i
}

func dji(start, end [2]int, walls [][2]int) int {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	heap := [][3]int{{start[0], start[1], 0}}
	visited := [][2]int{}

	for len(heap) > 0 {
		node := heap[0]
		heap = heap[1:]

		if slices.Contains(visited, [2]int{node[0], node[1]}) {
			continue
		}

		if node[0] == end[0] && node[1] == end[1] {
			return node[2]
		}

		for _, d := range dirs {
			nextNode := [2]int{node[0] + d[0], node[1] + d[1]}
			if nextNode[0] >= 0 && nextNode[0] <= end[0] && nextNode[1] >= 0 && nextNode[1] <= end[1] {
				if !slices.Contains(walls, nextNode) && !slices.Contains(visited, [2]int{nextNode[0], nextNode[1]}) {
					heap = append(heap, [3]int{nextNode[0], nextNode[1], node[2] + 1})
				}
			}
		}
		visited = append(visited, [2]int{node[0], node[1]})
	}
	return -1
}

func solve(walls [][2]int, dim int) int {
	return dji([2]int{0, 0}, [2]int{dim, dim}, walls)
}

func partOne(input string) {
	walls := readInput(input)
	start := time.Now()
	fmt.Printf("result: %v %s\n", solve(walls[:1024], 70), time.Since(start))
}

func partTwo(input string) {
	walls := readInput(input)

	start := time.Now()
	low := 1024
	high := len(walls)
	for low <= high {
		mid := (low + high) / 2

		if solve(walls[:mid], 70) != -1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Printf("result %v, %s\n", walls[high], time.Since(start))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
