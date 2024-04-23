package sol

import (
	"fmt"
)

type starLocation struct {
	i    int
	j    int
	dist int
}

// func (s starLocation) print() { fmt.Printf("%d_%d\n", s.i, s.j) }
func Day11(data []string) int {
	empty_rows, empty_cols := expandUniverse(data)
	fmt.Println(empty_rows)
	fmt.Println(empty_cols)
	star_locations := getStarLocations(data, empty_rows, empty_cols)
	// for _, val := range star_locations {
	// 	val.print()
	// }
	num_rows := len(data) + len(empty_rows)
	// num_rows := len(data) + len(empty_rows)*999999
	// num_cols := len(data) + len(empty_cols)*999999
	num_cols := len(data) + len(empty_cols)
	sum_shortest_paths := 0
	for i := 0; i < len(star_locations)-1; i++ {
		for j := i + 1; j < len(star_locations); j++ {
			sum_shortest_paths += getShortestPath(star_locations[i], star_locations[j], num_rows, num_cols)
		}
	}
	fmt.Println(sum_shortest_paths)
	return sum_shortest_paths
}

func getShortestPath(source starLocation, dest starLocation, num_rows int, num_cols int) int {
	// creting a visited array with all false
	visited := [][]bool{}
	for i := 0; i < num_rows; i++ {
		visi := []bool{}
		for j := 0; j < num_cols; j++ {
			visi = append(visi, false)
		}
		visited = append(visited, visi)
	}

	isValid := func(i int, j int) bool {
		if i >= 0 && j >= 0 && i < num_rows && j < num_cols && !visited[i][j] {
			return true
		}
		return false
	}

	// running BFS code for finding shortest distances
	queue := make([]starLocation, 0)
	visited[source.i][source.j] = true
	queue = append(queue, source)
	for len(queue) != 0 {
		curr_location := queue[0]
		queue = queue[1:]

		if curr_location.i == dest.i && curr_location.j == dest.j {
			return curr_location.dist
		}

		// top
		if isValid(curr_location.i-1, curr_location.j) {
			queue = append(queue, starLocation{curr_location.i - 1, curr_location.j, curr_location.dist + 1})
			visited[curr_location.i-1][curr_location.j] = true
		}
		// bot
		if isValid(curr_location.i+1, curr_location.j) {
			queue = append(queue, starLocation{curr_location.i + 1, curr_location.j, curr_location.dist + 1})
			visited[curr_location.i+1][curr_location.j] = true
		}

		// left
		if isValid(curr_location.i, curr_location.j-1) {
			queue = append(queue, starLocation{curr_location.i, curr_location.j - 1, curr_location.dist + 1})
			visited[curr_location.i][curr_location.j-1] = true
		}

		// right
		if isValid(curr_location.i, curr_location.j+1) {
			queue = append(queue, starLocation{curr_location.i, curr_location.j + 1, curr_location.dist + 1})
			visited[curr_location.i][curr_location.j+1] = true
		}
	}

	return -1
}

func getStarLocations(data []string, empty_rows []int, empty_cols []int) []starLocation {
	starPositions := []starLocation{}
	getStarPosition := func(i int, j int) starLocation {
		count := 0
		for _, val := range empty_cols {
			if j < val {
				break
			}
			count++
		}
		j = j + count
		// j = j + count*999999
		count = 0
		for _, val := range empty_rows {
			if i < val {
				break
			}
			count++
		}
		i = i + count
		// i = i + count*999999
		return starLocation{i: i, j: j}
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] == '#' {
				star_position := getStarPosition(i, j)
				starPositions = append(starPositions, star_position)
			}
		}
	}
	return starPositions
}

// returns the rows and cols that need to be expanded
func expandUniverse(data []string) ([]int, []int) {
	// fmt.Println(data)
	empty_rows := []int{}
	for i := range data {
		count_stars := 0
		for j := range data[i] {
			if data[i][j] == '#' {
				count_stars++
			}
		}
		if count_stars == 0 {
			empty_rows = append(empty_rows, i)
		}
	}
	empty_cols := []int{}
	for j := range data[0] {
		count_stars := 0
		for i := range data {
			if data[i][j] == '#' {
				count_stars++
			}
		}
		if count_stars == 0 {
			empty_cols = append(empty_cols, j)
		}
	}
	return empty_rows, empty_cols
}
