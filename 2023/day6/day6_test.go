package day6

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestParseRaces(t *testing.T) {
	racesLines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expected := []race{
		{
			time:    7,
			recDist: 9,
		},
		{
			time:    15,
			recDist: 40,
		},
		{
			time:    30,
			recDist: 200,
		},
	}
	races := parseRaces(racesLines)
	assert.Equal(t, expected, races)
}

func TestFindRaceWinningConditions(t *testing.T) {
	tests := []struct {
		race     race
		expected []int
	}{
		{
			race: race{
				time:    7,
				recDist: 9,
			},
			expected: []int{2, 3, 4, 5},
		},
		{
			race: race{
				time:    15,
				recDist: 40,
			},
			expected: []int{4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			race: race{
				time:    30,
				recDist: 200,
			},
			expected: []int{11, 12, 13, 14, 15, 16, 17, 18, 19},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("findRaceWinningConditions - test: %d", i), func(t *testing.T) {
			winBtnTimes := findRaceWinningConditions(test.race)
			assert.Equal(t, test.expected, winBtnTimes)
		})
	}
}

func TestParseRace(t *testing.T) {
	raceLines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	expected := race{
		time:    71530,
		recDist: 940200,
	}
	races := parseRace(raceLines)
	assert.Equal(t, expected, races)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 288,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 131376,
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
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 71503,
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
