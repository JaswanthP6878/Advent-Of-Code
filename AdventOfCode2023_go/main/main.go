package main

import (
	"advent2023/sol"
	"advent2023/utils"
	"fmt"
	"time"
)

func main() {
	data, err := utils.ReadFile("day11.txt")
	if err != nil {
		fmt.Printf("Data cannot be read. got error %s", err.Error())
	}
	start := time.Now()
	sol := sol.Day11(data)
	fmt.Printf("Time taken for running code %s\n", time.Since(start).String())
	fmt.Println(sol)
}
