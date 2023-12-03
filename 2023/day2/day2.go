package day2

import (
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id    int
	rolls []map[string]int
}

func part1(gameLines []string) int {
	total := 0
	for _, line := range gameLines {
		g := parseGameLine(line)
		if !isGamePossible(g) {
			continue
		}
		total += g.id
	}
	return total
}

func part2(gameLines []string) int {
	powerSum := 0
	for _, line := range gameLines {
		g := parseGameLine(line)
		set := maxCubesPerGamePerColor(g)
		power := cubeSetPower(set)
		powerSum += power
	}
	return powerSum
}

func parseGameLine(gameLine string) game {
	recordMatches := regexp.MustCompile(`(Game\s\d+|\d+\sblue|\d+\sred|\d+\sgreen)`).FindAllStringSubmatch(gameLine, -1)
	idParts := strings.Split(recordMatches[0][1], " ")
	id, _ := strconv.Atoi(idParts[1])
	g := game{
		id:    id,
		rolls: []map[string]int{},
	}

	for _, record := range recordMatches[1:] {
		rolls := map[string]int{}
		if len(record) > 1 {
			roll := strings.Split(record[1], " ")
			qnt, _ := strconv.Atoi(roll[0])
			color := roll[1]
			rolls[color] += qnt
		}
		g.rolls = append(g.rolls, rolls)
	}

	return g
}

var existingCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isGamePossible(g game) bool {
	for _, rolls := range g.rolls {
		for color, roll := range rolls {
			if roll > existingCubes[color] {
				return false
			}
		}
	}
	return true
}

func maxCubesPerGamePerColor(g game) map[string]int {
	maxVals := map[string]int{}
	for _, rolls := range g.rolls {
		for color, qnt := range rolls {
			prevQnt, ok := maxVals[color]
			if !ok {
				maxVals[color] = qnt
				continue
			}

			if qnt > prevQnt {
				maxVals[color] = qnt
			}
		}
	}
	return maxVals
}

func cubeSetPower(set map[string]int) int {
	power := 1
	for _, qtn := range set {
		power *= qtn
	}
	return power
}
