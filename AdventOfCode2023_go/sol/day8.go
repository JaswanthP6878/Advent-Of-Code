package sol

import (
	"fmt"
	"strings"
)

type adjacent struct {
	left  string
	right string
}

func (a *adjacent) getLR(b byte) string {
	if b == 'L' {
		return a.left
	}
	return a.right
}

func Day8(data []string) int {
	Instructions := data[0]
	location_maps := createMapGraph(data[2:])
	// starting location is AAA; given in question
	curr_Location := "AAA"
	len_instructions := len(Instructions)
	i := 0 // instruction pointer
	steps := 0
	for curr_Location != "ZZZ" {
		if i == len_instructions {
			i = 0
		}
		curr_Location = location_maps[curr_Location].getLR(Instructions[i])
		i++
		steps++
	}
	return steps
}

func Day8_2opti(data []string) int {
	instructions := data[0]
	// nodes that end with A
	location_maps, start_nodes := createMapGraphWithAendings(data[2:])

	len_instructions := len(instructions)
	for _, start_node := range start_nodes {
		tmp := start_node
		// instruction pointer
		i := 0
		steps := 0
		for tmp[2] != 'Z' {
			if i == len_instructions {
				i = 0
			}
			tmp = location_maps[tmp].getLR(instructions[i])
			i++
			steps++
		}

		fmt.Println(steps)
	}
	// of all the steps take lcm; that is the output
	return 0
}

// map key is current location, value is the left and right location
func createMapGraph(data []string) map[string]*adjacent {
	sol := map[string]*adjacent{}
	for _, val := range data {
		location, left, right := parseLocation(val)
		if _, ok := sol[location]; !ok {
			sol[location] = &adjacent{left: left, right: right}
		} else {
			fmt.Printf("location already in map %s\n", location)
			panic(1)
		}
	}
	return sol

}
func createMapGraphWithAendings(data []string) (map[string]*adjacent, []string) {
	start_nodes := []string{}
	sol := map[string]*adjacent{}
	for _, val := range data {
		location, left, right := parseLocation(val)
		if location[2] == 'A' {
			start_nodes = append(start_nodes, location)
		}
		if _, ok := sol[location]; !ok {
			sol[location] = &adjacent{left: left, right: right}
		} else {
			fmt.Printf("location already in map %s\n", location)
			panic(1)
		}
	}
	return sol, start_nodes

}

func parseLocation(str string) (string, string, string) {
	str_split := strings.Split(str, " = ")

	location := str_split[0]

	lrsplit := strings.Split(str_split[1][1:9], ", ")

	left := lrsplit[0]
	right := lrsplit[1]
	return location, left, right

}
