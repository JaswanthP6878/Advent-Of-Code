package sol

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_stringParse(t *testing.T) {
	val := "1-3 a: abcdaaae"
	values := strings.Split(val, " ")
	values[1] = string(values[1][0])
	bounds := strings.Split(values[0], "-")
	lo, _ := strconv.Atoi(bounds[0])
	high, _ := strconv.Atoi(bounds[1])
	count := 0
	for i := range values[2] {
		if string(values[2][i]) == values[1] {
			count += 1
		}
	}
	fmt.Print(lo, high, values, count)
	if len(values) != 3 {
		t.Error("not correct")
	}
}
