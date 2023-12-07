package day7

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestParseHand(t *testing.T) {
	line := "32T3K 765"
	expected := hand{
		cards: []card{
			{symbol: "3", rank: 2},
			{symbol: "2", rank: 1},
			{symbol: "T", rank: 9},
			{symbol: "3", rank: 2},
			{symbol: "K", rank: 12}},
		bid: 765,
	}

	h := parseHand(line, ranks)
	assert.Equal(t, expected, h)
}

func TestFindHandStrength(t *testing.T) {
	tests := []struct {
		hand     hand
		expected int
	}{
		{
			hand: hand{
				cards: []card{
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
				},
				bid: 765,
			},
			expected: 1,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "3", rank: 2},
				},
				bid: 765,
			},
			expected: 2,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "3", rank: 2},
					{symbol: "3", rank: 2},
				},
				bid: 765,
			},
			expected: 3,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "T", rank: 9},
					{symbol: "3", rank: 2},
					{symbol: "2", rank: 1},
				},
				bid: 765,
			},
			expected: 4,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "4", rank: 3},
					{symbol: "4", rank: 3},
					{symbol: "3", rank: 2},
					{symbol: "3", rank: 2},
					{symbol: "2", rank: 1},
				},
				bid: 765,
			},
			expected: 5,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "K", rank: 12},
					{symbol: "K", rank: 12},
					{symbol: "6", rank: 5},
					{symbol: "5", rank: 4},
					{symbol: "4", rank: 3},
				},
				bid: 765,
			},
			expected: 6,
		},
		{
			hand: hand{
				cards: []card{
					{symbol: "K", rank: 12},
					{symbol: "T", rank: 9},
					{symbol: "4", rank: 3},
					{symbol: "3", rank: 2},
					{symbol: "2", rank: 1},
				},
				bid: 765,
			},
			expected: 7,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("findHandStrength - test %d", i), func(t *testing.T) {
			s := findHandStrength(test.hand, false)
			assert.Equal(t, test.expected, s)
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
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 6440,
		},
		{
			lines: []string{
				"2345A 1",
				"Q2KJJ 13",
				"Q2Q2Q 19",
				"T3T3J 17",
				"T3Q33 11",
				"2345J 3",
				"J345A 2",
				"32T3K 5",
				"T55J5 29",
				"KK677 7",
				"KTJJT 34",
				"QQQJA 31",
				"JJJJJ 37",
				"JAAAA 43",
				"AAAAJ 59",
				"AAAAA 61",
				"2AAAA 23",
				"2JJJJ 53",
				"JJJJ2 41",
			},
			expected: 6592,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 253313241,
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
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			expected: 5905,
		},
		{
			lines: []string{
				"2345A 1",
				"Q2KJJ 13",
				"Q2Q2Q 19",
				"T3T3J 17",
				"T3Q33 11",
				"2345J 3",
				"J345A 2",
				"32T3K 5",
				"T55J5 29",
				"KK677 7",
				"KTJJT 34",
				"QQQJA 31",
				"JJJJJ 37",
				"JAAAA 43",
				"AAAAJ 59",
				"AAAAA 61",
				"2AAAA 23",
				"2JJJJ 53",
				"JJJJ2 41",
			},
			expected: 6839,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 34123437,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part2(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}
