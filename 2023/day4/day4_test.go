package day4

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestParseScratchCart(t *testing.T) {
	cardStg := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := card{
		id:             1,
		winningNumbers: []int{41, 48, 83, 86, 17},
		gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	card := parseScratchCart(cardStg)
	assert.Equal(t, expected, card)
}

func TestCalculateCardWorth(t *testing.T) {
	tests := []struct {
		card     card
		expected int
	}{
		{
			card: card{
				id:             1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expected: 8,
		},
		{
			card: card{
				id:             2,
				winningNumbers: []int{13, 32, 20, 16, 61},
				gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expected: 2,
		},
		{
			card: card{
				id:             3,
				winningNumbers: []int{1, 21, 53, 59, 44},
				gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			expected: 2,
		},
		{
			card: card{
				id:             4,
				winningNumbers: []int{41, 92, 73, 84, 69},
				gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			expected: 1,
		},
		{
			card: card{
				id:             5,
				winningNumbers: []int{87, 83, 26, 28, 32},
				gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expected: 0,
		},
		{
			card: card{
				id:             6,
				winningNumbers: []int{31, 18, 13, 56, 72},
				gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			expected: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("calculateCardWorth - test: %d", i), func(t *testing.T) {
			worth := calculateCardWorth(test.card)
			assert.Equal(t, test.expected, worth)
		})
	}
}

func TestCalculateWinningNumbersQuantity(t *testing.T) {
	tests := []struct {
		card     card
		expected int
	}{
		{
			card: card{
				id:             1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expected: 4,
		},
		{
			card: card{
				id:             2,
				winningNumbers: []int{13, 32, 20, 16, 61},
				gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expected: 2,
		},
		{
			card: card{
				id:             3,
				winningNumbers: []int{1, 21, 53, 59, 44},
				gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			expected: 2,
		},
		{
			card: card{
				id:             4,
				winningNumbers: []int{41, 92, 73, 84, 69},
				gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			expected: 1,
		},
		{
			card: card{
				id:             5,
				winningNumbers: []int{87, 83, 26, 28, 32},
				gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expected: 0,
		},
		{
			card: card{
				id:             6,
				winningNumbers: []int{31, 18, 13, 56, 72},
				gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			expected: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("calculateCardWorth - test: %d", i), func(t *testing.T) {
			qtn := calculateWinningNumbersQuantity(test.card)
			assert.Equal(t, test.expected, qtn)
		})
	}
}

func TestCalculateCopiesWon(t *testing.T) {
	tests := []struct {
		initCard card
		cards    []card
		expected int
	}{
		{
			initCard: card{
				id:             1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 15,
		},
		{
			initCard: card{
				id:             2,
				winningNumbers: []int{13, 32, 20, 16, 61},
				gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 7,
		},
		{
			initCard: card{
				id:             3,
				winningNumbers: []int{1, 21, 53, 59, 44},
				gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 4,
		},
		{
			initCard: card{
				id:             4,
				winningNumbers: []int{41, 92, 73, 84, 69},
				gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 2,
		},
		{
			initCard: card{
				id:             5,
				winningNumbers: []int{87, 83, 26, 28, 32},
				gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 1,
		},
		{
			initCard: card{
				id:             6,
				winningNumbers: []int{31, 18, 13, 56, 72},
				gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			},
			cards: []card{
				{
					id:             1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					gameNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{
					id:             2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					gameNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:             3,
					winningNumbers: []int{1, 21, 53, 59, 44},
					gameNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
				},
				{
					id:             4,
					winningNumbers: []int{41, 92, 73, 84, 69},
					gameNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
				},
				{
					id:             5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					gameNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
				{
					id:             6,
					winningNumbers: []int{31, 18, 13, 56, 72},
					gameNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
				},
			},
			expected: 1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("calculateCardWorth - test: %d", i), func(t *testing.T) {
			qtn := calculateCopiesWon(test.initCard, test.cards)
			assert.Equal(t, test.expected, qtn)
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: 13,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 25231,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part1(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			expected: 30,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 9721255,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part2(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}
