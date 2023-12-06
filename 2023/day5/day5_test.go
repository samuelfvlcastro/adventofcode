package day5

import (
	"aoc/2023/utils"
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestParseAlmanac(t *testing.T) {
	almaLines := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	expected := almanac{
		seeds: []mapRange{
			{start: 79, end: 79},
			{start: 14, end: 14},
			{start: 55, end: 55},
			{start: 13, end: 13},
		},
		seedToSoil: []mapping{
			{
				destinations: mapRange{start: 50, end: 51},
				sources:      mapRange{start: 98, end: 99},
			},
			{
				destinations: mapRange{start: 52, end: 99},
				sources:      mapRange{start: 50, end: 97},
			},
		},
		soilToFertilizer: []mapping{
			{
				destinations: mapRange{start: 0, end: 36},
				sources:      mapRange{start: 15, end: 51},
			},
			{
				destinations: mapRange{start: 37, end: 38},
				sources:      mapRange{start: 52, end: 53},
			},
			{
				destinations: mapRange{start: 39, end: 53},
				sources:      mapRange{start: 0, end: 14},
			},
		},
		fertilizerToWater: []mapping{
			{
				destinations: mapRange{start: 49, end: 56},
				sources:      mapRange{start: 53, end: 60},
			},
			{
				destinations: mapRange{start: 0, end: 41},
				sources:      mapRange{start: 11, end: 52},
			},
			{
				destinations: mapRange{start: 42, end: 48},
				sources:      mapRange{start: 0, end: 6},
			},
			{
				destinations: mapRange{start: 57, end: 60},
				sources:      mapRange{start: 7, end: 10},
			},
		},
		waterToLight: []mapping{
			{
				destinations: mapRange{start: 88, end: 94},
				sources:      mapRange{start: 18, end: 24},
			},
			{
				destinations: mapRange{start: 18, end: 87},
				sources:      mapRange{start: 25, end: 94},
			},
		},
		lightToTemperature: []mapping{
			{
				destinations: mapRange{start: 45, end: 67},
				sources:      mapRange{start: 77, end: 99},
			},
			{
				destinations: mapRange{start: 81, end: 99},
				sources:      mapRange{start: 45, end: 63},
			},
			{
				destinations: mapRange{start: 68, end: 80},
				sources:      mapRange{start: 64, end: 76},
			},
		},
		temperatureToHumidity: []mapping{
			{
				destinations: mapRange{start: 0, end: 0},
				sources:      mapRange{start: 69, end: 69},
			},
			{
				destinations: mapRange{start: 1, end: 69},
				sources:      mapRange{start: 0, end: 68},
			},
		},
		humidityToLocation: []mapping{
			{
				destinations: mapRange{start: 60, end: 96},
				sources:      mapRange{start: 56, end: 92},
			},
			{
				destinations: mapRange{start: 56, end: 59},
				sources:      mapRange{start: 93, end: 96},
			},
		},
	}

	alma := parseAlmanac(almaLines, false)
	assert.Equal(t, alma, expected)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{
			lines: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"				",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"				",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"				",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"				",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"				",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"				",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 35,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 51580674,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := Part1(test.lines)
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
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"				",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"				",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"				",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"				",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"				",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"				",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 46,
		},
		{
			lines:    utils.MustFetchInput("./input.txt"),
			expected: 99751240,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("part1 - test: %d", i), func(t *testing.T) {
			total := Part2(test.lines)
			assert.Equal(t, test.expected, total)
		})
	}
}
