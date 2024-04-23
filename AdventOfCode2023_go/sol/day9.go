package sol

import (
	"fmt"
	"strconv"
	"strings"
)

// Maybe think in terms of AP series finding and then
// backtrack from there.
func Day9(data []string) int {
	sol := 0
	for _, val := range data {
		sol += getNextNumberInSequence(val)
	}
	return sol

}

// reverse the strings and take subractions
func Day9_2(data []string) int {
	sol := 0
	for _, val := range data {
		sol += getPrevNumberInSequence(val)
	}
	return sol

}
func getPrevNumberInSequence(val string) int {
	input_arr := []int{}
	str_split := strings.Split(val, " ")
	for _, val := range str_split {
		str_int, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("str %s cannot be converted to int", val)
			panic(1)
		}
		input_arr = append(input_arr, str_int)
	}
	first_vals := []int{input_arr[0]}
	for !allZeros(input_arr) {
		diff_arr := []int{}
		for i := 0; i < len(input_arr)-1; i++ {
			diff_arr = append(diff_arr, input_arr[i+1]-input_arr[i])
		}
		first_vals = append(first_vals, diff_arr[0])
		input_arr = diff_arr
	}
	// fmt.Println(first_vals)
	sol := 0
	for i := len(first_vals) - 1; i >= 0; i-- {
		// fmt.Println(sol)
		sol = first_vals[i] - sol
	}
	return sol
}

func getNextNumberInSequence(val string) int {
	input_arr := []int{}
	str_split := strings.Split(val, " ")
	for _, val := range str_split {
		str_int, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("str %s cannot be converted to int", val)
			panic(1)
		}
		input_arr = append(input_arr, str_int)
	}
	last_vals := []int{input_arr[len(input_arr)-1]}
	for !allZeros(input_arr) {
		diff_arr := []int{}
		for i := 0; i < len(input_arr)-1; i++ {
			diff_arr = append(diff_arr, input_arr[i+1]-input_arr[i])
		}
		last_vals = append(last_vals, diff_arr[len(diff_arr)-1])
		input_arr = diff_arr
	}
	sol := 0
	for i := len(last_vals) - 1; i >= 0; i-- {
		sol += last_vals[i]
	}
	return sol
}

func allZeros(data []int) bool {
	count := 0
	for _, val := range data {
		if val == 0 {
			count++
		}
	}
	return count == len(data)
}
