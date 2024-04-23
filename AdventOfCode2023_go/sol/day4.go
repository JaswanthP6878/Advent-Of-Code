package sol

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day4(data []string) int {
	sum := 0
	for _, game := range data {
		sum += getScoreForCard(game)
	}
	return sum
}

// create a struct [card_no, mathing_cards]

// check how much would down card would produce and add them
// card_A_produces := card_A + how_much_would_all_below_A_produces
func Day4_2(data []string) int {
	matches := [][]int{}
	for _, line := range data {
		match := getMatchesPerCard(line)
		matches = append(matches, match)
	}
	sols := make([]int, len(matches))
	spawned_cards := func(match int, index int) int {
		sol := 0
		if match == 0 {
			return sol
		}
		for match != 0 && index < len(matches) {
			sol += sols[index]
			index++
			match--
		}
		return sol
	}
	for j := len(matches) - 1; j >= 0; j-- {
		sols[j] = 1 + spawned_cards(matches[j][1], j+1)
	}
	sum := 0
	for _, val := range sols {
		sum += val
	}
	return sum
}

func getMatchesPerCard(game string) []int {
	card_no, winning_nums, nums_we_have := parseCard(game)
	matches := 0

	for _, our := range nums_we_have {
		for _, num := range winning_nums {
			if our == num {
				matches += 1
			}
		}
	}
	return []int{card_no, matches}

}

func getScoreForCard(game string) int {
	// struct == CARD no: <winning-numbers> | <numbers_we_have>
	_, winning_nums, nums_we_have := parseCard(game)
	winnings := 0
	for _, our := range nums_we_have {
		for _, num := range winning_nums {
			if our == num {
				if winnings == 0 {
					winnings = 1
				} else {
					winnings = winnings * 2
				}
			}
		}
	}
	return winnings

}

func parseCard(game string) (int, []int, []int) {
	card_split := strings.Split(game, ": ")
	if card_split == nil || len(card_split) != 2 {
		fmt.Printf("cannot split card data %s", game)
	}
	re := regexp.MustCompile("[0-9]+")
	// get card number
	card_number := re.FindAllString(card_split[0], -1)
	if len(card_number) != 1 {
		fmt.Printf("Error getting card number, for %T", card_number)
	}
	card_number_int, _ := strconv.Atoi(card_number[0])

	cards := strings.Split(card_split[1], " | ")
	if cards == nil || len(cards) != 2 {
		fmt.Printf("Cards not parsed properly: %T", cards)
	}

	winning_cards := re.FindAllString(cards[0], -1)
	our_cards := re.FindAllString(cards[1], -1)

	winning_cards_int := []int{}
	our_cards_int := []int{}

	for _, win := range winning_cards {
		val, err := strconv.Atoi(win)
		// fmt.Println(val)
		if err != nil {
			fmt.Printf("cannot convert val %s to int", win)
			panic(1)
		}
		winning_cards_int = append(winning_cards_int, val)
	}

	for _, win := range our_cards {
		val, err := strconv.Atoi(win)
		// fmt.Println(val)
		if err != nil {
			fmt.Printf("cannot convert val %s to int", win)
			panic(1)
		}
		our_cards_int = append(our_cards_int, val)
	}

	return card_number_int, winning_cards_int, our_cards_int

}
