package sol

import (
	"testing"
)

func TestGetNextNumberInSequence(t *testing.T) {
	input := "1 3 6 10 15 21"
	sol := getNextNumberInSequence(input)
	if sol != 28 {
		t.Errorf("got %d not 28", sol)
	}

}

func TestPrevNumber(t *testing.T) {
	input := "10 13 16 21 30 45"
	sol := getPrevNumberInSequence(input)
	if sol != 5 {
		t.Errorf("got %d, have to get 5", sol)
	}
}
