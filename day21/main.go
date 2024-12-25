package main

import (
	"fmt"
	"math"
	"strconv"
)

var numpad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'Z', '0', 'A'},
}

var dirpad = [][]rune{
	{'Z', '^', 'A'},
	{'<', 'v', '>'},
}

type Dir struct {
	nr, nc int
	ch     string
}

var (
	LEFT  = Dir{0, -1, "<"}
	RIGHT = Dir{0, 1, ">"}
	UP    = Dir{-1, 0, "^"}
	DOWN  = Dir{1, 0, "v"}
	DIRS  = []Dir{LEFT, RIGHT, UP, DOWN}
)

type Node struct {
	r, c int
	path string
}

type CacheKey struct {
	target string
	depth  int
}

var cache = make(map[CacheKey]int64)

func expand(current, next rune) []string {
	keypad := dirpad
	_, err := strconv.Atoi(string(current))
	_, err2 := strconv.Atoi(string(next))
	if err == nil || err2 == nil {
		keypad = numpad
	}

	start := [2]int{}
	for row := 0; row < len(keypad); row++ {
		for col := 0; col < len(keypad[row]); col++ {
			if keypad[row][col] == current {
				start = [2]int{row, col}
				break
			}
		}
	}

	optimals := []string{}
	best := math.MaxInt32
	q := []Node{{start[0], start[1], ""}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if keypad[n.r][n.c] == 'Z' {
			continue
		}

		if len(n.path) >= best {
			break
		}

		if keypad[n.r][n.c] == next {
			optimals = append(optimals, n.path+"A")
			best = len(n.path) + 1
			continue
		}

		for _, dir := range DIRS {
			nn := [2]int{n.r + dir.nr, n.c + dir.nc}
			if nn[0] >= 0 && nn[1] >= 0 && nn[0] < len(keypad) && nn[1] < len(keypad[n.r]) {
				q = append(q, Node{nn[0], nn[1], n.path + dir.ch})
			}
		}
	}
	return optimals

}

func getMoveCount(current, next rune, depth int) int64 {
	if current == next {
		return 1
	}

	expanded := expand(current, next)
	var best int64 = math.MaxInt64
	for _, exp := range expanded {
		best = int64(math.Min(float64(best), float64(getSequenceLength(exp, depth-1))))
	}
	return best
}

func getSequenceLength(target string, depth int) int64 {
	if v, ok := cache[CacheKey{target, depth}]; ok {
		return v
	}

	var length int64
	if depth == 0 {
		length = int64(len(target))
	} else {
		current := 'A'
		for _, next := range target {
			length += getMoveCount(current, next, depth)
			current = next
		}
	}

	cache[CacheKey{target, depth}] = length
	return length
}

func findNumber(in string) int {
	for in[0] == '0' {
		in = in[1:]
	}
	i, _ := strconv.Atoi(in[:len(in)-1])
	return i
}

func partOne(input []string) {
	var sum int64 = 0
	for _, in := range input {
		len := getSequenceLength(in, 3)
		code := findNumber(in)
		fmt.Printf("l %v, code %v\n", len, code)
		sum += len * int64(code)
	}
	fmt.Printf("result %v\n", sum)
}

func partTwo(input []string) {
	var sum int64 = 0
	for _, in := range input {
		len := getSequenceLength(in, 26)
		code := findNumber(in)
		sum += len * int64(code)
	}
	fmt.Printf("result %v\n", sum)
}

func main() {
	// example := []string{"029A", "980A", "179A", "456A", "379A"}
	input := []string{"964A", "246A", "973A", "682A", "180A"}
	partOne(input)
	partTwo(input)
}
