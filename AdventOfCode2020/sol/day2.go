package sol

import (
	"strconv"
	"strings"
)

func Day2(data []string) int {
	count := 0
	for i := range data {
		if isValidPassword_part_2(data[i]) {
			count++
		}
	}
	return count
}

func isValidPassword_part_1(val string) bool {
	// strings.Split(s, sep)
	values := strings.Split(val, " ")
	values[1] = string(values[1][0])
	bounds := strings.Split(values[0], "-")
	lo, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])
	count := 0
	for i := range values[2] {
		if string(values[2][i]) == values[1] {
			count += 1
		}
	}
	if (lo < count && high > count) || lo == count || high == count {
		return true
	}
	return false
}

func isValidPassword_part_2(val string) bool {
	// strings.Split(s, sep)
	values := strings.Split(val, " ")
	values[1] = string(values[1][0])
	bounds := strings.Split(values[0], "-")
	lo, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])
	count := 0
	if string(values[2][lo-1]) == values[1] {
		count += 1
	}
	if string(values[2][high-1]) == values[1] {
		count += 1
	}
	if count == 1 {
		return true
	}
	return false
}
