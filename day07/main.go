package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = iota
	MULTIPLTY
	CONCAT
)

type Equation struct {
	result int
	parts  []int
}

func readInput(input string) []Equation {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	equations := []Equation{}
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		lineString := scanner.Text()

		colonIndex := strings.Index(lineString, ":")
		result, err := strconv.Atoi(lineString[:colonIndex])
		if err != nil {
			panic(err)
		}

		partsString := lineString[colonIndex+2:]
		partsS := strings.Split(partsString, " ")

		parts := []int{}
		for _, part := range partsS {
			iPart, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			parts = append(parts, iPart)
		}
		equations = append(equations, Equation{result, parts})
	}

	return equations
}

func operate(a, b, operation int) int {
	if operation == MULTIPLTY {
		return a * b
	} else if operation == ADD {
		return a + b
	} else if operation == CONCAT {
		s, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		if err != nil {
			panic(err)
		}
		return s
	} else {
		panic("invalid operation")
	}
}

// Recursively check the solution space.
func canSolve(result int, numbers, operations []int) bool {
	if len(numbers) == 1 {
		return result == numbers[0]
	}

	solvable := false
	for op := range operations {
		a := operate(numbers[0], numbers[1], op)
		n := []int{a}
		n = append(n, numbers[2:]...)

		solvable = solvable || canSolve(result, n, operations)
	}
	return solvable
}

func solve(equations []Equation, operations []int) int {
	sum := 0
	for _, eq := range equations {
		if canSolve(eq.result, eq.parts, operations) {
			sum += eq.result
		}
	}
	return sum
}

func partOne(input string) {
	equations := readInput(input)
	fmt.Printf("Output: %v\n", solve(equations, []int{ADD, MULTIPLTY}))
}

func partTwo(input string) {
	equations := readInput(input)
	fmt.Printf("Output: %v\n", solve(equations, []int{ADD, MULTIPLTY, CONCAT}))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
