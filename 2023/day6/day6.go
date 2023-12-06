package day6

import (
	"regexp"
	"strconv"
)

type race struct {
	time    int
	recDist int
}

func part1(raceLines []string) int {
	races := parseRaces(raceLines)

	margin := 1
	for _, r := range races {
		wConds := findRaceWinningConditions(r)
		margin *= len(wConds)
	}

	return margin
}

func part2(raceLines []string) int {
	r := parseRace(raceLines)

	wConds := findRaceWinningConditions(r)
	margin := len(wConds)

	return margin
}

func parseRace(raceLines []string) race {
	matchTimes := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(raceLines[0], -1)
	matchDistances := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(raceLines[1], -1)

	raceTime := ""
	raceDist := ""
	for i := 0; i < len(matchTimes); i++ {
		raceTime += matchTimes[i][1]
		raceDist += matchDistances[i][1]
	}
	time, _ := strconv.Atoi(raceTime)
	dist, _ := strconv.Atoi(raceDist)

	return race{time: time, recDist: dist}
}

func parseRaces(raceLines []string) []race {
	matchTimes := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(raceLines[0], -1)
	matchDistances := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(raceLines[1], -1)

	races := []race{}
	for i := 0; i < len(matchTimes); i++ {
		time, _ := strconv.Atoi(matchTimes[i][1])
		dist, _ := strconv.Atoi(matchDistances[i][1])
		races = append(races, race{time: time, recDist: dist})
	}

	return races
}

// Button time -> travelled distance (bt*rt)
// 0 -> 0*7 = 0
// 1 -> 1*6 = 6
// 2 -> 2*5 = 10
// 3 -> 3*4 = 12
// 4 -> 4*3 = 12
// 5 -> 5*2 = 10
// 6 -> 6*1 = 6
// 7 -> 7*0 = 0
func findRaceWinningConditions(r race) []int {
	raceTime := r.time

	winBtnTimes := []int{}
	for bt := 1; bt < raceTime; bt++ {
		rt := raceTime - bt
		dist := bt * rt
		if dist <= r.recDist {
			continue
		}
		winBtnTimes = append(winBtnTimes, bt)
	}
	return winBtnTimes
}
