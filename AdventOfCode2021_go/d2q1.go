package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type dir struct {
	direction string
	value int
}

func d2q1() {
	x := []dir{}
	file, err := os.Open("./1.txt")
	if err !=  nil {
		panic("File cannot be opend")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		i, err := strconv.Atoi(strs[1])
		if err != nil {
			panic("string cannot be converted")
		}
		dirtemp := dir{strs[0],i}
		x = append(x,dirtemp)
	}
	aim := 0
	depth := 0
	position := 0
	for i := 0; i < len(x); i++ {
		if x[i].direction == "forward" {
			position += x[i].value
			depth += x[i].value * aim
		} 

		if x[i].direction == "down" {
			aim += x[i].value
		}
		if x[i].direction == "up" {
			aim -= x[i].value
		}
	}

	fmt.Println(depth * position)

		
}
