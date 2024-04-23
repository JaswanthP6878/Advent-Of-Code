package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func ReadInput(input_file_name string) []int {
	output := []int{}
	working_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get working dir")
		os.Exit(1)
	}
	fileLoc := filepath.Join(working_dir, "input", input_file_name)
	fd, err := os.Open(fileLoc)
	if err != nil {
		fmt.Printf("Cannot open file, loc : %s", fileLoc)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		output = append(output, val)
	}

	return output
}

func ReadInputString(input_file_name string) []string {
	output := []string{}
	working_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get working dir")
		os.Exit(1)
	}
	fileLoc := filepath.Join(working_dir, "input", input_file_name)
	fd, err := os.Open(fileLoc)
	if err != nil {
		fmt.Printf("Cannot open file, loc : %s", fileLoc)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}
func ReadInputBytes(input_file_name string) [][]byte {
	output := [][]byte{}
	working_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get working dir")
		os.Exit(1)
	}
	fileLoc := filepath.Join(working_dir, "input", input_file_name)
	fd, err := os.Open(fileLoc)
	if err != nil {
		fmt.Printf("Cannot open file, loc : %s", fileLoc)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		output = append(output, scanner.Bytes())
	}
	return output
}
