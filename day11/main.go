package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(input string) []int {
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

	stones := []int{}
	for scanner.Scan() {
		lineString := scanner.Text()

		strStones := strings.Split(lineString, " ")

		for _, stone := range strStones {
			i, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			stones = append(stones, i)
		}
	}

	return stones
}

func blinkImpl(st map[int]int) map[int]int {
	stones := make(map[int]int)
	for stone := range maps.Keys(st) {
		count := st[stone]

		strStone := strconv.Itoa(stone)
		n := len(strStone)
		if stone == 0 {
			stones[1] += count
		} else if n%2 == 0 {
			left, err := strconv.Atoi(strStone[:n/2])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(strStone[n/2:])
			if err != nil {
				panic(err)
			}
			stones[left] += count
			stones[right] += count
		} else {
			stones[stone*2024] += count
		}
	}
	return stones
}

func blink(stonesI []int, iterations int) map[int]int {
	stones := make(map[int]int)
	for _, st := range stonesI {
		stones[st]++
	}

	for i := 0; i < iterations; i++ {
		stones = blinkImpl(stones)
	}
	return stones
}

func sum(stones map[int]int) int {
	sum := 0
	for k := range maps.Keys(stones) {
		sum += stones[k]
	}
	return sum
}

func partOne(input string) {
	stones := readInput(input)
	stonesMap := blink(stones, 25)
	fmt.Printf("Output %v\n", sum(stonesMap))
}

func partTwo(input string) {
	now := time.Now()
	stones := readInput(input)
	stonesMap := blink(stones, 75)
	fmt.Printf("Output %v %s\n", sum(stonesMap), time.Since(now))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
