package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) ([]string, error) {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/resources/" + path)
	if err != nil {
		fmt.Printf("file cannot be opened: %s", err.Error())
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	file.Close()

	return fileLines, nil

}

func ReadFileForDay5(path string) ([]string, error) {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/resources/" + path)
	if err != nil {
		fmt.Printf("file cannot be opened: %s", err.Error())
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			fileLines = append(fileLines, scanner.Text())
		}
	}
	file.Close()

	return fileLines, nil

}

func ReadFileFromSol(path string) ([]string, error) {
	file, err := os.Open("/Users/jaswanthpinnepu/Desktop/AdventOfCode2023_go/resources/" + path)
	if err != nil {
		fmt.Printf("file cannot be opened: %s", err.Error())
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	file.Close()

	return fileLines, nil

}
