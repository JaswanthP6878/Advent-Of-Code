package sol

import (
	"fmt"
	"sort"
	"testing"
)

func TestGetCardType(t *testing.T) {
	input := []struct {
		card  string
		cardt cardType
	}{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPair},
		{"23456", Highcard},
	}
	for _, tt := range input {
		card := getCardType(tt.card)
		if card != tt.cardt {
			t.Errorf("Got wrong Card Type")
		}
	}
}

func TestSortCard(t *testing.T) {
	input := []Card{
		Card{value: "2AAAA", cardType: FiveOfAKind},
		Card{value: "3333A", cardType: FiveOfAKind},
		// Card{value: "22222", cardType: FiveOfAKind},
		// Card{value: "QQQQQ", cardType: FiveOfAKind},
		// Card{value: "44444", cardType: FiveOfAKind},
	}

	sort.Slice(input, func(i, j int) bool {
		return CardLess(input[i], input[j])
	})

	fmt.Println(input)

}

func TestCardTypeWithJoker(t *testing.T) {
	input := "6J528"
	val := getCardTypeWithJoker(input)
	fmt.Println(val)
}
