package day4

import (
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	id             int
	winningNumbers []int
	gameNumbers    []int
}

func part1(cardLines []string) int {
	total := 0
	for _, cardStg := range cardLines {
		c := parseScratchCart(cardStg)
		worth := calculateCardWorth(c)
		total += worth
	}
	return total
}

func part2(cardLines []string) int {
	cards := []card{}
	for _, cardStg := range cardLines {
		c := parseScratchCart(cardStg)
		cards = append(cards, c)
	}

	cardsWon := 0
	for _, card := range cards {
		cardsWon += calculateCopiesWon(card, cards)
	}

	return cardsWon
}

func calculateCopiesWon(c card, deck []card) int {
	// The card itself
	cardsWon := 1
	lastCardIdx := len(deck) - 1
	wNumbers := calculateWinningNumbersQuantity(c)
	if wNumbers > 0 {
		nextCardIdx := c.id
		for i := 0; i < wNumbers; i++ {
			if nextCardIdx > lastCardIdx {
				break
			}

			// Cards won by playing the winning card
			cardsWon += calculateCopiesWon(deck[nextCardIdx], deck)

			nextCardIdx++
		}
	}
	return cardsWon
}

func parseScratchCart(cardStg string) card {
	cardMatches := regexp.MustCompile(`(\|\s|(\d+))`).FindAllStringSubmatch(cardStg, -1)

	id, _ := strconv.Atoi(cardMatches[0][1])

	gameNumbers := []int{}
	winningNumbers := []int{}
	switchToGameNumbers := false
	for i := 1; i < len(cardMatches); i++ {
		if strings.TrimSpace(cardMatches[i][1]) == "|" {
			switchToGameNumbers = true
			continue
		}

		numb, _ := strconv.Atoi(cardMatches[i][1])
		if switchToGameNumbers {
			gameNumbers = append(gameNumbers, numb)
			continue
		}
		winningNumbers = append(winningNumbers, numb)
	}

	return card{
		id:             id,
		winningNumbers: winningNumbers,
		gameNumbers:    gameNumbers,
	}
}

func calculateCardWorth(c card) int {
	worth := 0
	firstWin := true
	for _, wNumber := range c.winningNumbers {
		for _, number := range c.gameNumbers {
			if wNumber == number {
				if firstWin {
					worth = 1
					firstWin = false
					continue
				}
				worth *= 2
			}
		}
	}

	return worth
}

func calculateWinningNumbersQuantity(c card) int {
	winningNumbers := 0
	for _, wNumber := range c.winningNumbers {
		for _, number := range c.gameNumbers {
			if wNumber == number {
				winningNumbers++
			}
		}
	}

	return winningNumbers
}
