package sol

import (
	"advent2023/utils"
	"testing"
)

func TestExpandedUniverse(t *testing.T) {
	data, _ := utils.ReadFileFromSol("day11_test.txt")
	_, _ = expandUniverse(data)
}
func TestChangedLocations(t *testing.T) {
	data, _ := utils.ReadFileFromSol("day11_test.txt")
	_ = Day11(data)

}
