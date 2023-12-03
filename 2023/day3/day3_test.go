package day3

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestCreateRow(t *testing.T) {
	line := "467..114.."

	p1 := &slot{value: "467"}
	p2 := &slot{value: "114"}
	expected := []*slot{
		p1,
		p1,
		p1,
		nil,
		nil,
		p2,
		p2,
		p2,
		nil,
		nil,
	}

	row := createRow(line)
	assert.Equal(t, expected, row)
}

func TestCreateMatrix(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
	}
	p1 := &slot{value: "467"}
	p2 := &slot{value: "114"}
	p3 := &slot{value: "35"}
	p4 := &slot{value: "633"}
	p5 := &slot{
		value: "617",
	}
	symbolP1 := &slot{isSymbol: true, value: "*"}
	symbolP2 := &slot{isSymbol: true, value: "#"}

	expected := [][]*slot{
		{
			p1,
			p1,
			p1,
			nil,
			nil,
			p2,
			p2,
			p2,
			nil,
			nil,
		},
		{
			nil,
			nil,
			nil,
			symbolP1,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
		{
			nil,
			nil,
			p3,
			p3,
			nil,
			nil,
			p4,
			p4,
			p4,
			nil,
		},
		{
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			symbolP2,
			nil,
			nil,
			nil,
		},
		{
			p5,
			p5,
			p5,
			symbolP1,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
	}

	res := createMatrix(lines)

	assert.Equal(t, expected, res)
}

func TestHasAdjacentSymbol(t *testing.T) {
	p1 := &slot{value: "467"}
	p2 := &slot{value: "114"}
	symbolP := &slot{isSymbol: true, value: "*"}

	tests := []struct {
		matrix   [][]*slot
		curY     int
		curX     int
		expected bool
	}{
		{
			matrix: [][]*slot{
				{
					p1,
					p1,
					p1,
					nil,
					nil,
					p2,
					p2,
					p2,
					nil,
					nil,
				},
				{
					nil,
					nil,
					nil,
					symbolP,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
				},
			},
			curY:     0,
			curX:     2,
			expected: true,
		},
		{
			matrix: [][]*slot{
				{
					p1,
					p1,
					p1,
					nil,
					nil,
					p2,
					p2,
					p2,
					nil,
					nil,
				},
				{
					nil,
					nil,
					nil,
					symbolP,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
				},
			},
			curY:     0,
			curX:     5,
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("hasAdjacentSymbol - test: %d", i), func(t *testing.T) {
			res := hasAdjacentSymbol(test.matrix, test.curY, test.curX)
			assert.Equal(t, res, test.expected)
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
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 517021,
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
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 467835,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 81296995,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := part2(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}
