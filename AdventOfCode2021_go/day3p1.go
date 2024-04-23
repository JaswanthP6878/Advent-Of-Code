package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getOxygenvalue(values []string) int{

	return 0


}

func getCo2Value(values []string) int {
	return 0
}


func main() {
	items := []string{}
	file, err := os.Open("./day3.txt")
	if err != nil {
		panic("cannot open file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := fmt.Sprintf("%s", scanner.Text())
		items = append(items, x)
	}
	oxval := getOxygenvalue(items)
	co2val := getCo2Value(items)

//
//	var epsilon, gamma string
//	for i := 0; i<len(items[0]); i++ {
//		var zeros, ones int
//		for _, val := range items {
//			if val[i] == '0' {
//				zeros += 1
//			} else {
//				ones += 1
//			}
//		}
//		if zeros > ones {
//			gamma += "0"
//			epsilon += "1"
//		} else {
//			gamma += "1"
//			epsilon += "0"
//		}
//
//	}
//	// parsing the string to int values
//	e, err := strconv.ParseInt(epsilon, 2, 64)
//	if err != nil {
//		panic("cannot convert epsilon to int")
//	}
//	g, err := strconv.ParseInt(gamma, 2, 64)
//	if err != nil {
//		panic("cannot convert gamma to int")
//	}
//	fmt.Println(e * g)
}
