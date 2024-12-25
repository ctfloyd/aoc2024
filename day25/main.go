package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeLock(lines []string) []int {
	lock := []int{}
	for i := 0; i < len(lines[0]); i++ {
		height := -1
		for _, l := range lines {
			if l[i] == '#' {
				height++
			}
		}
		lock = append(lock, height)
	}
	return lock
}

func makeKey(lines []string) []int {
	key := []int{}
	for i := 0; i < len(lines[0]); i++ {
		height := -1
		for _, l := range lines {
			if l[i] == '#' {
				height++
			}
		}
		key = append(key, height)
	}
	return key
}

func readInput(input string) (locks, keys [][]int) {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if lines[0][0] == '#' {
				locks = append(locks, makeLock(lines))
			} else {
				keys = append(keys, makeKey(lines))
			}
			lines = []string{}
		} else {
			lines = append(lines, line)
		}
	}

	if lines[0][0] == '#' {
		locks = append(locks, makeLock(lines))
	} else {
		keys = append(keys, makeKey(lines))
	}

	return locks, keys
}

func doesFit(key, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if key[i]+lock[i] >= 6 {
			return false
		}
	}
	return true
}

func partOne(input string) {
	locks, keys := readInput(input)
	fmt.Printf("keys %v\n", keys)

	sum := 0
	for _, lock := range locks {
		for _, key := range keys {
			fmt.Printf("lock %v, key %v (%v)\n", lock, key, doesFit(key, lock))
			if doesFit(key, lock) {
				sum++
			}
		}
	}
	fmt.Printf("result %v\n", sum)
}

func partTwo(input string) {
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
