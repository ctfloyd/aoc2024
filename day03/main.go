package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readInput(input string) string {
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
	var fullText string
	for scanner.Scan() {
		fullText += scanner.Text()
	}

	return fullText
}

func partOne(input string) {
	fullText := readInput(input)
	pattern, err := regexp.Compile("mul\\((\\d+),(\\d+)\\)")
	if err != nil {
		panic(err)
	}

	result := 0

	matches := pattern.FindAllStringSubmatch(fullText, -1)
	for _, match := range matches {
		i1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		result += i1 * i2
	}

	fmt.Printf("output: %d\n", result)
}

func partTwo(input string) {
	fullText := readInput(input)
	pattern, err := regexp.Compile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)")
	if err != nil {
		panic(err)
	}

	matches := pattern.FindAllStringSubmatch(fullText, -1)

	do := true
	result := 0
	for _, match := range matches {
		if match[0] == "don't()" {
			do = false
		} else if match[0] == "do()" {
			do = true
		} else {
			i1, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			i2, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			if do {
				result += i1 * i2
			}
		}
	}

	fmt.Printf("output: %d\n", result)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
