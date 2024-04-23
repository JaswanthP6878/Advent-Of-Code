package sol

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// 50 98 2
// 52 50 48
// The first line has a destination range start of 50, a source range start of 98, and a range length of 2.
// This line means that the source range starts at 98 and contains two values: 98 and 99.
// The destination range is the same length, but it starts at 50, so its two values are 50 and 51.
// With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

func Day5(data []string) int {
	seeds := getSeeds(data[0])
	// maps are like:
	// [ [ [50 98 2] [52 50 48] ], <next-map>]
	maps := getMapsFromData(data)
	map_funcs := []func(int) int{}
	for i := 0; i < len(maps); i++ {
		map_funcs = append(map_funcs, getMapFuncFromMap(maps[i]))
	}

	for i := 0; i < len(seeds); i++ {
		for _, map_func := range map_funcs {
			seeds[i] = map_func(seeds[i])
		}
	}
	min := 2147483647
	for _, val := range seeds {
		// fmt.Println(val)
		if val < min {
			min = val
		}
	}
	// for _, val := range maps {
	// 	fmt.Println(val)
	// }
	return min
}

func Day5_2(data []string) int {
	seed_ranges := getSeeds(data[0])
	// get ranges from i,i+1
	ranges := [][]int{}
	for i := 0; i <= len(seed_ranges)-2; i = i + 2 {
		start_range, end_range := seed_ranges[i], seed_ranges[i]+seed_ranges[i+1]-1
		ranges = append(ranges, []int{start_range, end_range})
	}
	fmt.Println(ranges)
	maps := getMapsFromData(data)
	map_funcs := []func(int) int{}
	for i := 0; i < len(maps); i++ {
		map_funcs = append(map_funcs, getMapFuncFromMap(maps[i]))
	}

	sols := make([]int, len(ranges))
	var wg sync.WaitGroup
	// we have map_func constructed
	for i, r := range ranges {
		wg.Add(1)
		go func(r []int, i int) {
			fmt.Println(r)
			defer wg.Done()
			min := 2147483647
			for k := r[0]; k <= r[1]; k++ {
				tmp := k
				for _, map_func := range map_funcs {
					tmp = map_func(tmp)
				}
				if tmp < min {
					min = tmp
				}
			}
			sols[i] = min
		}(r, i)
	}

	wg.Wait()
	fmt.Println(sols)

	min := sols[0]
	for _, val := range sols {
		if val <= min {
			min = val
		}
	}
	return min

}

func getMapFuncFromMap(m [][]int) func(int) int {
	map_func := func(x int) int {
		// map struct == [dest, source, range]
		for _, map_val := range m {
			if x >= map_val[1] && x < map_val[1]+map_val[2] {
				return map_val[0] + (x - map_val[1])
			}
		}
		return x
	}
	return map_func

}

func getMapsFromData(data []string) [][][]int {
	maps := [][][]int{}
	for i := 2; i < len(data); i++ {
		if strings.Contains(data[i], "map:") {
			temp_map := [][]int{}
			i++
			for i < len(data) && len(data[i]) != 0 {
				vals := strings.Split(data[i], " ")
				val_ints := []int{}
				for _, val := range vals {
					k, _ := strconv.Atoi(val)
					val_ints = append(val_ints, k)
				}
				temp_map = append(temp_map, val_ints)
				i++
			}
			maps = append(maps, temp_map)
		}
	}
	return maps
}

func getSeeds(line string) []int {
	seeds := []int{}
	data := strings.Split(line, ": ")
	vals := strings.Split(data[1], " ")
	for _, val := range vals {
		k, _ := strconv.Atoi(val)
		seeds = append(seeds, k)
	}
	return seeds

}
