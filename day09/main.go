package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func readInput(input string) []int {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fi)
	scanner.Scan()

	out := []int{}
	for _, ch := range scanner.Text() {
		out = append(out, int(ch-'0'))
	}
	return out
}

func makeBlocks(disk []int) []int {
	blocks := []int{}

	if len(disk)%2 != 0 {
		disk = append(disk, 0)
	}

	id := 0
	for i := 0; i < len(disk); i += 2 {
		len := disk[i]
		fs := disk[i+1]
		for j := 0; j < len; j++ {
			blocks = append(blocks, id)
		}
		for j := 0; j < fs; j++ {
			blocks = append(blocks, -1)
		}
		id += 1
	}
	return blocks
}

func compactBlocks(blocks []int) []int {
	bcopy := make([]int, len(blocks))
	copy(bcopy, blocks)

	for i := len(blocks) - 1; i >= len(blocks)/2; i-- {
		block := blocks[i]
		if block != -1 {
			bcopy[i] = -1
			for j := 0; j < len(bcopy); j++ {
				if bcopy[j] == -1 {
					bcopy[j] = block
					break
				}
			}
		}
	}
	return bcopy
}

type Empty struct {
	start int
	count int
}

func findEmpties(blocks []int) []Empty {
	empties := []Empty{}

	start := -1
	count := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == -1 {
			if start == -1 {
				start = i
				count = 1
			} else {
				count++
			}
		} else {
			if start > -1 {
				empties = append(empties, Empty{start, count})
			}
			start = -1
			count = 0
		}
	}
	return empties
}

func compactFiles(blocks []int) []int {
	bcopy := make([]int, len(blocks))
	copy(bcopy, blocks)

	fileStartIndex := len(blocks) - 1
	lastFileId := blocks[fileStartIndex]

	empties := findEmpties(blocks)

	for i := len(blocks) - 1; i >= 0; i-- {
		fileId := blocks[i]
		// We found the end of the file iterating backwards.
		if fileId != lastFileId {
			// If the last run of files we were checking was actually a file and not empty space.
			if lastFileId != -1 {
				// The size of the file.
				size := fileStartIndex - i
				for i, empty := range empties {
					if empty.count >= size && empty.start < fileStartIndex {
						// Place the file in the empty space.
						for j := 0; j < size; j++ {
							bcopy[empty.start+j] = lastFileId
						}
						// Remove the file from its original place.
						for j := 0; j < size; j++ {
							bcopy[fileStartIndex-j] = -1
						}

						empties[i] = Empty{start: empty.start + size, count: empty.count - size}
						break
					}
				}
			}

			// The fild id changed, start tracking the next file id.
			fileStartIndex = i
			lastFileId = fileId
		}
	}
	return bcopy

}

func checksum(compact []int) int {
	sum := 0
	for idx, i := range compact {
		if i != -1 {
			sum += idx * i
		}
	}
	return sum
}

func partOne(input string) {
	fmt.Printf("result: %v\n", checksum(compactBlocks(makeBlocks(readInput(input)))))
}

func partTwo(input string) {
	now := time.Now()
	fmt.Printf("result: %v, %s\n", checksum(compactFiles(makeBlocks(readInput(input)))), time.Since(now))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
