package sol

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// cardType enum
type cardType int

const (
	_ cardType = iota
	Highcard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Card struct {
	value    string
	cardType cardType
	bid      int // bid amount
}

func Day7(data []string) int {

	// all arrays for each kind
	FiveofKindCards := []Card{}
	FourOfAKindCards := []Card{}
	FullHouses := []Card{}
	ThreeOfAKinds := []Card{}
	TwoPairs := []Card{}
	OnePairs := []Card{}
	HighCards := []Card{}

	for _, val := range data {
		card := readCard(val)
		switch card.cardType {
		case 0:
			fmt.Println("zero card")
			fmt.Println(card)
		case Highcard:
			HighCards = append(HighCards, card)
		case OnePair:
			OnePairs = append(OnePairs, card)
		case TwoPair:
			TwoPairs = append(TwoPairs, card)
		case ThreeOfAKind:
			ThreeOfAKinds = append(ThreeOfAKinds, card)
		case FullHouse:
			FullHouses = append(FullHouses, card)
		case FourOfAKind:
			FourOfAKindCards = append(FourOfAKindCards, card)
		case FiveOfAKind:
			FiveofKindCards = append(FiveofKindCards, card)
		}
	}
	// sorting all of them
	sort.Slice(HighCards, func(i, j int) bool {
		return CardLess(HighCards[i], HighCards[j])
	})
	sort.Slice(OnePairs, func(i, j int) bool {
		return CardLess(OnePairs[i], OnePairs[j])
	})
	sort.Slice(TwoPairs, func(i, j int) bool {
		return CardLess(TwoPairs[i], TwoPairs[j])
	})
	sort.Slice(ThreeOfAKinds, func(i, j int) bool {
		return CardLess(ThreeOfAKinds[i], ThreeOfAKinds[j])
	})
	sort.Slice(FullHouses, func(i, j int) bool {
		return CardLess(FullHouses[i], FullHouses[j])
	})
	sort.Slice(FourOfAKindCards, func(i, j int) bool {
		return CardLess(FourOfAKindCards[i], FourOfAKindCards[j])
	})
	sort.Slice(FiveofKindCards, func(i, j int) bool {
		return CardLess(FiveofKindCards[i], FiveofKindCards[j])
	})

	// fmt.Println(HighCards)
	// fmt.Println(OnePairs)
	// fmt.Println(TwoPairs)
	// fmt.Println(ThreeOfAKinds)
	// fmt.Println(FullHouses)
	// fmt.Println(FourOfAKindCards)
	// fmt.Println(FiveofKindCards)

	rank := 1
	winnings := 0
	for _, val := range HighCards {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range OnePairs {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range TwoPairs {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range ThreeOfAKinds {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range FullHouses {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range FourOfAKindCards {
		winnings += val.bid * rank
		rank++
	}
	for _, val := range FiveofKindCards {
		winnings += val.bid * rank
		rank++
	}

	return winnings
}

// util sorting a list
// returns true if c1 < c2; returns false if c1>c2
func CardLess(c1, c2 Card) bool {
	for i := 0; i < len(c1.value); i++ {
		if GetCardOrder(c1.value[i]) < GetCardOrder(c2.value[i]) {
			return true
		} else if GetCardOrder(c1.value[i]) > GetCardOrder(c2.value[i]) {
			return false
		} else {
			continue
		}
	}
	return true
}

func GetCardOrder(b byte) int {
	switch b {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 11
	default:
		val, _ := strconv.Atoi(string(b))
		return val
	}
}

func readCard(data string) Card {
	// fmt.Println(data)
	vals := strings.Split(data, " ")
	new_card := Card{}
	new_card.value = vals[0]
	val, err := strconv.Atoi(vals[1])
	if err != nil {
		fmt.Printf("%s bid cannot be converted to int\n", vals[1])
		panic(1)
	}
	new_card.bid = val
	new_card.cardType = getCardTypeWithJoker(new_card.value)

	return new_card

}

// made using freqArray check for ways to be modified

// have to make changes to getting joker card J
func getCardType(str string) cardType {
	// fmt.Println(str)
	vec_int := make([]int, 13)
	// parsing the cards and storing in a
	// freq array
	for i := 0; i < len(str); i++ {
		if str[i] == 'T' {
			vec_int[8]++
		} else if str[i] == 'J' {
			vec_int[9]++
		} else if str[i] == 'Q' {
			vec_int[10]++
		} else if str[i] == 'K' {
			vec_int[11]++
		} else if str[i] == 'A' {
			vec_int[12]++
		} else {
			conv, _ := strconv.Atoi(string(str[i]))
			vec_int[conv-2]++
		}
	}
	max_int := 0
	for _, val := range vec_int {
		if val > max_int {
			max_int = val
		}
	}
	// fmt.Println(vec_int)
	if max_int == 5 {
		return FiveOfAKind
	} else if max_int == 4 {
		return FourOfAKind
	} else if max_int == 3 {
		for _, val := range vec_int {
			if val == 2 {
				return FullHouse
			}
		}
		return ThreeOfAKind
	} else if max_int == 2 {
		count := 0
		for _, val := range vec_int {
			if val == 2 {
				count++
			}
		}
		if count == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	} else {
		return Highcard
	}

}

// for 7_2
// gettting too-low answer
func getCardTypeWithJoker(str string) cardType {
	// fmt.Println(str)
	vec_int := make([]int, 13)
	jokers := 0
	// parsing the cards and storing in a
	// freq array
	for i := 0; i < len(str); i++ {
		if str[i] == 'T' {
			vec_int[8]++
		} else if str[i] == 'J' {
			jokers += 1
		} else if str[i] == 'Q' {
			vec_int[10]++
		} else if str[i] == 'K' {
			vec_int[11]++
		} else if str[i] == 'A' {
			vec_int[12]++
		} else {
			conv, _ := strconv.Atoi(string(str[i]))
			vec_int[conv-2]++
		}
	}

	// fmt.Printf("jokers are := %d\n", jokers)
	max_int := 0
	for _, val := range vec_int {
		if val > max_int {
			max_int = val
		}
	}
	// fmt.Printf("max int is %d\n", max_int)

	if jokers == 5 {
		return FiveOfAKind
	} else if jokers == 4 {
		return FiveOfAKind
	} else if jokers == 3 {
		if max_int == 2 {
			return FiveOfAKind
		}
		return FourOfAKind
	} else if jokers == 2 {
		if max_int == 3 {
			return FiveOfAKind
		} else if max_int == 2 {
			return FourOfAKind
		} else {
			return ThreeOfAKind
		}
	} else if jokers == 1 {
		if max_int == 4 {
			return FiveOfAKind
		} else if max_int == 3 {
			return FourOfAKind
		} else if max_int == 2 {
			count := 0
			for _, val := range vec_int {
				if val == 2 {
					count++
				}
			}
			if count == 2 {
				return FullHouse
			} else {
				return ThreeOfAKind
			}
		} else {
			return OnePair
		}
	} else { // jokers == 0 case;
		return getCardType(str)
	}
}
