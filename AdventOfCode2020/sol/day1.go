package sol

import (
	"fmt"
	"sort"
)

func Day1(data []int) int {
	size := len(data)
	for i := 0; i < size-2; i++ {
		for j := 1; j < size-1; j++ {
			for k := 2; k < size; k++ {
				if data[i]+data[j]+data[k] == 2020 {
					return data[i] * data[j] * data[k]
				}
			}
		}
	}
	return -1
}

// two pointer approach
func Day1_better(data []int) int {
	sort.Ints(data)
	fmt.Println(data)
	i, j := 0, len(data)-1
	for i <= j {
		if data[i]+data[j] > 2020 {
			j--
		} else if data[i]+data[j] < 2020 {
			i++
		} else {
			return data[i] * data[j]
		}
	}
	return -1
}
