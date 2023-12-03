package day2

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestPart1ParseGameLine(t *testing.T) {
	expected := game{
		id: 13,
		rolls: []map[string]int{
			{
				"green": 8,
			},
			{
				"blue": 6,
			},
			{
				"red": 20,
			},
			{
				"blue": 5,
			},
			{
				"red": 4,
			},
			{
				"green": 13,
			},
			{
				"green": 5,
			},
			{
				"red": 1,
			},
		},
	}
	res := parseGameLine("Game 13: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	assert.Equal(t, expected, res)
}

func TestPart1IsPossible(t *testing.T) {
	tests := []struct {
		line     string
		expected bool
	}{
		{
			line:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: true,
		},
		{
			line:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: true,
		},
		{
			line:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: false,
		},
		{
			line:     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: false,
		},
		{
			line:     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("parseGameLine - test: %d", i), func(t *testing.T) {
			game := parseGameLine(test.line)
			result := isGamePossible(game)
			assert.Equal(t, test.expected, result)
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
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expected: 8,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 8,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part1(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}

func TestMaxCubesPerGamePerColor(t *testing.T) {
	tests := []struct {
		g                game
		expectedMinCubes map[string]int
		expectedPower    int
	}{
		{
			g: game{
				id: 13,
				rolls: []map[string]int{
					{
						"blue": 3,
					},
					{
						"red": 4,
					},
					{
						"red": 1,
					},
					{
						"green": 2,
					},
					{
						"blue": 6,
					},
					{
						"green": 2,
					},
				},
			},
			expectedMinCubes: map[string]int{
				"red":   4,
				"green": 2,
				"blue":  6,
			},
			expectedPower: 48,
		},
		{
			g: game{
				id: 13,
				rolls: []map[string]int{
					{
						"blue": 1,
					},
					{
						"green": 2,
					},
					{
						"green": 3,
					},
					{
						"blue": 4,
					},
					{
						"red": 1,
					},
					{
						"green": 1,
					},
					{
						"blue": 1,
					},
				},
			},
			expectedMinCubes: map[string]int{
				"red":   1,
				"green": 3,
				"blue":  4,
			},
			expectedPower: 12,
		},
		{
			g: game{
				id: 13,
				rolls: []map[string]int{
					{
						"green": 8,
					},
					{
						"blue": 6,
					},
					{
						"red": 20,
					},
					{
						"blue": 5,
					},
					{
						"red": 4,
					},
					{
						"green": 13,
					},
					{
						"green": 5,
					},
					{
						"red": 1,
					},
				},
			},
			expectedMinCubes: map[string]int{
				"red":   20,
				"green": 13,
				"blue":  6,
			},
			expectedPower: 1560,
		},
		{
			g: game{
				id: 13,
				rolls: []map[string]int{
					{
						"green": 1,
					},
					{
						"red": 3,
					},
					{
						"blue": 6,
					},
					{
						"green": 3,
					},
					{
						"red": 6,
					},
					{
						"green": 3,
					},
					{
						"blue": 15,
					},
					{
						"red": 14,
					},
				},
			},
			expectedMinCubes: map[string]int{
				"red":   14,
				"green": 3,
				"blue":  15,
			},
			expectedPower: 630,
		},
		{
			g: game{
				id: 13,
				rolls: []map[string]int{
					{
						"red": 6,
					},
					{
						"blue": 1,
					},
					{
						"green": 3,
					},
					{
						"blue": 2,
					},
					{
						"red": 1,
					},
					{
						"green": 2,
					},
				},
			},
			expectedMinCubes: map[string]int{
				"red":   6,
				"green": 3,
				"blue":  2,
			},
			expectedPower: 36,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("maxCubesPerGamePerColor - test: %d", i), func(t *testing.T) {
			set := maxCubesPerGamePerColor(test.g)
			assert.Equal(t, test.expectedMinCubes, set)

			power := cubeSetPower(set)
			assert.Equal(t, test.expectedPower, power)
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
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			expected: 2286,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 8,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part2(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}
