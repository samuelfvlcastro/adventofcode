package day1

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines:    []string{"1abc2"},
			expected: 12,
		},
		{
			lines:    []string{"pqr3stu8vwx"},
			expected: 38,
		},
		{
			lines:    []string{"a1b2c3d4e5f"},
			expected: 15,
		},
		{
			lines:    []string{"treb7uchet"},
			expected: 77,
		},
		{
			lines: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			expected: 142,
		},
		// puzzle answer
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 54644,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			result := part1(test.lines)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines:    []string{"two1nine"},
			expected: 29,
		},
		{
			lines:    []string{"eightwothree"},
			expected: 83,
		},
		{
			lines:    []string{"abcone2threexyz"},
			expected: 13,
		},
		{
			lines:    []string{"xtwone3four"},
			expected: 24,
		},
		{
			lines:    []string{"4nineeightseven2"},
			expected: 42,
		},
		{
			lines:    []string{"zoneight234"},
			expected: 14,
		},
		{
			lines:    []string{"7pqrstsixteen"},
			expected: 76,
		},
		{
			lines: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expected: 281,
		},
		// puzzle answer
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 53348,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			result := part2(test.lines)
			assert.Equal(t, test.expected, result)
		})
	}
}
