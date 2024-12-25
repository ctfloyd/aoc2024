package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(input string) []int {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)
	nums := []int{}
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, i)
	}
	return nums
}

func generateSequence(num, iterations int) []int {
	nums := []int{}
	for range iterations {
		num = (num ^ (num << 6)) % 16777216
		num = (num ^ (num >> 5)) % 16777216
		num = (num ^ (num << 11)) % 16777216
		nums = append(nums, num)
	}
	return nums
}
func getOnes(seq []int) []int {
	ones := []int{}
	for _, num := range seq {
		ones = append(ones, num%10)
	}
	return ones
}

func getAllOnes(seqs [][]int) [][]int {
	allOnes := [][]int{}
	for _, seq := range seqs {
		allOnes = append(allOnes, getOnes(seq))
	}
	return allOnes
}

func getAllWindows(seqs [][]int) [][4]int {
	windows := [][4]int{}
	m := make(map[[4]int]bool)
	for _, seq := range seqs {
		for i := 1; i < len(seq)-3; i += 1 {
			window := [4]int{seq[i] - seq[i-1], seq[i+1] - seq[i], seq[i+2] - seq[i+1], seq[i+3] - seq[i+2]}
			if _, ok := m[window]; !ok {
				windows = append(windows, window)
				m[window] = true
			}
		}
	}
	return windows
}

type M map[[4]int]int

func calculateMaxBananas(windows [][4]int, ones [][]int) int {
	ml := make([]M, len(ones))
	for j, seq := range ones {
		m := M{}
		for i := 1; i < len(seq)-3; i += 1 {
			sw := [4]int{seq[i] - seq[i-1], seq[i+1] - seq[i], seq[i+2] - seq[i+1], seq[i+3] - seq[i+2]}
			if _, ok := m[sw]; !ok {
				m[sw] = seq[i+3]
			}
		}
		ml[j] = m
	}

	maxBananas := -1
	for _, window := range windows {
		bananas := 0
		for i := 0; i < len(ones); i++ {
			bananas += ml[i][window]
		}
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}
	return maxBananas
}

func partOne(input string) {
	nums := readInput(input)
	sum := 0
	for _, n := range nums {
		res := generateSequence(n, 2000)
		sum += res[len(res)-1]
	}
	fmt.Printf("result %v\n", sum)
}

func partTwo(input string) {
	nums := readInput(input)
	allSeqs := [][]int{}
	for _, n := range nums {
		seq := []int{n}
		seq = append(seq, generateSequence(n, 2000)...)
		allSeqs = append(allSeqs, seq)
	}
	ones := getAllOnes(allSeqs)
	windows := getAllWindows(ones)
	maxBananas := calculateMaxBananas(windows, ones)
	fmt.Printf("result %v\n", maxBananas)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
