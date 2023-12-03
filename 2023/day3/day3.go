package day3

import (
	"slices"
	"strconv"
)

type slot struct {
	isSymbol bool
	value    string
}

// Create a matrix for each [x, y] point in space
// 0   1   2   3   4   5   6   7   8   10
// 0  467 467 467  .   .  114 114 114  .   .
// 1   .   .   .   *   .   .   .   .   .   .
// 2   .   .   35  35  .   .  633 633 633  .
// 3   .   .   .   .   .   .   #   .   .   .

// Using [2, 2] as an example:
//     [1, 1] [2, 2] [3, 1]
//     [1, 2] [2, 2] [3, 2]
//     [1, 3] [2, 2] [3, 3]

// Starting on the top left and iterating x+2 and y+2:
// all surronding spaces will be checked
//     [2, 2] -> for x  -> for y
//              [1, 1]
//               [1, 2]
//               [1, 3]
//              [2, 1]
//               [2, 2]
//               [2, 3]
//              [3, 1]
//               [3, 2]
//               [3, 3]
// So 3 rows and 3 columns

func part1(gameLines []string) int {
	matrix := createMatrix(gameLines)

	total := 0
	pts := []*slot{}
	for y, row := range matrix {
		for x, slot := range row {
			if slot != nil && !slot.isSymbol && hasAdjacentSymbol(matrix, y, x) {
				if ok := slices.Contains(pts, slot); !ok {
					pts = append(pts, slot)
					numb, _ := strconv.Atoi(slot.value)
					total += numb
				}
			}
		}
	}
	return total
}

func part2(gameLines []string) int {
	matrix := createMatrix(gameLines)

	total := 0
	for y, row := range matrix {
		for x, slot := range row {
			if slot != nil && slot.isSymbol {
				parts := fetchGearParts(matrix, y, x)
				if len(parts) < 2 {
					continue
				}

				mult := 1
				for _, part := range parts {
					numb, _ := strconv.Atoi(part.value)
					mult *= numb
				}
				total += mult
			}
		}
	}
	return total
}

func fetchGearParts(matrix [][]*slot, y int, x int) []*slot {
	parts := []*slot{}
	// UP-LEFT
	if slot, ok := findSlot(matrix, y-1, x-1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// UP
	if slot, ok := findSlot(matrix, y-1, x); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// UP-RIGHT
	if slot, ok := findSlot(matrix, y-1, x+1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// LEFT
	if slot, ok := findSlot(matrix, y, x-1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// RIGHT
	if slot, ok := findSlot(matrix, y, x+1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// DOWN-LEFT
	if slot, ok := findSlot(matrix, y+1, x-1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// DOWN
	if slot, ok := findSlot(matrix, y+1, x); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	// DOWN-RIGHT
	if slot, ok := findSlot(matrix, y+1, x+1); ok &&
		slot != nil &&
		!slot.isSymbol &&
		!slices.Contains(parts, slot) {
		parts = append(parts, slot)
	}

	return parts
}

func hasAdjacentSymbol(matrix [][]*slot, y int, x int) bool {
	// UP-LEFT
	if slot, ok := findSlot(matrix, y-1, x-1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// UP
	if slot, ok := findSlot(matrix, y-1, x); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// UP-RIGHT
	if slot, ok := findSlot(matrix, y-1, x+1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// LEFT
	if slot, ok := findSlot(matrix, y, x-1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// RIGHT
	if slot, ok := findSlot(matrix, y, x+1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// DOWN-LEFT
	if slot, ok := findSlot(matrix, y+1, x-1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// DOWN
	if slot, ok := findSlot(matrix, y+1, x); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	// DOWN-RIGHT
	if slot, ok := findSlot(matrix, y+1, x+1); ok &&
		slot != nil &&
		slot.isSymbol {
		return true
	}

	return false
}

func findSlot(matrix [][]*slot, y int, x int) (*slot, bool) {
	if x < 0 || y < 0 || len(matrix)-1 < y || len(matrix[y])-1 < x {
		return nil, false
	}
	return matrix[y][x], true
}

func createMatrix(lines []string) [][]*slot {
	rowSize := len(lines)
	matrix := make([][]*slot, rowSize)

	for i, line := range lines {
		row := createRow(line)
		matrix[i] = row
	}
	return matrix
}

func createRow(line string) []*slot {
	colSize := len(line)
	row := make([]*slot, colSize)

	var curNumb *slot
	for i := 1; i <= colSize; i++ {
		if curNumb == nil {
			curNumb = &slot{
				value: "",
			}
		}

		char := line[i-1 : i]
		if _, err := strconv.Atoi(char); err == nil {
			curNumb.value += char
			row[i-1] = curNumb
		} else if char != "." {
			curNumb = nil
			row[i-1] = &slot{isSymbol: true, value: char}
		} else {
			curNumb = nil
			row[i-1] = curNumb
		}
	}

	return row
}
