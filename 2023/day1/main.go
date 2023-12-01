package main

import (
	"aoc/2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

// Day1: https://adventofcode.com/2023/day/1
func main() {
	partOne()
	partTwo()
}

// The newly-improved calibration document consists of lines of text;
// each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
//
// For example:
//
//	1abc2
//	pqr3stu8vwx
//	a1b2c3d4e5f
//	treb7uchet
//
// In this example, the calibration values of these four lines are 12, 38, 15, and 77.
// Adding these together produces 142.
// Resp: 54644
func partOne() {
	input, err := utils.FetchInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	for _, line := range input {
		chars := []rune(line)

		left := 0
		right := len(chars) - 1

		var lNumb, rNumb string
		for i := 0; i <= len(chars)-1; i++ {
			if ok := unicode.IsNumber(chars[left]); ok && lNumb == "" {
				lNumb = string(chars[left])
			}

			if ok := unicode.IsNumber(chars[right]); ok && rNumb == "" {
				rNumb = string(chars[right])
			}

			if lNumb == "" {
				left++
			}

			if rNumb == "" {
				right--
			}
		}
		numb, _ := strconv.Atoi(lNumb + rNumb)
		total += numb
	}
	fmt.Println(total)
}

// Your calculation isn't quite right.
// It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".
//
// Equipped with this new information, you now need to find the real first and last digit on each line. For example:
// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
// In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
// Resp: 53348
func partTwo() {
	input, err := utils.FetchInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, line := range input {
		chars := []rune(line)

		l := 0
		var lSub, lNumb string

		r := len(chars) - 1
		var rSub, rNumb string

		for i := 0; i < len(chars); i++ {
			lc := chars[l]
			if unicode.IsNumber(lc) && lNumb == "" {
				lNumb = string(lc)
			}

			lSub += string(lc)
			if numb := fetchNumberStr(lSub); numb != "" && lNumb == "" {
				lNumb = numb
			}
			l++

			rc := chars[r]
			if unicode.IsNumber(rc) && rNumb == "" {
				rNumb = string(rc)
			}

			rSub = string(rc) + rSub
			if numb := fetchNumberStr(rSub); numb != "" && rNumb == "" {
				rNumb = numb
			}
			r--

			if lNumb != "" && rNumb != "" {
				break
			}
		}

		numb, _ := strconv.Atoi(lNumb + rNumb)
		total += numb
	}
	fmt.Println(total)
}

var strToNumbMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func fetchNumberStr(sub string) string {
	for str, numb := range strToNumbMap {
		if strings.Contains(sub, str) {
			return numb
		}
	}
	return ""
}
