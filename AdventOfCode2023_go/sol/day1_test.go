package sol

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDay1_1(t *testing.T) {
	input := []string{"1abc2", "a1b2c3d4e5f", "treb7uchet", "pqr3stu8vwx"}

	if Day1(input) != 142 {
		t.Errorf("Expected 142 . got %d", Day1(input))
	}

}

func TestRegExp(t *testing.T) {
	var re = regexp.MustCompile(`(?m)(seven)`)

	var str = `pvqqdfcfourxqvsfcs1two9sevenninetwoseven`
	if !re.MatchString(str) {
		t.Errorf("cannnot match string")
	}

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}

}

func TestDay1_2(t *testing.T) {
	input := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	// input := []string{"eightthree"}
	if Day1_2BSI(input) != 281 {
		t.Errorf("Expected 281 . got %d", Day1_2BSI(input))
	}
}
