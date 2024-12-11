package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ordering struct {
	before int
	after  int
}

type Update struct {
	updates []int
}

func readInput(input string) ([]Ordering, []Update) {
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

	orderings := []Ordering{}
	updates := []Update{}

	orderingState := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			orderingState = false
		} else if orderingState {
			parts := strings.Split(line, "|")
			l, lerr := strconv.Atoi(parts[0])
			if lerr != nil {
				panic(lerr)
			}
			r, rerr := strconv.Atoi(parts[1])
			if rerr != nil {
				panic(rerr)
			}
			orderings = append(orderings, Ordering{before: l, after: r})
		} else {
			update := []int{}
			parts := strings.Split(line, ",")
			for _, p := range parts {
				n, nerr := strconv.Atoi(p)
				if nerr != nil {
					panic(nerr)
				}
				update = append(update, n)
			}
			updates = append(updates, Update{updates: update})
		}
	}

	return orderings, updates
}

func findMiddleNumber(update Update) int {
	nums := update.updates
	return nums[len(nums)/2]
}

func indexOfSlice(slice []int, value int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return i
		}
	}

	return -1
}

func indexOf(update Update, value int) int {
	return indexOfSlice(update.updates, value)
}

func isOrderingViolated(update Update, ordering Ordering) bool {
	beforePageIndex := indexOf(update, ordering.before)
	afterPageIndex := indexOf(update, ordering.after)
	if beforePageIndex > -1 && afterPageIndex > -1 {
		if beforePageIndex > afterPageIndex {
			return true
		}
	}
	return false
}

func isCorrectlyOrdered(update Update, orderings []Ordering) bool {
	for _, ordering := range orderings {
		if isOrderingViolated(update, ordering) {
			return false
		}
	}

	return true
}

func swap(slice []int, i, j int) []int {
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
	return slice
}

func fixOrdering(update Update, orderings []Ordering) Update {
	for !isCorrectlyOrdered(update, orderings) {
		for _, ordering := range orderings {
			if isOrderingViolated(update, ordering) {
				update.updates = swap(update.updates, indexOf(update, ordering.before), indexOf(update, ordering.after))
			}
		}

	}

	return update
}

func partOne(input string) {
	orderings, updates := readInput(input)

	sum := 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, orderings) {
			sum += findMiddleNumber(update)
		}
	}

	fmt.Printf("Output: %d\n", sum)
}

func partTwo(input string) {
	orderings, updates := readInput(input)
	sum := 0
	for _, update := range updates {
		if !isCorrectlyOrdered(update, orderings) {
			correctOrdering := fixOrdering(update, orderings)
			sum += findMiddleNumber(correctOrdering)
		}
	}

	fmt.Printf("Output: %d\n", sum)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
