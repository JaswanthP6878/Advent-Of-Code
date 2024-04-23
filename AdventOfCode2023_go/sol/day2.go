package sol

import (
	"fmt"
	"strconv"
	"strings"
)

type gameset struct {
	red   int
	blue  int
	green int
}

func (g *gameset) printBalls() {
	fmt.Printf("red: %d, blue: %d, green: %d\n", g.red, g.blue, g.green)
}

func Day2_1(data []string) int {
	sum := 0

	for _, str := range data {
		if isPossible(str) {
			sum += gameId(str)
		}
	}
	return sum
}

func Day2_2(data []string) int {
	prod := 0
	for _, str := range data {
		games := parseGameSets(str)
		var max_red, max_green, max_blue int = 0, 0, 0
		for _, game := range games {
			if game.blue > max_blue {
				max_blue = game.blue
			}
			if game.red > max_red {
				max_red = game.red
			}
			if game.green > max_green {
				max_green = game.green
			}
		}
		if max_red == 0 {
			max_red = 1
		}
		if max_green == 0 {
			max_green = 1
		}
		if max_blue == 0 {
			max_blue = 1
		}
		prod += max_blue * max_green * max_red
	}
	return prod
}

func isPossible(str string) bool {
	// parse game
	gameSets := parseGameSets(str)
	for _, gs := range gameSets {
		if gs.blue > 14 {
			return false
		}
		if gs.red > 12 {
			return false
		}
		if gs.green > 13 {
			return false
		}
	}

	return true

}

func parseGameSets(str string) []gameset {
	out := []gameset{}
	gamesetfull := strings.Split(str, ": ")
	gamesets := gamesetfull[1]
	gamesetSplit := strings.Split(gamesets, "; ")
	for _, gamesetstring := range gamesetSplit {
		g := gameset{}
		balls := strings.Split(gamesetstring, ", ")
		for _, ball := range balls {
			countPerColor := strings.Split(ball, " ")
			if countPerColor[1] == "blue" {
				val, err := strconv.Atoi(countPerColor[0])
				if err != nil {
					fmt.Println(err)
					panic(1)
				}
				g.blue = val
			}
			if countPerColor[1] == "red" {
				val, err := strconv.Atoi(countPerColor[0])
				if err != nil {
					fmt.Println(err)
					panic(1)
				}
				g.red = val
			}
			if countPerColor[1] == "green" {
				val, err := strconv.Atoi(countPerColor[0])
				if err != nil {
					fmt.Println(err)
					panic(1)
				}
				g.green = val
			}
		}
		out = append(out, g)
	}
	return out
}

func gameId(str string) int {
	ss := strings.Split(str, ": ")
	gameNumber := strings.Split(ss[0], " ")[1]
	val, _ := strconv.Atoi(gameNumber)
	return val
}
