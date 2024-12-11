package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func innerIsSafe(levels []int, skipIdx int) bool {
	decreasing := true
	safe := true
	first := true
	for i := 0; i < len(levels)-1; i++ {
		currentIdx := i
		nextIdx := i + 1
		if i == skipIdx {
			currentIdx += 1
			nextIdx += 1
			i += 1
		}

		if i+1 == skipIdx {
			nextIdx += 1
			i += 1
		}

		if nextIdx >= len(levels) {
			break
		}

		l1 := levels[currentIdx]
		l2 := levels[nextIdx]

		if first && l2 > l1 {
			decreasing = false
		}

		if decreasing && l2 > l1 {
			safe = false
		}

		if !decreasing && l2 < l1 {
			safe = false
		}

		abs := int(math.Abs(float64(l1 - l2)))
		if abs < 1 || abs > 3 {
			safe = false
		}
		first = false
	}
	return safe
}

func isSafe(levels []int, tolerance bool) bool {
	safe := innerIsSafe(levels, -1)
	if safe || !tolerance {
		return safe
	}

	for j := 0; j < len(levels); j++ {
		safe = innerIsSafe(levels, j)
		if safe {
			fmt.Println(fmt.Sprintf("Levels %d, Remove: %d", levels, j))
			return true
		}
	}
	return false
}

func solve(input string, tolerance bool) {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	sum := 0
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, " ")

		var iFields []int
		for _, n := range fields {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			iFields = append(iFields, i)
		}

		if isSafe(iFields, tolerance) {
			sum++
		}
	}

	fmt.Println(fmt.Sprintf("output: %d", sum))

}

func partOne(input string) {
	solve(input, false)
}

func partTwo(input string) {
	solve(input, true)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
