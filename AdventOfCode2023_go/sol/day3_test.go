package sol

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegEx(t *testing.T) {

	re := regexp.MustCompile("[0-9]+")
	input := "12..324..1.."
	vals := re.FindAllStringIndex(input, -1)
	fmt.Print(vals)

}

func TestGetNumbers(t *testing.T) {
	input := []string{"123.11.*.2", "123*11...2"}
	vals := getSpecialLocations(input)
	fmt.Println(vals)
}

func TestDay3(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	nums := getNumbers(input)
	fmt.Println(nums)
	locs := getSpecialLocations(input)
	fmt.Println(locs)
	sum := Day3(input)
	if sum != 4361 {
		t.Fatalf("expected 4361 got . %d", sum)
	}
	sum_2 := Day3_2(input)
	if sum_2 != 467835 {
		t.Errorf("Got %d", sum_2)

	}
}
