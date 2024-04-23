package sol

import "fmt"

func Day3_2(data []string) int {
	tree_count := 1
	tree_count *= traverse_x_y(data, 1, 1)
	tree_count *= traverse_x_y(data, 3, 1)
	tree_count *= traverse_x_y(data, 5, 1)
	tree_count *= traverse_x_y(data, 7, 1)
	tree_count *= traverse_x_y(data, 1, 2)

	return tree_count
}

func traverse_x_y(data []string, right int, down int) int {
	count_trees := 0
	starting_x, starting_y := 0, 0
	fmt.Println((data[0]))
	for starting_y < len(data) {
		if data[starting_y][starting_x%len(data[0])] == '#' {
			count_trees++
		}
		starting_y += down
		starting_x += right
	}

	return count_trees
}
