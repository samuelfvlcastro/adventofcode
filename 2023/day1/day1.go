package day1

import (
	"strconv"
	"strings"
	"unicode"
)

// The newly-improved calibration document consists of lines of text;
// each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit (in that order)
// to form a single two-digit number.
func part1(lines []string) int {
	total := 0
	for _, line := range lines {
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
	return total
}

// Your calculation isn't quite right.
// It looks like some of the digits are actually spelled out with letters:
// one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".
func part2(lines []string) int {
	total := 0
	for _, line := range lines {
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

	return total
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
