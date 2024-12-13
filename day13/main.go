package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	x, y int
}

type Prize struct {
	x, y int
}

type Claw struct {
	x, y int
}

type Game struct {
	a, b  Button
	prize Prize
}

const SENTINEL = 0xFFFFFFFF

func readInput(input string, offset int) []Game {
	fi, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fi)

	r, _ := regexp.Compile(`\d+`)

	games := []Game{}

	game := Game{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			games = append(games, game)
			game = Game{}
			continue
		}

		m := r.FindAllString(line, -1)
		m0, _ := strconv.Atoi(m[0])
		m1, _ := strconv.Atoi(m[1])

		if game.a.x == 0 {
			game.a = Button{x: m0, y: m1}
		} else if game.b.x == 0 {
			game.b = Button{x: m0, y: m1}
		} else {
			game.prize = Prize{x: m0 + offset, y: m1 + offset}
		}
	}

	if game.a.x != 0 {
		games = append(games, game)
	}

	return games
}

func calc(g Game, ignoreLimit bool) int {
	b := (g.prize.y*g.a.x - g.prize.x*g.a.y) / (g.b.y*g.a.x - g.b.x*g.a.y)
	a := (g.prize.x - b*g.b.x) / g.a.x

	if !ignoreLimit && (b > 100 || a > 100) {
		return 0
	}

	if a*g.a.x+b*g.b.x == g.prize.x && a*g.a.y+b*g.b.y == g.prize.y {
		return a*3 + b
	}
	return 0
}

func partOne(input string) {
	games := readInput(input, 0)
	sum := 0
	for _, g := range games {
		sum += calc(g, false)
	}
	fmt.Printf("result: %v\n", sum)
}

func partTwo(input string) {
	games := readInput(input, 10000000000000)
	sum := 0
	for _, g := range games {
		sum += calc(g, true)
	}
	fmt.Printf("result: %v\n", sum)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
