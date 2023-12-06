package day5

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type mapRange struct {
	start int
	end   int
}
type mapping struct {
	destinations mapRange
	sources      mapRange
}

type almanac struct {
	seeds                 []mapRange
	seedToSoil            []mapping
	soilToFertilizer      []mapping
	fertilizerToWater     []mapping
	waterToLight          []mapping
	lightToTemperature    []mapping
	temperatureToHumidity []mapping
	humidityToLocation    []mapping
}

func part1(lines []string) int {
	alma := parseAlmanac(lines, false)
	return findMinLocation(alma)
}

func part2(lines []string) int {
	alma := parseAlmanac(lines, true)
	return findMinLocation(alma)
}

func findMinLocation(alma almanac) int {
	var mu sync.Mutex
	var wg sync.WaitGroup

	minLoc := -1
	for _, seedRange := range alma.seeds {
		for seed := seedRange.start; seed <= seedRange.end; seed++ {
			wg.Add(1)
			go func(seed int) {
				defer wg.Done()
				soilDest := findDestination(seed, alma.seedToSoil)
				fertDest := findDestination(soilDest, alma.soilToFertilizer)
				waterDest := findDestination(fertDest, alma.fertilizerToWater)
				lightDest := findDestination(waterDest, alma.waterToLight)
				tempDest := findDestination(lightDest, alma.lightToTemperature)
				humidDest := findDestination(tempDest, alma.temperatureToHumidity)
				locDest := int(findDestination(humidDest, alma.humidityToLocation))

				mu.Lock()
				if minLoc == -1 {
					minLoc = locDest
				} else if minLoc > locDest {
					minLoc = locDest
				}
				mu.Unlock()
			}(seed)
		}
	}
	wg.Wait()

	return minLoc
}

func findDestination(seed int, mapps []mapping) int {
	for _, mapp := range mapps {
		if seed >= mapp.sources.start && seed <= mapp.sources.end {
			slot := mapp.sources.end - seed
			return mapp.destinations.end - slot
		}
	}
	return seed
}
func parseAlmanac(lines []string, useRange bool) almanac {
	alma := almanac{
		seeds: []mapRange{},
	}

	almaPart := 1
	isDescLine := false
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			almaPart += 1
			isDescLine = true
			continue
		}

		if isDescLine {
			isDescLine = false
			continue
		}

		switch almaPart {
		case 1:
			seedsMatcher := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(line, -1)
			if useRange {
				for i := 0; i < len(seedsMatcher); i += 2 {
					seedStart, _ := strconv.Atoi(seedsMatcher[i][1])
					ran, _ := strconv.Atoi(seedsMatcher[i+1][1])
					length := ran - 1
					alma.seeds = append(alma.seeds, mapRange{start: seedStart, end: seedStart + length})
				}
			} else {
				for _, match := range seedsMatcher {
					numb, _ := strconv.Atoi(match[1])
					alma.seeds = append(alma.seeds, mapRange{start: numb, end: numb})
				}
			}
		case 2:
			numbsStg := strings.Split(line, " ")
			alma.seedToSoil = append(alma.seedToSoil, parseMapping(numbsStg))
		case 3:
			numbsStg := strings.Split(line, " ")
			alma.soilToFertilizer = append(alma.soilToFertilizer, parseMapping(numbsStg))
		case 4:
			numbsStg := strings.Split(line, " ")
			alma.fertilizerToWater = append(alma.fertilizerToWater, parseMapping(numbsStg))
		case 5:
			numbsStg := strings.Split(line, " ")
			alma.waterToLight = append(alma.waterToLight, parseMapping(numbsStg))
		case 6:
			numbsStg := strings.Split(line, " ")
			alma.lightToTemperature = append(alma.lightToTemperature, parseMapping(numbsStg))
		case 7:
			numbsStg := strings.Split(line, " ")
			alma.temperatureToHumidity = append(alma.temperatureToHumidity, parseMapping(numbsStg))
		case 8:
			numbsStg := strings.Split(line, " ")
			alma.humidityToLocation = append(alma.humidityToLocation, parseMapping(numbsStg))
		}
	}
	return alma
}

func parseMapping(numbsStr []string) mapping {
	dest, _ := strconv.Atoi(numbsStr[0])
	sour, _ := strconv.Atoi(numbsStr[1])
	ran, _ := strconv.Atoi(numbsStr[2])
	length := ran - 1

	mapp := mapping{
		destinations: mapRange{
			start: dest,
			end:   dest + length,
		},
		sources: mapRange{
			start: sour,
			end:   sour + length,
		},
	}

	return mapp
}
