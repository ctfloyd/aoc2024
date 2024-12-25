package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
	"strings"
	"time"
)

func readInput(input string) [][2]string {
	fi, _ := os.Open(input)
	scanner := bufio.NewScanner(fi)

	computers := [][2]string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		computers = append(computers, [2]string{parts[0], parts[1]})
	}
	return computers
}

func connectComputers(connections [][2]string) map[string][]string {
	m := make(map[string][]string)
	for _, connection := range connections {
		if !slices.Contains(m[connection[0]], connection[1]) {
			m[connection[0]] = append(m[connection[0]], connection[1])
		}
		if !slices.Contains(m[connection[1]], connection[0]) {
			m[connection[1]] = append(m[connection[1]], connection[0])
		}
	}
	return m
}

func generateCombos(set []string, n int) (subsets [][]string) {
	length := uint(len(set))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		subset := []string{}
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func isValidCombo(combo []string, allConnections map[string][]string) bool {
	valid := areAllComputersConnected([]string(combo), allConnections)
	if valid {
		valid = false
		for _, elem := range combo {
			if elem[0] == 't' {
				valid = true
			}
		}
	}

	return valid
}

func findLanParties(allConnections map[string][]string) [][3]string {
	result := [][3]string{}
	checked := make(map[[3]string]bool)

	for computer, connections := range allConnections {
		conn := []string{computer}
		conn = append(conn, connections...)

		if len(conn) >= 3 {
			combinations := generateCombos(conn, 3)
			for _, combo := range combinations {
				slices.Sort(combo)
				smallCombo := [3]string(combo)
				if !checked[smallCombo] && isValidCombo(combo, allConnections) {
					result = append(result, smallCombo)
					checked[smallCombo] = true
				}
			}
		}
	}

	return result
}

func areAllComputersConnected(combo []string, allConnections map[string][]string) bool {
	for i := 0; i < len(combo); i++ {
		elem := combo[i]
		for j := 0; j < len(combo); j++ {
			if i != j {
				v, ok := allConnections[elem]
				if !ok || !slices.Contains(v, combo[j]) {
					return false
				}
			}
		}
	}
	return true
}

func findLargestLanParty(allConnections map[string][]string) []string {
	result := []string{}
	for computer, connections := range allConnections {
		conn := []string{computer}
		conn = append(conn, connections...)

		if len(conn) < len(result) {
			continue
		}

		for i := len(conn); i >= len(result); i-- {
			combinations := generateCombos(conn, i)
			for _, combo := range combinations {
				slices.Sort(combo)
				if areAllComputersConnected(combo, allConnections) {
					result = combo
				}
			}
		}
	}

	return result
}

func partOne(input string) {
	computers := readInput(input)
	start := time.Now()
	connections := connectComputers(computers)
	parties := findLanParties(connections)
	fmt.Printf("result %v (%s)\n", len(parties), time.Since(start))
}

func partTwo(input string) {
	computers := readInput(input)
	start := time.Now()
	connections := connectComputers(computers)
	party := findLargestLanParty(connections)
	fmt.Printf("result %v (%s) \n", strings.Join(party, ","), time.Since(start))
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
