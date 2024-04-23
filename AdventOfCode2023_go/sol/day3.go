package sol

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3(data []string) int {
	sum := 0
	// sols := []int{}
	specialLocs := getSpecialLocations(data)
	nums := getNumbers(data)
	for _, loc := range specialLocs {
		j_l, i_l := loc[0], loc[1]
		for _, num := range nums {
			// if num already included in sum
			if num[0] == 1 {
				continue
			}
			if num[2] == j_l+1 || num[2] == j_l-1 {
				if (num[3] <= i_l && num[4] >= i_l) || (num[3] == i_l+1 || num[4] == i_l-1) {
					// sols = append(sols, num[1])
					sum += num[1]
					num[0] = 1
				}
			}

			if num[2] == j_l {
				if num[3] == i_l+1 || num[4] == i_l-1 {
					// sols = append(sols, num[1])
					sum += num[1]
					num[0] = 1
				}
			}
		}
	}
	// fmt.Println("printing solutions")
	// fmt.Println(sols)
	return sum

}

func Day3_2(data []string) int {
	sum := 0
	// traverse and find * symbols
	nums := getNumbers(data)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[1]); j++ {
			if data[i][j] == '*' {
				parts := getPartsforGears(i, j, nums)
				if len(parts) != 2 {
					continue
				}
				sum += parts[0] * parts[1]
			}

		}
	}

	return sum

}

func getPartsforGears(j_l int, i_l int, nums [][]int) []int {
	sols := []int{}
	for _, num := range nums {
		if num[2] == j_l+1 || num[2] == j_l-1 {
			if (num[3] <= i_l && num[4] >= i_l) || (num[3] == i_l+1 || num[4] == i_l-1) {
				sols = append(sols, num[1])
			}
		}
		if num[2] == j_l {
			if num[3] == i_l+1 || num[4] == i_l-1 {
				sols = append(sols, num[1])
			}
		}
	}
	return sols
}

func getSpecialLocations(data []string) [][]int {
	locs := [][]int{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[1]); j++ {
			if isSpecialChar(data[i][j]) {
				locs = append(locs, []int{i, j})
			}
		}
	}
	return locs
}

// value skeleton -> [included, number, line_number, left_pointer, right_pointer]
func getNumbers(data []string) [][]int {

	sol := [][]int{}

	re := regexp.MustCompile("[0-9]+")
	for i, line := range data {
		indexes := re.FindAllStringIndex(line, -1)
		if indexes == nil {
			continue
		}
		nums := getNumbersFromIndexes(indexes, line)

		if len(indexes) != len(nums) {
			fmt.Println("indexes and nums length does not match")
			fmt.Println(indexes)
			fmt.Println(nums)
		}
		for k := range nums {
			if indexes[k][0]+1 == indexes[k][1] {
				val := nums[k]
				sol = append(sol, []int{0, val, i, indexes[k][0], indexes[k][0]})
			} else {
				val := nums[k]
				sol = append(sol, []int{0, val, i, indexes[k][0], indexes[k][1] - 1})
			}
		}
	}

	return sol
}

// util
func isSpecialChar(char byte) bool {
	if (char >= '0' && char <= '9') || (char == '.') {
		return false
	}
	return true
}

func getNumbersFromIndexes(ind [][]int, line string) []int {
	sols := []int{}
	for _, arr := range ind {
		l, r := arr[0], arr[1]-1
		if l == r {
			val, err := strconv.Atoi(string(line[l]))
			if err != nil {
				fmt.Printf("cannot convert %s to int", string(line[l]))
			}
			sols = append(sols, val)
		} else {
			var sb strings.Builder
			for l <= r {
				sb.WriteByte(line[l])
				l++
			}
			val, err := strconv.Atoi(sb.String())
			if err != nil {
				fmt.Printf("cannto convert string %s, to int", sb.String())
			}
			sols = append(sols, val)
		}
	}
	return sols

}
