package sol

import (
	"fmt"
	"testing"
)

func TestParseLocation(t *testing.T) {
	input := "AAA = (BBB, CCC)"
	loc, l, r := parseLocation(input)
	fmt.Println(loc)
	fmt.Println(l)
	fmt.Println(r)

}
