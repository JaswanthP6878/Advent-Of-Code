package main

import (
	"fmt"

	"advent2020.jaswanthp/sol"
	"advent2020.jaswanthp/utils"
)

func main() {
	out := utils.ReadInputString("day4.txt")
	solution := sol.Day3_2(out)
	fmt.Println(solution)
}
