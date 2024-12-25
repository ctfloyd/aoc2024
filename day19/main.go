package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(input string) ([]string, []string) {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	towels := []string{}
	designs := []string{}

	first := true
	for scanner.Scan() {
		line := scanner.Text()
		if first {
			towels = strings.Split(line, ", ")
			first = false
		} else if line == "" {
			continue
		} else {
			designs = append(designs, line)
		}
	}
	return towels, designs
}

var memo = make(map[[2]string]int)

func isDesignPossibleRecursive(current string, target string, towels []string) int {
	if v, ok := memo[[2]string{current, target}]; ok {
		return v
	}

	if current == target {
		return 1
	}

	ways := 0
	for _, towel := range towels {
		option := current + towel
		if len(option) <= len(target) && option == target[:len(option)] {
			ways += isDesignPossibleRecursive(option, target, towels)
		}
	}
	memo[[2]string{current, target}] = ways
	return ways
}

func isDesignPossible(design string, towels []string) int {
	return isDesignPossibleRecursive("", design, towels)
}

func partOne(input string) {
	towels, designs := readInput(input)
	cnt := 0
	for _, design := range designs {
		if isDesignPossible(design, towels) > 0 {
			cnt += 1
		}
	}
	fmt.Printf("Result: %v\n", cnt)
}

func partTwo(input string) {
	towels, designs := readInput(input)
	cnt := 0
	for _, design := range designs {
		cnt += isDesignPossible(design, towels)
	}
	fmt.Printf("Result: %v\n", cnt)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
