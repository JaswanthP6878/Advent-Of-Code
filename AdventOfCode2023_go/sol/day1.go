package sol

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// day1_1 solution:
func Day1(input []string) int {
	var sum int = 0
	for _, data := range input {
		sum += getNumber(data)
	}
	return sum

}

// BAT-SHIT Insane(BSI) strat.. cause fck i dont wanna write and handle all edgecases writing a lexer
// This also works on the first question i.e day1_1
func Day1_2BSI(input []string) int {
	sum := 0
	var re1 = regexp.MustCompile(`(?m)(one|1)`)
	var re2 = regexp.MustCompile(`(?m)(two|2)`)
	var re3 = regexp.MustCompile(`(?m)(three|3)`)
	var re4 = regexp.MustCompile(`(?m)(four|4)`)
	var re5 = regexp.MustCompile(`(?m)(five|5)`)
	var re6 = regexp.MustCompile(`(?m)(six|6)`)
	var re7 = regexp.MustCompile(`(?m)(seven|7)`)
	var re8 = regexp.MustCompile(`(?m)(eight|8)`)
	var re9 = regexp.MustCompile(`(?m)(nine|9)`)

	rgs := []*regexp.Regexp{re1, re2, re3, re4, re5, re6, re7, re8, re9}
	for _, str := range input {
		sum += getNumberFromMatchBSI(str, rgs)

	}
	return sum
}

func getNumberFromMatchBSI(str string, rgs []*regexp.Regexp) int {
	// key : each digit from 1 to nine
	// value : the location of min and max location of that number
	tmp := map[int][]int{}

	for i, rg := range rgs {
		arr := rg.FindAllStringIndex(str, -1)
		if arr != nil {
			new_arr := []int{}
			for _, ar := range arr {
				new_arr = append(new_arr, ar[0])
			}
			max := -1
			min := 2000
			for _, val := range new_arr {
				if val > max {
					max = val
				}
				if val < min {
					min = val
				}
			}
			tmp[i+1] = []int{min, max}
		}
	}
	min := 2000
	max := -1
	min_digit := 0
	max_digit := 0
	for k, v := range tmp {
		if v[0] < min {
			min_digit = k
			min = v[0]
		}

		if v[1] > max {
			max_digit = k
			max = v[1]
		}
	}
	return min_digit*10 + max_digit
}

func getNumber(line string) int {
	var i int = 0
	var j int = len(line) - 1
	for !isDigit(line[i]) {
		i++
	}
	for !isDigit(line[j]) {
		j--
	}
	var sb strings.Builder
	sb.WriteByte(line[i])
	sb.WriteByte(line[j])
	st := sb.String()

	i, err := strconv.Atoi(st)
	if err != nil {
		fmt.Printf("Cannot convert to int . got string %s", st)
		panic(1)
	}

	return i
}

func isDigit(k byte) bool {
	return k >= '0' && k <= '9'
}
