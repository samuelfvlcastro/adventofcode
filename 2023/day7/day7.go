package day7

import (
	"strconv"
	"strings"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var ranksWildCard = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type card struct {
	symbol string
	rank   int
}

type hand struct {
	desc     string
	cards    []card
	strength int
	bid      int
}

func part1(handLines []string) int {
	useWildCard := false
	currRanks := ranks
	if useWildCard {
		currRanks = ranksWildCard
	}

	hands := []hand{}
	for _, line := range handLines {
		h := parseHand(line, currRanks)
		h.strength = findHandStrength(h, useWildCard)
		hands = append(hands, h)
	}

	sortedHands := sortHands(hands)

	total := 0
	for i, h := range sortedHands {
		total += (i + 1) * h.bid
	}

	return total
}

func part2(handLines []string) int {
	useWildCard := true
	currRanks := ranks
	if useWildCard {
		currRanks = ranksWildCard
	}

	hands := []hand{}
	for _, line := range handLines {
		h := parseHand(line, currRanks)
		h.strength = findHandStrength(h, useWildCard)
		hands = append(hands, h)
	}

	sortedHands := sortHands(hands)

	total := 0
	for i, h := range sortedHands {
		total += (i + 1) * h.bid
	}

	return total
}

func parseHand(line string, currentRank []string) hand {
	parts := strings.Split(line, " ")

	cards := strings.Split(parts[0], "")
	bid, _ := strconv.Atoi(parts[1])
	h := hand{
		desc:  "",
		cards: []card{},
		bid:   bid,
	}
	for _, c := range cards {
		for i := len(currentRank) - 1; i >= 0; i-- {
			if c == currentRank[i] {
				h.desc += c
				h.cards = append(h.cards, card{symbol: c, rank: i + 1})
			}
		}
	}

	return h
}

func findHandStrength(h hand, useWildCard bool) int {
	sortedCards := sortCards(h.cards)
	handGroups := groupCardsByHandType(sortedCards)
	if useWildCard {
		handGroups = regroupWithJokeWildCardRule(handGroups)
	}
	numbGroups := len(handGroups)

	switch numbGroups {
	// Five of a Kind ([T T T T T])
	case 1:
		return 7
	case 2:
		for _, group := range handGroups {
			// Four of a Kind ( [T T T T] [2] )
			if group == 4 {
				return 6
			}
		}
		// Full House ([T T T] [2 2])
		return 5
	case 3:
		for _, group := range handGroups {
			// Three of a Kind ([T T T] [3 2])
			if group == 3 {
				return 4
			}
		}
		//  Two Pairs ([T T] [K K] [3])
		return 3
	case 4:
		// One Pair ([T T] [4] [3] [2])
		return 2
	default:
		// High Card ([T] [T] [4] [3] [2])
		return 1
	}
}

func groupCardsByHandType(cards []card) map[string]int {
	groups := map[string]int{}

	groupSize := 1
	curCard := cards[0].symbol
	for c1 := 1; c1 < len(cards); c1++ {
		if curCard == cards[c1].symbol {
			groupSize++
		} else {
			groups[curCard] = groupSize
			groupSize = 1
		}
		curCard = cards[c1].symbol
	}
	groups[curCard] = groupSize

	return groups
}

func regroupWithJokeWildCardRule(groups map[string]int) map[string]int {
	jGroup, ok := groups["J"]
	if !ok {
		return groups
	}

	prevGroupSize := 1
	groupId := ""
	for i, g := range groups {
		if prevGroupSize <= g && i != "J" {
			prevGroupSize = g
			groupId = i
		}
	}
	delete(groups, "J")

	groups[groupId] = groups[groupId] + jGroup

	return groups
}

func sortCards(cards []card) []card {
	sortedCards := []card{}
	sortedCards = append(sortedCards, cards...)
	for i := 0; i < len(sortedCards); i++ {
		for j := 0; j < len(sortedCards)-1-i; j++ {
			if sortedCards[j].rank < sortedCards[j+1].rank {
				temp := sortedCards[j]
				sortedCards[j] = sortedCards[j+1]
				sortedCards[j+1] = temp
			}
		}
	}
	return sortedCards
}

func sortHands(hands []hand) []hand {
	sortedHands := []hand{}
	sortedHands = append(sortedHands, hands...)
	for i := 0; i < len(sortedHands); i++ {
		for j := 0; j < len(sortedHands)-1-i; j++ {
			if sortedHands[j].strength > sortedHands[j+1].strength {
				temp := sortedHands[j]
				sortedHands[j] = sortedHands[j+1]
				sortedHands[j+1] = temp
			}
		}
	}

	for i := 0; i < len(sortedHands); i++ {
		for j := 0; j < len(sortedHands)-1-i; j++ {
			if sortedHands[j].strength == sortedHands[j+1].strength && handAHasStrongerFirstCard(sortedHands[j], sortedHands[j+1]) {
				temp := sortedHands[j]
				sortedHands[j] = sortedHands[j+1]
				sortedHands[j+1] = temp
			}
		}
	}

	return sortedHands
}

func handAHasStrongerFirstCard(ha hand, hb hand) bool {
	for i := 0; i < len(ha.cards); i++ {
		if ha.cards[i].rank == hb.cards[i].rank {
			continue
		}

		if ha.cards[i].rank > hb.cards[i].rank {
			return true
		}
		return false
	}
	return false
}
