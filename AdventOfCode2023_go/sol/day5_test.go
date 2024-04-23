package sol

import (
	"fmt"
	"testing"
)

// func TestGetMapsFromData(t *testing.T) {
// 	input := "soil-to-fertilizer map: 0 15 37 37 52 2 39 0 15"
// 	_ = getMapsFromData(input)
// }

func TestMapfunc(t *testing.T) {
	map_vals := [][]int{
		[]int{50, 98, 2},
		[]int{52, 50, 48},
	}
	seeds := []int{79, 14, 55, 13}
	map_func := getMapFuncFromMap(map_vals)

	for _, val := range seeds {
		fmt.Println(map_func(val))
	}

}

// func Test82(t *testing.T) {
// 	in := 8

// }
