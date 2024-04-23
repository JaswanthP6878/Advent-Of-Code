package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func q2() {
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
	l := 0
	r := 3
	count := 0
	for r < len(x) {
		if x[r] - x[l] > 0 {
			count++;
		}
		l += 1
		r = l + 3
	}
	fmt.Println(count)
	
}
