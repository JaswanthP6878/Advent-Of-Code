package sol

import (
	"fmt"
	"testing"
)

func Test_Daty5(t *testing.T) {
	input := "BFFFBBFRRR"
	val := GetSeatNumber(input)
	fmt.Println(val)
}
