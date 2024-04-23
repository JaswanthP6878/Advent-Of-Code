package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func q1() {

	x := []int{}
	file, err := os.Open("./1.txt")
	if err !=  nil {
		panic("File cannot be opend")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("string cannot be converted")
		}
		x = append(x,i)

	}
	prev := x[0]
	count := 0
	for i := 1; i < len(x); i++ {
		if x[i] - prev > 0 {
			count += 1
		}
		prev = x[i]
	}
	fmt.Println(count)

}
