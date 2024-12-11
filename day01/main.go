package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math"
	"sort"
	"strconv"
)

func partOne() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	var left []int
	var right []int

	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, "   ")

		li, lerr := strconv.Atoi(fields[0])
		ri, rerr := strconv.Atoi(fields[1])

		if lerr != nil {
			panic(lerr)
		}

		if rerr != nil {
			panic(rerr)
		}


		left = append(left, li)
		right = append(right, ri)
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("output: %d", sum));

}

func partTwo() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	var left []int
	right := make(map[int]int)

	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, "   ")

		li, lerr := strconv.Atoi(fields[0])
		ri, rerr := strconv.Atoi(fields[1])

		if lerr != nil {
			panic(lerr)
		}

		if rerr != nil {
			panic(rerr)
		}


		left = append(left, li)
		right[ri]++
	}

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += left[i] * right[left[i]]
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("output2: %d", sum));

}

func main()  {
	partOne()
	partTwo()
}
